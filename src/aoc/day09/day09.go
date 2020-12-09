package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func findBreakingNumber(numbers []int, preamble int) int {
	var idx = 0
	var idxToCheck = preamble

	for {
		var found = false
		val := numbers[idxToCheck]
		numbersToCheck := numbers[idx : idx+preamble]
		for i := 0; i < len(numbersToCheck)-1; i++ {
			for j := i + 1; j < len(numbersToCheck); j++ {
				if numbersToCheck[i]+numbersToCheck[j] == val {
					found = true
				}
			}
		}

		if !found {
			return numbers[idxToCheck]
		}

		idx++
		idxToCheck++
	}
}

func findContiguousSetOfNumbers(numbers []int, toFind int) []int {
	var idx = 0
	var res []int

	var currentTotal = 0

	for {
		var currIdx = idx
		for currentTotal < toFind {
			currentTotal += numbers[currIdx]
			res = append(res, numbers[currIdx])

			currIdx++
		}

		if currentTotal == toFind {
			return res
		}

		res = []int{}
		currentTotal = 0
		idx++
	}

}

func main() {
	file, err := os.Open("../../data/day09.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var numbers []int

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, value)
	}

	file.Close()

	preamble := 25
	var part1 = findBreakingNumber(numbers, preamble)

	res := findContiguousSetOfNumbers(numbers, part1)
	sort.Ints(res)

	var part2 = res[0] + res[len(res)-1]

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
