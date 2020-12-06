package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../../data/day06.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var answersMap = make(map[rune]int)
	var people = 0

	var part1 = 0
	var part2 = 0

	var canBreak = false

	scanner.Scan()

	for {
		line := scanner.Text()
		if len(line) != 0 {
			people++

			for _, ch := range line {
				answersMap[ch]++
			}

			if scanner.Scan() {
				continue
			}

			canBreak = true
		}

		// blank line, evaluate answers stored in map
		part1 += len(answersMap)

		for _, v := range answersMap {
			if v == people {
				part2++
			}
		}

		answersMap = make(map[rune]int)
		people = 0

		if canBreak {
			break
		}

		scanner.Scan()
	}

	file.Close()

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
