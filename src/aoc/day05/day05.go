package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

type seat struct {
	column int
	row    int
}

func calculateRowColumn(reader *bufio.Reader, isUpper bool, data []int) int {

	ch, _, err := reader.ReadRune()
	// reached end of reader
	if err != nil {
		if isUpper {
			return data[len(data)-1]
		}

		return data[0]
	}

	middle := math.Floor(float64(len(data) / 2))

	switch ch {
	case 'L':
		// lower
		fallthrough
	case 'F':
		// lower
		return calculateRowColumn(reader, isUpper, data[:int(middle)])
	case 'B':
		// higher
		fallthrough
	case 'R':
		// higher
		return calculateRowColumn(reader, isUpper, data[int(middle):])
	default:
		// should not get here
		fmt.Println("Error, invalid ch:", ch)
		return -1
	}
}

func main() {
	file, err := os.Open("../../data/day05.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var part1 = -1

	var seatsMap = make(map[seat]int)

	var rows []int
	for i := 0; i < 128; i++ {
		rows = append(rows, i)
	}

	var columns []int
	for i := 0; i < 8; i++ {
		columns = append(columns, i)
	}

	for scanner.Scan() {
		line := scanner.Text()
		row := line[:7]
		column := line[7:]

		var rowIsUpper = false
		if line[len(row)-1] == 'B' {
			rowIsUpper = true
		}

		var columnIsUpper = false
		if line[len(column)-1] == 'R' {
			columnIsUpper = true
		}

		r := calculateRowColumn(bufio.NewReader(strings.NewReader(row)), rowIsUpper, rows)
		c := calculateRowColumn(bufio.NewReader(strings.NewReader(column)), columnIsUpper, columns)

		seatID := (r * 8) + c
		seatsMap[seat{c, r}] = seatID
		if seatID > part1 {
			part1 = seatID
		}

	}

	file.Close()

	var part2 = 0
	var seatsIds []int

	for _, v := range seatsMap {
		seatsIds = append(seatsIds, v)
	}

	sort.Ints(seatsIds)

	for i := 0; i < len(seatsIds)-1; i++ {
		if seatsIds[i+1]-seatsIds[i] == 2 {
			part2 = seatsIds[i] + 1
		}
	}

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
