package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type action string

const (
	// move
	north action = "N"
	east         = "E"
	south        = "S"
	west         = "W"

	// rotate
	left  = "L"
	right = "R"

	// move in facing direction
	forward = "F"
)

type instruction struct {
	op    action
	value int
}

type ship struct {
	X         int
	Y         int
	facing    action
	movements map[action]int
}

func rotateShip(facing action, dir action, degrees int) action {
	var directions = [4]action{
		north,
		east,
		south,
		west,
	}

	var idx = 0
	for i, d := range directions {
		if d == facing {
			idx = i
			break
		}
	}

	var turns = degrees / 90
	if dir == left {
		turns *= -1
	}

	idx += turns
	idx = int(math.Abs(float64(4+idx))) % 4

	return directions[idx]
}

func navigateShip(s ship, navigation []instruction) int {

	for _, ins := range navigation {
		switch ins.op {
		case north:
			s.movements[north] += ins.value
		case east:
			s.movements[east] += ins.value
		case south:
			s.movements[south] += ins.value
		case west:
			s.movements[west] += ins.value
		case left:
			fallthrough
		case right:
			s.facing = rotateShip(s.facing, ins.op, ins.value)
		case forward:
			s.movements[s.facing] += ins.value
		default:
			// should not get here
			fmt.Println("Error invalid action:", ins.op)
			os.Exit(2)
		}
	}

	y := math.Abs(float64(s.movements[north]) - float64(s.movements[south]))
	x := math.Abs(float64(s.movements[east]) - float64(s.movements[west]))

	return int(x + y)
}

func abs(val int) int {
	if val >= 0 {
		return val
	}

	return -val
}

func rotateWayPointClockwise(x int, y int, radians float64) (int, int) {
	var wayPointX = (x * int(math.Cos(radians))) + (y * int(math.Sin(radians)))
	var wayPointY = -(x * int(math.Sin(radians))) + (y * int(math.Cos(radians)))
	return wayPointX, wayPointY
}

func rotateWayPointAntiClockwise(x int, y int, radians float64) (int, int) {
	var wayPointX = (x * int(math.Cos(radians))) - (y * int(math.Sin(radians)))
	var wayPointY = (x * int(math.Sin(radians))) + (y * int(math.Cos(radians)))
	return wayPointX, wayPointY
}

func navigateShipToWayPoint(s ship, navigation []instruction) int {
	var wayPointX = 10
	var wayPointY = 1

	for _, ins := range navigation {
		switch ins.op {
		case north:
			wayPointY += ins.value
		case east:
			wayPointX += ins.value
		case south:
			wayPointY -= ins.value
		case west:
			wayPointX -= ins.value
		case left:
			radians := float64(ins.value) * (math.Pi / 180)
			wayPointX, wayPointY = rotateWayPointAntiClockwise(wayPointX, wayPointY, radians)
		case right:
			radians := float64(ins.value) * (math.Pi / 180)
			wayPointX, wayPointY = rotateWayPointClockwise(wayPointX, wayPointY, radians)
		case forward:
			s.X += wayPointX * ins.value
			s.Y += wayPointY * ins.value
		default:
			// should not get here
			fmt.Println("Error invalid action:", ins.op)
			os.Exit(2)
		}
	}

	return abs(s.X) + abs(s.Y)
}

func main() {
	file, err := os.Open("../../data/day12.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var navigation []instruction

	for scanner.Scan() {
		line := scanner.Text()
		op := action(line[0])
		value, _ := strconv.Atoi(line[1:])
		navigation = append(navigation, instruction{op, value})
	}

	file.Close()

	var ferry = ship{0, 0, east, map[action]int{
		north: 0,
		east:  0,
		south: 0,
		west:  0,
	}}
	var part1 = navigateShip(ferry, navigation)

	var ferry2 = ship{0, 0, east, map[action]int{
		north: 0,
		east:  0,
		south: 0,
		west:  0,
	}}
	var part2 = navigateShipToWayPoint(ferry2, navigation)

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
