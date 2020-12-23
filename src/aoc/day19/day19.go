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

func containsDigit(s string) bool {
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			return true
		}
	}

	return false
}

func remove(s []int, r int) []int {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func parseRules(rulesMap map[int]string, options map[int][]int) {
	for {
		var count = 0
		for _, v := range rulesMap {
			if containsDigit(v) {
				// more rule parsing to do..
				count++
			}
		}

		if count == 0 {
			// all rules have been parsed!
			// get out
			break
		}

		for k, v := range rulesMap {
			if containsDigit(v) {
				// needs replacing
				idxs, _ := options[k]

				var optionsDefined = true
				for _, idx := range idxs {
					t, _ := rulesMap[idx]
					if containsDigit(t) {
						optionsDefined = false
						break
					}
				}

				// out of options, can now build
				// string
				if optionsDefined {
					val, _ := rulesMap[k]
					var regex = ""

					split := strings.Split(val, " ")
					for _, x := range split {
						if x == "|" {
							regex += x
							continue
						}

						num, _ := strconv.Atoi(x)
						dep, _ := rulesMap[num]
						regex += dep
					}

					regex = "(" + regex + ")"
					rulesMap[k] = regex
				}
			}
		}
	}
}

func main() {
	file, err := os.Open("../../data/day19.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var rulesMap = make(map[int]string)
	var rulesMap2 = make(map[int]string)

	var options = make(map[int][]int)
	var options2 = make(map[int][]int)

	var messages []string

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 0 {
			// initially parsing the rules
			split := strings.Split(line, ": ")
			k, _ := strconv.Atoi(split[0])
			var v = split[1]

			if containsDigit(v) {
				idxs := strings.Split(v, " ")
				var l = []int{}
				for _, key := range idxs {
					if key == "|" {
						continue
					}

					knum, _ := strconv.Atoi(key)
					l = append(l, knum)
				}

				options[k] = l
				options2[k] = l
			} else {
				v = strings.ReplaceAll(v, "\"", "")
			}

			rulesMap[k] = v
			rulesMap2[k] = v
			continue
		}

		// empty line, start parsing messages
		for scanner.Scan() {
			message := scanner.Text()
			messages = append(messages, message)
		}
	}

	file.Close()

	parseRules(rulesMap, options)

	var part1 = 0
	var regex = strings.ReplaceAll("^"+rulesMap[0]+"$", " ", "")
	for _, m := range messages {
		isValid, _ := regexp.MatchString(regex, m)
		if isValid {
			part1++
		}
	}

	// Part2: replace 8: 42, 11: 42 31
	// with
	// 8: 42 | 42 8
	// 11: 42 31 | 42 11 31
	// super hacky...
	rulesMap[8] = "42 | 42 42 | 42 42 42 | 42 42 42 42 | 42 42 42 42 42 | 42 42 42 42 42 42 | 42 42 42 42 42 42 42 | 42 42 42 42 42 42 42 42 | 42 42 42 42 42 42 42 42 42 | 42 42 42 42 42 42 42 42 42 42"
	rulesMap[11] = "42 31 | 42 42 31 31 | 42 42 42 31 31 31 | 42 42 42 42 31 31 31 31 | 42 42 42 42 42 31 31 31 31 31 | 42 42 42 42 42 42 31 31 31 31 31 31 | 42 42 42 42 42 42 42 31 31 31 31 31 31 31"

	parseRules(rulesMap, options)

	var part2 = 0
	var regex2 = strings.ReplaceAll("^"+rulesMap[8]+rulesMap[11]+"$", " ", "")
	var r = regexp.MustCompile(regex2)

	for _, m := range messages {
		isValid := r.MatchString(m)
		if isValid {
			part2++
		}
	}

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
