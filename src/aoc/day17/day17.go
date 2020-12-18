package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type state string

const (
	inactive state = "."
	active         = "#"
)

type pos struct {
	X int
	Y int
	Z int
	W int
}

func checkActiveCubes(cube pos, pocketDimension map[pos]state, isPart2 bool) int {

	var result = 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {

				if isPart2 {
					for w := -1; w <= 1; w++ {
						if x == 0 && y == 0 && z == 0 && w == 0 {
							continue
						}
						_, e := pocketDimension[pos{cube.X + x, cube.Y + y, cube.Z + z, cube.W + w}]
						if e {
							result++
						}

					}
				} else {
					if x == 0 && y == 0 && z == 0 {
						continue
					}
					_, e := pocketDimension[pos{cube.X + x, cube.Y + y, cube.Z + z, 0}]
					if e {
						result++
					}
				}
			}
		}
	}

	return result
}

func getCubeNeighbours(cube pos, isPart2 bool) []pos {

	var result []pos
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {

				if isPart2 {
					for w := -1; w <= 1; w++ {
						if x == 0 && y == 0 && z == 0 && w == 0 {
							continue
						}
						result = append(result, pos{cube.X + x, cube.Y + y, cube.Z + z, cube.W + w})
					}
				} else {
					if x == 0 && y == 0 && z == 0 {
						continue
					}
					result = append(result, pos{cube.X + x, cube.Y + y, cube.Z + z, 0})
				}
			}
		}
	}

	return result
}

func findActiveCubes(pocketDimension map[pos]state, cycles int, isPart2 bool) int {
	var c = cycles

	for c > 0 {
		var changeMap = make(map[pos]state)

		for k := range pocketDimension {
			var neighbours = getCubeNeighbours(k, isPart2)
			for _, cube := range neighbours {
				_, e := pocketDimension[cube]

				if !e && checkActiveCubes(cube, pocketDimension, isPart2) == 3 {
					changeMap[cube] = active
				}
			}

			res := checkActiveCubes(k, pocketDimension, isPart2)
			if res != 2 && res != 3 {
				changeMap[k] = inactive
			}
		}

		// apply changes
		for k, v := range changeMap {
			if v == inactive {
				delete(pocketDimension, k)
				continue
			}
			pocketDimension[k] = v
		}

		c--
	}

	return len(pocketDimension)
}

func main() {
	file, err := os.Open("../../data/day17.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var depth = 0
	var pocketDimension = make(map[pos]state)
	var pocketDimension2 = make(map[pos]state)

	for scanner.Scan() {
		line := scanner.Text()
		for i, ch := range line {
			var cube = pos{depth, i, 0, 0}
			if ch == '#' {
				pocketDimension[cube] = state(string(ch))
				pocketDimension2[cube] = state(string(ch))
			}
		}
		depth++
	}

	var part1 = findActiveCubes(pocketDimension, 6, false)
	var part2 = findActiveCubes(pocketDimension2, 6, true)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)

	file.Close()
}
