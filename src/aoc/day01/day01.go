package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type void struct{}

func main() {
	file, err := os.Open("../../data/day01.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var entries = make(map[int]void)

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		result := 2020 - i
		_, exists := entries[result]
		if exists {
			fmt.Println("Part1:", i*result)
		}

		entries[i] = struct{}{}
	}

	for key := range entries {
		var found = false
		toFind := 2020 - key
		for e := range entries {
			if e == key {
				continue
			}

			result := toFind - e
			_, exists := entries[result]
			if exists {
				fmt.Println("Part2:", key*e*result)
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	file.Close()
}
