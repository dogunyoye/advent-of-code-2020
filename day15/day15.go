package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Plays the memory game, returning the numbers spoken on part1Turn and
// part2Turn. Only the most recent turn of each number is tracked, which is
// all the game rules need.
func playMemoryGame(initialNumbers []int, part1Turn, part2Turn int) (int, int) {
	lastSpoken := make(map[int]int, len(initialNumbers)*2)
	for i, n := range initialNumbers[:len(initialNumbers)-1] {
		lastSpoken[n] = i + 1 // turns are 1-based
	}

	part1 := 0

	last := initialNumbers[len(initialNumbers)-1]
	for turn := len(initialNumbers) + 1; turn <= part2Turn; turn++ {
		if prev, seen := lastSpoken[last]; seen {
			lastSpoken[last] = turn - 1
			last = turn - 1 - prev
		} else {
			lastSpoken[last] = turn - 1
			last = 0
		}

		if turn == part1Turn {
			part1 = last
		}
	}

	return part1, last
}

func main() {
	file, err := os.Open("../data/day15.txt")
	if err != nil {
		file, err = os.Open("data/day15.txt")
	}

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var initialNumbers []int

	scanner.Scan()
	for _, n := range strings.Split(scanner.Text(), ",") {
		number, _ := strconv.Atoi(n)
		initialNumbers = append(initialNumbers, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input: %s", err)
	}

	file.Close()

	part1, part2 := playMemoryGame(initialNumbers, 2020, 30000000)

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
