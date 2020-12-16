package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func playMemoryGame(initialNumbers []int, turnNumber int) int {
	var turnsLeft = turnNumber - len(initialNumbers)
	var memoryGameMap = make(map[int]*[]int)

	var currentTurn = len(initialNumbers)

	// initialise the game
	for i, n := range initialNumbers {
		memoryGameMap[n] = &[]int{i + 1}
	}

	for turnsLeft != 0 {
		currentTurn++
		prev := initialNumbers[len(initialNumbers)-1]
		turnsSpoken, exists := memoryGameMap[prev]
		if exists {
			var spoken = 0
			if len(*turnsSpoken) == 1 {
				spoken = 0
			} else {
				// spoken before
				timesSpoken := len(*turnsSpoken)
				ref := *turnsSpoken
				spoken = ref[timesSpoken-1] - ref[timesSpoken-2]
			}

			initialNumbers = append(initialNumbers, spoken)
			l, e := memoryGameMap[spoken]
			if !e {
				memoryGameMap[spoken] = &[]int{currentTurn}
			} else {
				*l = append(*l, currentTurn)
			}
		}

		turnsLeft--
	}

	return initialNumbers[turnNumber-1]
}

func main() {
	file, err := os.Open("../../data/day15.txt")

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

	file.Close()

	var part1 = playMemoryGame(initialNumbers, 2020)
	var part2 = playMemoryGame(initialNumbers, 30000000)

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
