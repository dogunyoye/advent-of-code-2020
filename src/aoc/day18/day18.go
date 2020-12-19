package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Stack datatype
type Stack []string

// IsEmpty - check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push - a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Peek - retrieve the element at the top of the stack without removing
func (s *Stack) Peek() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(*s) - 1 // Get the index of the top most element.
	element := (*s)[index]
	return element, true
}

// Pop - Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.
	return element, true

}

func evaluate(tokens []string) int {

	var stack Stack

	for _, t := range tokens {
		switch t {
		case "*":
			v1, _ := stack.Pop()
			l, _ := strconv.Atoi(v1)

			v2, _ := stack.Pop()
			r, _ := strconv.Atoi(v2)

			res := strconv.Itoa(l * r)
			stack.Push(res)

		case "+":
			v1, _ := stack.Pop()
			l, _ := strconv.Atoi(v1)

			v2, _ := stack.Pop()
			r, _ := strconv.Atoi(v2)

			res := strconv.Itoa(l + r)
			stack.Push(res)

		default:
			// a number
			stack.Push(t)
		}
	}

	r, _ := stack.Pop()
	result, _ := strconv.Atoi(r)
	return result
}

func shuntingYardAlgorithm(line string) []string {

	// operation stack
	var stack Stack

	// number queue
	var queue []string

	for _, ch := range line {
		switch ch {
		case '*':
			fallthrough
		case '+':
			fallthrough
		case '(':
			stack.Push(string(ch))
		case ')':
			var next, _ = stack.Pop()
			for next != "(" {
				queue = append(queue, next)
				next, _ = stack.Pop()
			}
		default:
			// a number
			queue = append(queue, string(ch))
		}
	}

	for !stack.IsEmpty() {
		val, _ := stack.Pop()
		queue = append(queue, val)
	}

	return queue
}

func shuntingYardAlgorithmWithPrecedence(line string) []string {

	// operation stack
	var stack Stack

	// number queue
	var queue []string

	for _, ch := range line {
		switch ch {
		case '*':
			fallthrough
		case '+':
			head, _ := stack.Peek()
			if head == "+" {
				val, _ := stack.Pop()
				queue = append(queue, val)
			}
			stack.Push(string(ch))
		case '(':
			stack.Push(string(ch))
		case ')':
			var next, _ = stack.Pop()
			for next != "(" {
				queue = append(queue, next)
				next, _ = stack.Pop()
			}
		default:
			// a number
			queue = append(queue, string(ch))
		}
	}

	for !stack.IsEmpty() {
		val, _ := stack.Pop()
		queue = append(queue, val)
	}

	return queue
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	file, err := os.Open("../../data/day18.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var expressions []string

	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), " ", "")
		reversed := reverse(line)

		// put it in reverse!
		buf := []rune(reversed)
		for i, ch := range buf {
			if ch == ')' {
				buf[i] = '('
			}

			if ch == '(' {
				buf[i] = ')'
			}
		}

		expressions = append(expressions, string(buf))
	}

	file.Close()

	var part1 = 0
	for _, e := range expressions {
		part1 += evaluate(shuntingYardAlgorithm(e))
	}

	var part2 = 0
	for _, e := range expressions {
		part2 += evaluate(shuntingYardAlgorithmWithPrecedence(e))
	}

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
