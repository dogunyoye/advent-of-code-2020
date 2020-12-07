package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type contents struct {
	quantity  int
	bagColour string
}

func calculateShinyGoldBags(rulesMap map[string][]contents, bagToCheck string) bool {
	contents, _ := rulesMap[bagToCheck]
	var canHold = false

	for _, b := range contents {
		if b.bagColour == "shiny gold" {
			canHold = true
		} else {
			canHold = calculateShinyGoldBags(rulesMap, b.bagColour)
		}

		if canHold {
			return canHold
		}
	}

	return canHold
}

func calculateNumberofBags(rulesMap map[string][]contents, bagToCheck string, count *int) {
	contents, _ := rulesMap[bagToCheck]

	for _, b := range contents {
		for i := 0; i < b.quantity; i++ {
			*count++
			calculateNumberofBags(rulesMap, b.bagColour, count)
		}
	}
}

func main() {
	file, err := os.Open("../../data/day07.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var part1 = 0
	var rulesMap = make(map[string][]contents)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " contain ")
		colour := strings.ReplaceAll(split[0], " bags", "")

		var bagArr []contents

		if split[1] == "no other bags." {
			rulesMap[colour] = bagArr
			continue
		}

		bagContents := strings.Split(split[1], ", ")
		for _, b := range bagContents {
			bSplit := strings.Split(b, " ")
			bQuantity, err := strconv.Atoi(bSplit[0])
			if err != nil {
				fmt.Println("Error parsing bag quantity:", err)
				os.Exit(2)
			}
			bColour := bSplit[1] + " " + bSplit[2]
			bagArr = append(bagArr, contents{bQuantity, bColour})
		}

		rulesMap[colour] = bagArr
	}

	file.Close()
	var bagToCheck = "shiny gold"

	for k := range rulesMap {
		if k != bagToCheck {
			res := calculateShinyGoldBags(rulesMap, k)
			if res {
				part1++
			}
		}
	}

	var part2 = 0
	calculateNumberofBags(rulesMap, bagToCheck, &part2)

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
