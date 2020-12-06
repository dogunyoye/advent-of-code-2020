package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func withinRange(value string, lowerLimit int, upperLimit int) bool {
	num, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	if num < lowerLimit || num > upperLimit {
		return false
	}

	return true
}

func isCredentialsValid(credentialsMap map[string]string) bool {
	for k, v := range credentialsMap {

		switch k {
		case "byr":
			if !withinRange(v, 1920, 2002) {
				return false
			}
		case "iyr":
			if !withinRange(v, 2010, 2020) {
				return false
			}
		case "eyr":
			if !withinRange(v, 2020, 2030) {
				return false
			}
		case "hgt":
			units := v[len(v)-2:]
			switch units {
			case "cm":
				if !withinRange(v[:len(v)-2], 150, 193) {
					return false
				}
			case "in":
				if !withinRange(v[:len(v)-2], 59, 76) {
					return false
				}
			default:
				return false
			}
		case "hcl":
			isValid, _ := regexp.MatchString("#[0-9a-f]{6}", v)

			if !isValid {
				return false
			}
		case "ecl":
			isValid, _ := regexp.MatchString("amb|blu|brn|gry|grn|hzl|oth", v)

			if !isValid {
				return false
			}
		case "pid":
			if len(v) != 9 {
				return false
			}

			_, err := strconv.Atoi(v)
			if err != nil {
				return false
			}

		case "cid":
			// noop
		}
	}

	return true
}

func main() {
	file, err := os.Open("../../data/day04.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var credentialsMap = make(map[string]string)
	var part1 = 0
	var part2 = 0

	var canBreak = false

	scanner.Scan()

	for {
		line := scanner.Text()
		if len(line) != 0 {
			split := strings.Split(line, " ")
			for _, s := range split {
				keyVal := strings.Split(s, ":")
				credentialsMap[keyVal[0]] = keyVal[1]
			}

			if scanner.Scan() {
				continue
			}

			canBreak = true
		}

		// blank line, evaluate credentials stored in map
		_, cidExists := credentialsMap["cid"]
		if len(credentialsMap) == 8 || (len(credentialsMap) == 7 && !cidExists) {
			part1++
			if isCredentialsValid(credentialsMap) {
				part2++
			}
		}

		credentialsMap = make(map[string]string)
		if canBreak {
			break
		}

		scanner.Scan()
	}

	file.Close()

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
