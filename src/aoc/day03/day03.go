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

type traversal struct {
	right int
	down  int
}

func calculateTreesHit(forestMap map[position]string, mapWidth int, mapDepth int, t traversal) int {
	var currentPosition = position{0, 0}
	var treesHit = 0

	for {
		_, exists := forestMap[currentPosition]
		if !exists {
			// out of bounds
			// loop back round the map
			currentPosition.Y %= mapWidth
		}

		object, _ := forestMap[currentPosition]

		if object == "#" {
			treesHit++
		}

		currentPosition.X += t.down
		currentPosition.Y += t.right

		if currentPosition.X > mapDepth {
			return treesHit
		}
	}
}

func main() {
	file, err := os.Open("../../data/day03.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var forestMap = make(map[position]string)
	var width = 0
	var depth = 0

	for scanner.Scan() {
		line := scanner.Text()
		width = len(line)
		for i, c := range line {
			forestMap[position{depth, i}] = string(c)
		}
		depth++
	}

	file.Close()

	var slopeTraversal = traversal{3, 1}
	var part1 = calculateTreesHit(forestMap, width, depth, slopeTraversal)

	fmt.Println("Part1:", part1)

	slopeTraversals :=
		[4]traversal{
			traversal{1, 1},
			traversal{5, 1},
			traversal{7, 1},
			traversal{1, 2},
		}

	var part2 = 1
	for _, t := range slopeTraversals {
		res := calculateTreesHit(forestMap, width, depth, t)
		part2 *= res
	}

	fmt.Println("Part2:", part1*part2)
}
