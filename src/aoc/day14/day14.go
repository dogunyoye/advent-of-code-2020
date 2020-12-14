package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../../data/day14.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var maskString = ""

	var memMap = make(map[int]int64)
	var memMapPart2 = make(map[int64]int64)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "mask") {
			maskString = strings.Split(line, " = ")[1]
		} else {
			split := strings.Split(line, " = ")
			var temp = strings.ReplaceAll(split[0], "mem[", "")
			idxString := strings.ReplaceAll(temp, "]", "")

			idx, _ := strconv.Atoi(idxString)
			val, _ := strconv.Atoi(split[1])

			// PART 1 - STARTS
			valueBinary := strconv.FormatInt(int64(val), 2)
			padding := 36 - len(valueBinary)

			var normalised = strings.Repeat("0", padding) + valueBinary

			var masked = ""
			for i, ch := range maskString {
				if ch == 'X' {
					masked += string(normalised[i])
					continue
				}
				masked += string(ch)
			}

			// PART 1 - ENDS
			num, _ := strconv.ParseInt(masked, 2, 64)
			memMap[idx] = num

			// PART 2 - STARTS
			addressBinary := strconv.FormatInt(int64(idx), 2)
			var normalisedAddress = strings.Repeat("0", 36-len(addressBinary)) + addressBinary

			var addressMasked = ""
			var idxsToReplace []int
			for i, ch := range maskString {
				switch ch {
				case 'X':
					// floating
					addressMasked += "X"
					idxsToReplace = append(idxsToReplace, i)
				case '1':
					// set to 1
					addressMasked += "1"
				case '0':
					// unchanged
					addressMasked += string(normalisedAddress[i])
				}
			}

			for i := 0; i < int(math.Pow(float64(2), float64(len(idxsToReplace)))); i++ {
				var config = addressMasked
				var comboBinaryString = strconv.FormatInt(int64(i), 2)
				comboBinaryString = strings.Repeat("0", len(idxsToReplace)-len(comboBinaryString)) + comboBinaryString

				buf := []rune(config)
				for j, idx := range idxsToReplace {
					buf[idx] = rune(comboBinaryString[j])
				}

				// PART 2 - ENDS
				address, _ := strconv.ParseInt(string(buf), 2, 64)
				memMapPart2[address] = int64(val)
			}
		}
	}

	file.Close()

	var part1 = int64(0)
	for _, v := range memMap {
		part1 += v
	}

	var part2 = int64(0)
	for _, v := range memMapPart2 {
		part2 += v
	}

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
