package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type position struct {
	X int
	Y int
}

type state string
type direction int

const (
	floor    state = "."
	empty          = "L"
	occupied       = "#"
)

const (
	north direction = iota
	northEast
	east
	southEast
	south
	southWest
	west
	northWest
)

func findVisibleOccupiedSeat(seatsMap map[position]state, pos position, dir direction) int {

	var nextDir = pos

	switch dir {
	case north:
		nextDir = position{pos.X - 1, pos.Y}
	case northEast:
		nextDir = position{pos.X - 1, pos.Y + 1}
	case east:
		nextDir = position{pos.X, pos.Y + 1}
	case southEast:
		nextDir = position{pos.X + 1, pos.Y + 1}
	case south:
		nextDir = position{pos.X + 1, pos.Y}
	case southWest:
		nextDir = position{pos.X + 1, pos.Y - 1}
	case west:
		nextDir = position{pos.X, pos.Y - 1}
	case northWest:
		nextDir = position{pos.X - 1, pos.Y - 1}
	default:
		// should not get here
		fmt.Println("Error invalid direction:", dir)
		os.Exit(2)
	}

	v, exists := seatsMap[nextDir]
	if !exists {
		// out of bounds
		return 0
	}

	if v == occupied || v == empty {
		if v == occupied {
			return 1
		}

		return 0
	}

	return findVisibleOccupiedSeat(seatsMap, nextDir, dir)
}

func findSeatsAtNoChanges(seatsMap map[position]state, isPartTwo bool) int {
	var result = 0
	var changesMap = make(map[position]state)

	for {

		for k, v := range seatsMap {
			adjacents := [8]state{
				// north west
				seatsMap[position{k.X - 1, k.Y - 1}],
				// north
				seatsMap[position{k.X - 1, k.Y}],
				// north east
				seatsMap[position{k.X - 1, k.Y + 1}],
				// east
				seatsMap[position{k.X, k.Y + 1}],
				// south east
				seatsMap[position{k.X + 1, k.Y + 1}],
				// south
				seatsMap[position{k.X + 1, k.Y}],
				// south west
				seatsMap[position{k.X + 1, k.Y - 1}],
				// west
				seatsMap[position{k.X, k.Y - 1}],
			}

			switch v {
			case empty:

				var foundOccupied = false
				if !isPartTwo {
					for _, s := range adjacents {
						if s == occupied {
							foundOccupied = true
							break
						}
					}
				} else {
					var seatsOccupied = 0
					for dir := direction(0); dir < 8; dir++ {
						seatsOccupied += findVisibleOccupiedSeat(seatsMap, k, dir)
						if seatsOccupied != 0 {
							foundOccupied = true
							break
						}
					}
				}

				if !foundOccupied {
					changesMap[position{k.X, k.Y}] = occupied
				}

			case occupied:

				if !isPartTwo {
					var count = 0
					for _, s := range adjacents {
						if s == occupied {
							count++
						}
					}

					if count >= 4 {
						changesMap[position{k.X, k.Y}] = empty
					}

				} else {
					var seatsOccupied = 0
					for dir := direction(0); dir < 8; dir++ {
						seatsOccupied += findVisibleOccupiedSeat(seatsMap, k, dir)
					}

					if seatsOccupied >= 5 {
						changesMap[position{k.X, k.Y}] = empty
					}
				}

			case floor:
				// noop

			default:
				// should not get here
				fmt.Println("Error unknown state:", v)
				os.Exit(2)
			}
		}

		// no changes = equilibrium
		if len(changesMap) == 0 {
			break
		}

		for k, v := range changesMap {
			seatsMap[position{k.X, k.Y}] = v
		}

		changesMap = make(map[position]state)
	}

	for _, v := range seatsMap {
		if v == occupied {
			result++
		}
	}

	return result
}

func main() {
	file, err := os.Open("../../data/day11.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var seatsMap = make(map[position]state)
	var seatsMapPart2 = make(map[position]state)
	var depth = 0

	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			seatsMap[position{depth, i}] = state(c)
			seatsMapPart2[position{depth, i}] = state(c)
		}
		depth++
	}

	file.Close()

	part1 := findSeatsAtNoChanges(seatsMap, false)
	part2 := findSeatsAtNoChanges(seatsMapPart2, true)

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
