package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func getTileEdges(tile []string) []string {

	var edges []string

	// get top edge
	edges = append(edges, tile[0])
	edges = append(edges, reverse(tile[0]))

	// get left edge
	var leftEdge = ""
	for _, line := range tile {
		leftEdge += string(line[0])
	}
	edges = append(edges, leftEdge)
	edges = append(edges, reverse(leftEdge))

	// get bottom edge
	edges = append(edges, tile[len(tile)-1])
	edges = append(edges, reverse(tile[len(tile)-1]))

	// right edge
	var rightEdge = ""
	for _, line := range tile {
		rightEdge += string(line[len(line)-1])
	}
	edges = append(edges, rightEdge)
	edges = append(edges, reverse(rightEdge))

	return edges
}

func getTileEdgesNoReverse(tile []string) []string {

	var edges []string

	// get top edge
	edges = append(edges, tile[0])

	// get left edge
	var leftEdge = ""
	for _, line := range tile {
		leftEdge += string(line[0])
	}
	edges = append(edges, leftEdge)

	// get bottom edge
	edges = append(edges, tile[len(tile)-1])

	// right edge
	var rightEdge = ""
	for _, line := range tile {
		rightEdge += string(line[len(line)-1])
	}
	edges = append(edges, rightEdge)

	return edges
}

func findCornerTiles(tilesMap map[int][]string) []int {
	var tileIDs []int

	for k, v := range tilesMap {
		tEdges := getTileEdgesNoReverse(v)

		var count = 0
		for _, e := range tEdges {
			for kk, vv := range tilesMap {
				if kk == k {
					continue
				}

				var matchFound = false
				candidates := getTileEdges(vv)
				for _, c := range candidates {
					if c == e {
						count++
						matchFound = true
						break
					}
				}

				if matchFound {
					break
				}
			}
		}

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
	for _, ids := range findCornerTiles(tilesMap) {
		part1 *= ids
	}

	fmt.Println("Part1:", part1)

	file.Close()
}
