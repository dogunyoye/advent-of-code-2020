package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type edge struct {
	id   int
	edge string
}

type pair struct {
	id   int
	tile []string
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func copyTile(tile []string) []string {
	var copy []string
	for _, line := range tile {
		copy = append(copy, string(line))
	}

	return copy
}

func printTile(tile []string) {
	for _, line := range tile {
		fmt.Println(line)
	}
	fmt.Println()
}

// rotate the tile 90 degrees clockwise
func rotateTile(tile []string) []string {
	var rotated []string

	for i := 0; i < len(tile); i++ {
		var rotString = ""
		for j := len(tile) - 1; j >= 0; j-- {
			rotString += string(tile[j][i])
		}
		rotated = append(rotated, rotString)
	}

	return rotated
}

func flipTileHorizontal(tile []string) []string {
	var flippedHorizontal []string

	for i := 0; i < len(tile); i++ {
		flippedHorizontal = append(flippedHorizontal, reverse(tile[i]))
	}

	return flippedHorizontal
}

func allTileCombos(tile []string) [][]string {
	var allTiles [][]string
	var rotated = copyTile(tile)
	var flipped = copyTile(tile)

	allTiles = append(allTiles, copyTile(tile))
	for i := 0; i < 3; i++ {
		rotated = rotateTile(rotated)
		allTiles = append(allTiles, rotated)
	}

	flipped = flipTileHorizontal(tile)
	allTiles = append(allTiles, flipped)
	for i := 0; i < 3; i++ {
		flipped = rotateTile(flipped)
		allTiles = append(allTiles, flipped)
	}

	return allTiles
}

func orientateCorner(id int, tilesMap map[int][]string, tileNeighbours map[int][]edge) pair {
	var tile = tilesMap[id]
	var allCombos = allTileCombos(tile)

	var neighbours = tileNeighbours[id]
	var allCombosTileN0 = allTileCombos(tilesMap[neighbours[0].id])
	var allCombosTileN1 = allTileCombos(tilesMap[neighbours[1].id])

	for _, t := range allCombos {
		var edges = getTileEdgesNoReverse(id, t)
		var rightEdge = edges[3]
		var bottomEdge = edges[2]

		for _, t0 := range allCombosTileN0 {
			var edgesN0 = getTileEdgesNoReverse(neighbours[0].id, t0)
			for _, t1 := range allCombosTileN1 {
				var edgesN1 = getTileEdgesNoReverse(neighbours[1].id, t1)
				if (edgesN0[1].edge == rightEdge.edge || edgesN0[0].edge == bottomEdge.edge) &&
					(edgesN1[1].edge == rightEdge.edge || edgesN1[0].edge == bottomEdge.edge) {
					return pair{id, t}
				}
			}
		}
	}

	panic("No combination found")
}

func buildFirstColumn(tile pair, tilesMap map[int][]string, tileNeighbours map[int][]edge) []pair {
	var firstColumn []pair
	firstColumn = append(firstColumn, tile)

	var currentTile = tile

	for {
		var neighbours = tileNeighbours[currentTile.id]
		var found = false

		var currentTileEdges = getTileEdgesNoReverse(currentTile.id, currentTile.tile)
		for _, n := range neighbours {
			var allCombos = allTileCombos(tilesMap[n.id])
			for _, tc := range allCombos {
				var edges = getTileEdgesNoReverse(n.id, tc)
				if currentTileEdges[2].edge == edges[0].edge {
					firstColumn = append(firstColumn, pair{n.id, tc})
					currentTile = pair{n.id, tc}
					found = true
				}
			}
		}

		if !found {
			// reached the end
			return firstColumn
		}
	}
}

func buildRows(firstColumn []pair, tilesMap map[int][]string, tileNeighbours map[int][]edge) [][]pair {
	var rows [][]pair
	for i := 0; i < len(firstColumn); i++ {
		var currentTile = firstColumn[i]
		var currentRow []pair

		currentRow = append(currentRow, currentTile)

		for {
			var neighbours = tileNeighbours[currentTile.id]
			var found = false

			var currentTileEdges = getTileEdgesNoReverse(currentTile.id, currentTile.tile)
			for _, n := range neighbours {
				var allCombos = allTileCombos(tilesMap[n.id])
				for _, tc := range allCombos {
					var edges = getTileEdgesNoReverse(n.id, tc)
					if currentTileEdges[3].edge == edges[1].edge {
						currentRow = append(currentRow, pair{n.id, tc})
						currentTile = pair{n.id, tc}
						found = true
					}
				}
			}

			if !found {
				// reached the end
				break
			}
		}

		rows = append(rows, currentRow)
	}

	return rows
}

func countWater(image []string) int {
	var count = 0
	for _, line := range image {
		count += strings.Count(line, "#")
	}
	return count
}

func findSeaMonsters(image []string) int {
	var count = 0
	var length = len(image[0])

	for i := 0; i < len(image)-3; i++ {
		var line0 = image[i]
		var line1 = image[i+1]
		var line2 = image[i+2]

		for j, k := length, length-20; k >= 0; j, k = j-1, k-1 {
			var slice0 = line0[k:j]
			var slice1 = line1[k:j]
			var slice2 = line2[k:j]

			if slice0[18] == '#' && slice1[0] == '#' && slice1[5] == '#' &&
				slice1[6] == '#' && slice1[11] == '#' && slice1[12] == '#' &&
				slice1[17] == '#' && slice1[18] == '#' && slice1[19] == '#' &&
				slice2[1] == '#' && slice2[4] == '#' && slice2[7] == '#' &&
				slice2[10] == '#' && slice2[13] == '#' && slice2[16] == '#' {
				count++
			}
		}
	}

	return count
}

func getTileEdges(id int, tile []string) []edge {

	var edges []edge

	// get top edge
	edges = append(edges, edge{id, tile[0]})
	edges = append(edges, edge{id, reverse(tile[0])})

	// get left edge
	var leftEdge = ""
	for _, line := range tile {
		leftEdge += string(line[0])
	}

	edges = append(edges, edge{id, leftEdge})
	edges = append(edges, edge{id, reverse(leftEdge)})

	// get bottom edge
	edges = append(edges, edge{id, tile[len(tile)-1]})
	edges = append(edges, edge{id, reverse(tile[len(tile)-1])})

	// right edge
	var rightEdge = ""
	for _, line := range tile {
		rightEdge += string(line[len(line)-1])
	}
	edges = append(edges, edge{id, rightEdge})
	edges = append(edges, edge{id, reverse(rightEdge)})

	return edges
}

func getTileEdgesNoReverse(id int, tile []string) []edge {

	var edges = make([]edge, 4)

	// get top edge
	edges[0] = edge{id, tile[0]}

	// get left edge
	var leftEdge = ""
	for _, line := range tile {
		leftEdge += string(line[0])
	}
	edges[1] = edge{id, leftEdge}

	// get bottom edge
	edges[2] = edge{id, tile[len(tile)-1]}

	// right edge
	var rightEdge = ""
	for _, line := range tile {
		rightEdge += string(line[len(line)-1])
	}
	edges[3] = edge{id, rightEdge}

	return edges
}

func findCornerTiles(tilesMap map[int][]string, tileNeighbours map[int][]edge) []int {
	var tileIDs []int

	for k, v := range tilesMap {
		tEdges := getTileEdgesNoReverse(k, v)
		var neighbours []edge

		var count = 0
		for _, e := range tEdges {
			for kk, vv := range tilesMap {
				if kk == k {
					continue
				}

				var matchFound = false
				candidates := getTileEdges(kk, vv)
				for _, c := range candidates {
					if c.edge == e.edge {
						count++
						matchFound = true
						neighbours = append(neighbours, c)
						break
					}
				}

				if matchFound {
					break
				}
			}
		}

		tileNeighbours[k] = neighbours

		// corner tiles only have
		// two adjacent tiles
		if count == 2 {
			tileIDs = append(tileIDs, k)
		}
	}

	return tileIDs
}

func main() {
	file, err := os.Open("../../data/day20.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var tilesMap = make(map[int][]string)
	var tilesNeigbours = make(map[int][]edge)

	for scanner.Scan() {
		var line = scanner.Text()
		if len(line) != 0 {
			if strings.Contains(line, "Tile") {
				line = strings.ReplaceAll(line, ":", "")
				split := strings.Split(line, " ")
				tileID, _ := strconv.Atoi(split[1])

				scanner.Scan()

				var tile []string
				tile = append(tile, scanner.Text())
				for scanner.Scan() {
					if len(scanner.Text()) == 0 {
						// empty line
						break
					}

					tile = append(tile, scanner.Text())
				}

				tilesMap[tileID] = tile
			}
		}
	}

	var part1 = 1
	var cornerTileId = 0
	for _, id := range findCornerTiles(tilesMap, tilesNeigbours) {
		part1 *= id
		cornerTileId = id
	}

	fmt.Println("Part1:", part1)

	var topLeftCorner = orientateCorner(cornerTileId, tilesMap, tilesNeigbours)
	var firstColumn = buildFirstColumn(topLeftCorner, tilesMap, tilesNeigbours)
	var image = buildRows(firstColumn, tilesMap, tilesNeigbours)

	var bigImage []string

	for _, p := range image {
		var row = ""
		for i := 1; i < 9; i++ {
			for _, t := range p {
				row += t.tile[i][1:9]
			}
			bigImage = append(bigImage, row)
			row = ""
		}
	}

	for _, t := range allTileCombos(bigImage) {
		var seaMonsters = findSeaMonsters(t)
		if seaMonsters != 0 {
			fmt.Println("Part2:", countWater(bigImage)-(15*seaMonsters))
			break
		}
	}

	file.Close()
}
