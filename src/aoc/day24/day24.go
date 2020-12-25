package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type direction string

const (
	northEast direction = "ne"
	east                = "e"
	southEast           = "se"
	southWest           = "sw"
	west                = "w"
	northWest           = "nw"
)

type pos struct {
	X int
	Y int
}

func getTileNeighbours(position pos) []pos {
	var positions []pos

	var northEastPos pos
	if position.Y%2 == 0 {
		northEastPos = pos{position.X, position.Y - 1}
	} else {
		northEastPos = pos{position.X + 1, position.Y - 1}
	}

	var southEastPos pos
	if position.Y%2 == 0 {
		southEastPos = pos{position.X, position.Y + 1}
	} else {
		southEastPos = pos{position.X + 1, position.Y + 1}
	}

	var southWestPos pos
	if position.Y%2 == 0 {
		southWestPos = pos{position.X - 1, position.Y + 1}
	} else {
		southWestPos = pos{position.X, position.Y + 1}
	}

	var northWestPos pos
	if position.Y%2 == 0 {
		northWestPos = pos{position.X - 1, position.Y - 1}
	} else {
		northWestPos = pos{position.X, position.Y - 1}
	}

	positions = append(positions, northEastPos)
	positions = append(positions, pos{position.X + 1, position.Y})
	positions = append(positions, southEastPos)
	positions = append(positions, southWestPos)
	positions = append(positions, pos{position.X - 1, position.Y})
	positions = append(positions, northWestPos)

	return positions
}

func checkBlackTiles(position pos, tilesMap map[pos]rune) int {
	var neighbours = getTileNeighbours(position)
	var result = 0

	for _, n := range neighbours {
		_, exists := tilesMap[n]
		if exists {
			result++
		}
	}

	return result
}

func moveToTile(position pos, directions []direction) pos {
	// coordinate positioning for hexagonal
	// tiling researched from:
	// https://www.redblobgames.com/grids/hexagon
	for _, d := range directions {
		switch d {
		case northEast:
			if position.Y%2 == 0 {
				position.Y--
			} else {
				position.X++
				position.Y--
			}
		case east:
			position.X++
		case southEast:
			if position.Y%2 == 0 {
				position.Y++
			} else {
				position.X++
				position.Y++
			}
		case southWest:
			if position.Y%2 == 0 {
				position.X--
				position.Y++
			} else {
				position.Y++
			}
		case west:
			position.X--
		case northWest:
			if position.Y%2 == 0 {
				position.X--
				position.Y--
			} else {
				position.Y--
			}
		default:
			// should not get here
			fmt.Println("Error, invalid direction:", d)
			os.Exit(2)
		}
	}

	return position
}

func findBlackTilesAfterANumberofDays(directions [][]direction, days int) int {

	var tilesMap = make(map[pos]rune)

	for _, dList := range directions {
		var tPos = moveToTile(pos{0, 0}, dList)

		_, exists := tilesMap[tPos]
		if exists {
			delete(tilesMap, tPos)
		} else {
			tilesMap[tPos] = 'b'
		}
	}

	for days != 0 {

		var changeMap = make(map[pos]rune)
		for k := range tilesMap {
			var neighbours = getTileNeighbours(k)
			for _, n := range neighbours {
				_, e := tilesMap[n]

				if !e && checkBlackTiles(n, tilesMap) == 2 {
					changeMap[n] = 'b'
				}
			}

			res := checkBlackTiles(k, tilesMap)
			if res == 0 || res > 2 {
				changeMap[k] = 'w'
			}
		}

		for k, v := range changeMap {
			if v == 'w' {
				delete(tilesMap, k)
				continue
			}
			tilesMap[k] = v
		}

		days--
	}

	return len(tilesMap)
}

func findBlackTiles(directions [][]direction) int {
	var tilesMap = make(map[pos]rune)

	for _, dList := range directions {
		var position = moveToTile(pos{0, 0}, dList)

		_, exists := tilesMap[position]
		if exists {
			delete(tilesMap, position)
		} else {
			tilesMap[position] = 'b'
		}
	}

	return len(tilesMap)
}

func main() {
	file, err := os.Open("../../data/day24.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var directions [][]direction

	for scanner.Scan() {
		line := scanner.Text()
		var i = 0

		var tileDirections []direction
		for i < len(line) {
			first := line[i]

			// if this is the last character
			if i+1 == len(line) {
				tileDirections = append(tileDirections, direction(string(first)))
				i++
				continue
			}

			second := line[i+1]

			if (first == 's' || first == 'n') && (second == 'e' || second == 'w') {
				tileDirections = append(tileDirections, direction(string(first)+string(second)))
				i += 2
			} else {
				tileDirections = append(tileDirections, direction(string(first)))
				i++
			}
		}

		directions = append(directions, tileDirections)
	}

	var part1 = findBlackTiles(directions)
	var part2 = findBlackTilesAfterANumberofDays(directions, 100)

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)

	file.Close()
}
