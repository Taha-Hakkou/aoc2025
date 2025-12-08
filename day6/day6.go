package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(bytes), "\n")

	// operators
	operators := strings.Fields(lines[len(lines)-1])

	// Part 1
	var sum1 int
	operands := make([]int, len(strings.Fields(lines[0])))
	for _, line := range lines[:len(lines)-1] {
		elems := strings.Fields(line)
		for i, e := range elems {
			n, _ := strconv.Atoi(e)
			switch operators[i] {
			case "+":
				operands[i] += n
			case "*":
				operands[i] *= n
			}
		}
	}
	for _, operand := range operands {
		sum1 += operand
	}

	// Part 2
	var sum2 int
	problems := make([][]string, len(operators))
	for i, line := range lines[:len(lines)-1] {
		numbers := strings.Fields(line)
		for j, number := range numbers {
			if i == 0 {
				problems[j] = make([]string, 1)
			}
			for k, r := range number {
				if k >= len(problems[j]) {
					problems[j] = append(problems[j], "")
				}
				problems[j][k] += string(r) // index range problem
			}
		}
	}
	for _, problem := range problems {
		var n int
		for _, number := range problem {
			a, _ := strconv.Atoi(number)
			n += a
		}
		sum2 += n
	}

	// Result

	fmt.Printf("Result of math problems: %d\n", sum1)
	fmt.Printf("Result of math problems: %d\n", sum2)
}
