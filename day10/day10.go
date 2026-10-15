package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func findJoltageDifferences(currentRating int, ratings []int, differences *[]int) bool {

	if len(ratings) == 1 {
		*differences = append(*differences, 3)
		return true
	}

	var window []int
	var found = false

	for i := 0; i < len(ratings); i++ {
		diff := ratings[i] - currentRating
		if diff >= 1 && diff <= 3 {
			window = ratings[i:]
			*differences = append(*differences, diff)
			found = true
			break
		}
	}

	if !found {
		return false
	}

	return findJoltageDifferences(window[0], window, differences)
}

func findAllValidAdapterArrangements(ratings []int, idx int, memo map[int]int64) int64 {
	if idx == len(ratings)-1 {
		return 1
	}

	if value, exists := memo[ratings[idx]]; exists {
		return value
	}

	var result = int64(0)
	for next := idx + 1; next < len(ratings); next++ {
		if ratings[next]-ratings[idx] <= 3 {
			result += findAllValidAdapterArrangements(ratings, next, memo)
		} else {
			break
		}
	}

	memo[ratings[idx]] = result
	return result
}

func main() {
	file, err := os.Open("../data/day10.txt")
	if err != nil {
		file, err = os.Open("data/day10.txt")
	}

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var ratings []int
	var differences []int

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		ratings = append(ratings, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input: %s", err)
	}

	file.Close()

	sort.Ints(ratings)
	findJoltageDifferences(0, ratings, &differences)

	var ones = 0
	var threes = 0

	for _, d := range differences {
		switch d {
		case 1:
			ones++
		case 3:
			threes++
		default:
		}
	}

	var part1 = ones * threes

	ratings = append(ratings, 0)
	sort.Ints(ratings)
	ratings = append(ratings, ratings[len(ratings)-1]+3)

	memo := make(map[int]int64)
	var part2 = findAllValidAdapterArrangements(ratings, 0, memo)

	// Old solution
	// var part2 = findAllAdapterArrangements(ratings)

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
