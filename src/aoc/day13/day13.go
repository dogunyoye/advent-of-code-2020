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
	file, err := os.Open("../../data/day13.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var busIDs []int

	scanner.Scan()
	var timeStamp, _ = strconv.Atoi(scanner.Text())

	scanner.Scan()
	buses := strings.Split(scanner.Text(), ",")

	for _, id := range buses {
		busID, err := strconv.Atoi(id)
		if err == nil {
			busIDs = append(busIDs, busID)
		}
	}

	file.Close()

	var busFound = false
	var timeElapsed = 0
	var id = -1

	for {
		var currentTime = timeStamp + timeElapsed
		for _, b := range busIDs {
			if currentTime%b == 0 {
				busFound = true
				id = b
				break
			}
		}

		if busFound {
			break
		}

		timeElapsed++
	}

	var part1 = timeElapsed * id

	var remainders []int64
	var mods []int64

	for i, id := range buses {
		busID, err := strconv.Atoi(id)
		if err == nil {
			remainders = append(remainders, int64(-i))
			mods = append(mods, int64(busID))
		}
	}

	fmt.Println("Part1:", part1)

	// Part 2 is related to CRT (Chinese Remainder Theorem)
	// Haven't implemented this yet, but used an online calculator
	// https://www.dcode.fr/chinese-remainder with the following values
	// specific to my input:
	// remainder	modulo
	// ---------	------
	//  0			29
	// -19			41
	// -29			601
	// -37			23
	// -42			13
	// -46			17
	// -48			19
	// -60			463
	// -97			37
	fmt.Println("Part2:", remainders, mods)
}
