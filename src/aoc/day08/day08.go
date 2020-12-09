package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type operation string

const (
	acc operation = "acc"
	jmp           = "jmp"
	nop           = "nop"
)

type instruction struct {
	op    operation
	value int
}

func runBootCode(program []instruction) (int, bool) {
	var idx = 0
	var accumulator = 0
	var instructionIdxVisited = make(map[int]struct{})

	for idx < len(program) {
		_, exists := instructionIdxVisited[idx]
		if exists {
			return accumulator, false
		}

		instructionIdxVisited[idx] = struct{}{}
		instruction := program[idx]

		switch instruction.op {
		case acc:
			accumulator += instruction.value
			idx++
		case jmp:
			idx += instruction.value
		case nop:
			idx++
		default:
			// should not get here
			fmt.Println("Error, unknown operation:", instruction.op)
			os.Exit(2)
		}
	}

	return accumulator, true
}

func runTerminatingBootCode(program []instruction, idxs []int) int {

	for i := 0; i < len(idxs); i++ {
		temp := program[idxs[i]].op

		if temp == jmp {
			program[idxs[i]].op = nop
		} else {
			program[idxs[i]].op = jmp
		}

		accumulator, terminated := runBootCode(program)
		if terminated {
			return accumulator
		}

		// undo the change
		program[idxs[i]].op = temp
	}

	return -1
}

func main() {
	file, err := os.Open("../../data/day08.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var program []instruction

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		op := operation(split[0])
		value, _ := strconv.Atoi(split[1])
		program = append(program, instruction{op, value})
	}

	file.Close()

	var part1, _ = runBootCode(program)

	var instructionsToReplace []int
	for i, ins := range program {
		if ins.op == jmp || ins.op == nop {
			instructionsToReplace = append(instructionsToReplace, i)
		}
	}

	var part2 = runTerminatingBootCode(program, instructionsToReplace)

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
