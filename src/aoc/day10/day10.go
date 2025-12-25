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

func tribonacciGen(n int) int64 {
	// function to generate n tetranacci numbers
	switch n {
	case 0, 1:
		return 1
	case 2:
		return 2
	}

	var first = int64(1)
	var second = int64(1)
	var third = int64(2)

	for i := 0; i < n-2; i++ {
		var next = first + second + third
		first = second
		second = third
		third = next
	}

	return third
}

func findAllAdapterArrangements(ratings []int) int64 {
	var steps = make(map[int]int)

	for i := 0; i < len(ratings)-1; i++ {
		curr := ratings[i]
		next := ratings[i+1]

		steps[curr] = next - curr
	}

	// To store the keys in slice in sorted order
	keys := make([]int, len(steps))
	i := 0
	for k := range steps {
		keys[i] = k
		i++
	}
	sort.Ints(keys)

	var onesChain = 0
	var chainArr []int

	for _, i := range keys {
		if steps[i] == 1 {
			onesChain++
		} else {
			if onesChain >= 2 {
				chainArr = append(chainArr, onesChain)
			}

			onesChain = 0
		}
	}

	// append the remaining chain
	if onesChain > 1 {
		chainArr = append(chainArr, onesChain)
	}

	var result = int64(1)
	for _, t := range chainArr {
		result *= tribonacciGen(t)
	}

	return result
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
	file, err := os.Open("../../data/day10.txt")

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
