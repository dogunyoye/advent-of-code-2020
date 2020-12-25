package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func reverseEngineer(publicKey int) int {
	var counter = 0
	var result = 1

	for {
		counter++
		result *= 7
		result = result % 20201227

		if result == publicKey {
			return counter
		}
	}
}

func transform(publicKey int, loops int) int {
	var result = 1

	for i := 0; i < loops; i++ {
		result *= publicKey
		result = result % 20201227
	}

	return result
}

func main() {
	file, err := os.Open("../../data/day25.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	cardPublicKey, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	doorPublicKey, _ := strconv.Atoi(scanner.Text())

	var part1 = transform(doorPublicKey, reverseEngineer(cardPublicKey))

	fmt.Println("Part1:", part1)

	file.Close()
}
