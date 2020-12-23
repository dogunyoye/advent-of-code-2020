package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"log"
	"os"
	"strconv"
)

func contains(s map[int]struct{}, r int) bool {
	_, exists := s[r]
	return exists
}

func remove(s []int, r int) []int {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func getIndex(s []int, r int) int {
	for i, v := range s {
		if v == r {
			return i
		}
	}

	// this would be bad
	return -1
}

func insert(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func shiftLeft(a *[]int, i int) {
	x, b := (*a)[:i], (*a)[i:]
	*a = append(b, x...)
}

func hasDestination(cups *ring.Ring, destination int) bool {
	var found = false
	cups.Do(func(p interface{}) {
		if p.(int) == destination {
			found = true
		}
	})

	return found
}

func playCrabCupsOptimised(cups *ring.Ring, memo map[int]*ring.Ring, moves int) {

	for moves != 0 {
		removed := cups.Unlink(3)
		var destination = cups.Value.(int) - 1

		for destination != 0 && hasDestination(removed, destination) {
			destination--
		}

		if destination == 0 {
			destination = 1000000
		}

		memo[destination].Link(removed)
		cups = cups.Next()

		moves--
	}
}

func playCrabCups(cups []int, moves int, cupsSize int) []int {
	var counter = 0

	for moves != 0 {
		var idx = counter % cupsSize

		var cup = cups[idx]

		var pickup []int

		var pickupMap = make(map[int]struct{})
		var cupsMap = make(map[int]int)

		var wraparound = 0

		for i := idx + 1; i <= idx+3; i++ {
			pickup = append(pickup, cups[i%cupsSize])
			pickupMap[cups[i%cupsSize]] = struct{}{}

			if i >= 9 {
				wraparound++
			}
		}

		for _, p := range pickup {
			cups = remove(cups, p)
		}

		var max = -1
		for i, c := range cups {
			cupsMap[c] = i
			if c > max {
				max = c
			}
		}

		var destination = cup - 1
		if contains(pickupMap, destination) {

			destination--
			for contains(pickupMap, destination) && destination != 0 {
				destination--
			}

			// couldn't find a number
			// set it to highest value
			if destination == 0 {
				destination = max
			}

		} else {
			_, exists := cupsMap[destination]
			if !exists {
				destination = max
			}
		}

		var j = 0

		var destIdx = cupsMap[destination]
		var cupIdx = cupsMap[cup]

		for i := destIdx + 1; i <= destIdx+3; i++ {
			cups = insert(cups, i, pickup[j])
			j++
		}

		if destIdx < cupIdx {
			shiftLeft(&cups, 3-wraparound)
		}

		moves--
		counter++
	}

	return cups
}

func main() {
	file, err := os.Open("../../data/day23.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var cups []int
	var cupsPart2 []int

	for scanner.Scan() {
		line := scanner.Text()
		for _, ch := range line {
			num, _ := strconv.Atoi(string(ch))
			cups = append(cups, num)
			cupsPart2 = append(cupsPart2, num)
		}
	}

	file.Close()

	var part1 = ""
	cups = playCrabCups(cups, 100, 9)

	var idx = getIndex(cups, 1)
	for i := idx + 1; i < idx+9; i++ {
		part1 += strconv.Itoa(cups[i%len(cups)])
	}

	for i := 10; i <= 1000000; i++ {
		cupsPart2 = append(cupsPart2, i)
	}

	var cupsRing = ring.New(1000000)
	var memo = make(map[int]*ring.Ring)

	for _, c := range cupsPart2 {
		cupsRing.Value = c
		memo[c] = cupsRing
		cupsRing = cupsRing.Next()
	}

	playCrabCupsOptimised(cupsRing, memo, 10000000)

	var curr = memo[1].Next()
	var firstNum = curr.Value.(int)

	curr = curr.Next()
	var secondNum = curr.Value.(int)

	var part2 = firstNum * secondNum

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
