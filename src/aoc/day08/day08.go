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

func runTerminatingBootCode(program []instruction, idxs []int, replaced operation) int {

	for i := 0; i < len(idxs); i++ {
		temp := program[idxs[i]].op

		program[idxs[i]].op = replaced
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

	var jmpInstructions []int
	for i, ins := range program {
		if ins.op == jmp {
			jmpInstructions = append(jmpInstructions, i)
		}
	}

	var nopInstructions []int
	for i, ins := range program {
		if ins.op == nop {
			nopInstructions = append(nopInstructions, i)
		}
	}

	replaceJmpRes := runTerminatingBootCode(program, jmpInstructions, nop)
	replaceNopRes := runTerminatingBootCode(program, nopInstructions, jmp)

	// one of replaceJmpRes or replaceNopRes will return -1
	// find the maximum of the 2
	var part2 = math.Max(float64(replaceJmpRes), float64(replaceNopRes))

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
