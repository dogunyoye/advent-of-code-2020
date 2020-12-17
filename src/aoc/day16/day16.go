package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bound struct {
	name  string
	lower int
	upper int
}

type ticket struct {
	numbers []int
}

func findInvalidTickets(bounds []bound, nearbyTickets []ticket) ([]int, []int) {
	var invalid []int
	var invalidIdx []int

	for i, t := range nearbyTickets {
		for _, n := range t.numbers {
			var valid = false
			for _, b := range bounds {
				if n >= b.lower && n <= b.upper {
					valid = true
					break
				}
			}

			if !valid {
				invalid = append(invalid, n)
				invalidIdx = append(invalidIdx, i)
				break
			}
		}
	}

	return invalid, invalidIdx
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func determineFields(numberOfFields int, bounds []bound, nearbyTickets []ticket) []int {
	var pendingDetermined = make(map[int][]string)
	var determined = make(map[string]struct{})
	var fieldIdx = 0

	var departureIdxs []int

	for fieldIdx < numberOfFields {
		var candidates = [][]string{}

		for _, t := range nearbyTickets {
			f := t.numbers[fieldIdx]
			var possibleFields []string

			for _, b := range bounds {
				_, exists := determined[b.name]

				if exists {
					continue
				}

				if f >= b.lower && f <= b.upper {
					possibleFields = append(possibleFields, b.name)
				}
			}

			candidates = append(candidates, possibleFields)
		}

		var temp = make(map[string]int)
		for _, c := range candidates {
			for _, t := range c {
				temp[t]++
			}
		}

		var actualCandidates []string
		for k, v := range temp {
			// this key is present in all tickets
			if v == len(nearbyTickets) {
				actualCandidates = append(actualCandidates, k)
			}
		}

		if len(actualCandidates) == 1 {
			// our only option
			determined[actualCandidates[0]] = struct{}{}

			if strings.Contains(actualCandidates[0], "departure") {
				departureIdxs = append(departureIdxs, fieldIdx)
			}

			// eliminate this from the pending map
			for k, v := range pendingDetermined {
				for _, f := range v {
					if f == actualCandidates[0] {
						pendingDetermined[k] = remove(v, f)
					}
				}
			}

			for {
				var existsSingular = false
				var option = ""
				for k, v := range pendingDetermined {
					if len(v) == 1 {
						existsSingular = true
						option = v[0]

						if strings.Contains(option, "departure") {
							departureIdxs = append(departureIdxs, k)
						}
						break
					}
				}

				if !existsSingular {
					break
				}

				for k, v := range pendingDetermined {
					for _, f := range v {
						if f == option {
							pendingDetermined[k] = remove(v, f)
						}
					}
				}

				determined[option] = struct{}{}
			}

		} else {
			pendingDetermined[fieldIdx] = actualCandidates
		}

		fieldIdx++
	}

	return departureIdxs
}

func main() {
	file, err := os.Open("../../data/day16.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	var myTicket = ticket{[]int{}}
	var nearbyTickets []ticket

	var bounds []bound

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) != 0 {
			if strings.Contains(line, "your ticket:") {
				scanner.Scan()
				myTicketLine := strings.Split(scanner.Text(), ",")
				for _, n := range myTicketLine {
					num, _ := strconv.Atoi(n)
					myTicket.numbers = append(myTicket.numbers, num)
				}
			} else if strings.Contains(line, "nearby tickets:") {
				for scanner.Scan() {
					lSplit := strings.Split(scanner.Text(), ",")
					var nearby = ticket{[]int{}}
					for _, n := range lSplit {
						num, _ := strconv.Atoi(n)
						nearby.numbers = append(nearby.numbers, num)
					}
					nearbyTickets = append(nearbyTickets, nearby)
				}
			} else {
				split := strings.Split(line, ": ")
				split2 := strings.Split(split[1], " ")

				left := strings.Split(split2[0], "-")
				right := strings.Split(split2[2], "-")

				ll, _ := strconv.Atoi(left[0])
				lr, _ := strconv.Atoi(left[1])
				rl, _ := strconv.Atoi(right[0])
				rr, _ := strconv.Atoi(right[1])

				bounds = append(bounds, bound{split[0], ll, lr})
				bounds = append(bounds, bound{split[0], rl, rr})
			}
		}
	}

	file.Close()

	var part1 = 0

	invalid, invalidIdx := findInvalidTickets(bounds, nearbyTickets)
	for _, inv := range invalid {
		part1 += inv
	}

	invalidMap := make(map[int]struct{})
	for _, idx := range invalidIdx {
		invalidMap[idx] = struct{}{}
	}

	var validNearbyTickets []ticket

	for i, ticket := range nearbyTickets {
		_, exists := invalidMap[i]
		if !exists {
			validNearbyTickets = append(validNearbyTickets, ticket)
		}
	}

	validNearbyTickets = append(validNearbyTickets, myTicket)
	var idxs = determineFields(len(myTicket.numbers), bounds, validNearbyTickets)

	var part2 = 1
	for _, idx := range idxs {
		part2 *= myTicket.numbers[idx]
	}

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
