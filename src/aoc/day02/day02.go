package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../../data/day02.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var part1 = 0
	var part2 = 0

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		policy := strings.Split(split[0], "-")
		character := split[1][0]
		password := split[2]

		lowerBound, _ := strconv.Atoi(policy[0])
		upperBound, _ := strconv.Atoi(policy[1])

		count := strings.Count(password, string(character))
		if count >= lowerBound && count <= upperBound {
			part1++
		}

		if (password[lowerBound-1] == character) != (password[upperBound-1] == character) {
			part2++
		}
	}

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)

	file.Close()
}
