package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input.txt")
	rows := strings.Split(string(bytes), "\n")

	// Part 1
	sum, rows := CountAndRemoveRolls(rows)
	fmt.Printf("Sum of valid rolls of paper: %d\n", sum)

	// Part 2
	var total int = sum
	for sum > 0 {
		sum, rows = CountAndRemoveRolls(rows)
		total += sum
	}
	fmt.Printf("Total of valid rolls of paper: %d\n", total)
}

func CountAndRemoveRolls(rows []string) (int, []string) {
	var sum int
	for i, row := range rows {
		var indexes []int = []int{}
		for j, r := range row {
			if r == '@' {
				var n int // surrounding rolls of paper
				for a := max(i-1, 0); a < min(i+2, len(rows)); a++ {
					for b := max(j-1, 0); b < min(j+2, len(row)); b++ {
						if (a != i || b != j) && rows[a][b] == '@' {
							n++
						}
					}
				}
				if n < 4 {
					sum++
					indexes = append(indexes, j)
				}
			}
		}
		newRow := []byte(row)
		for _, index := range indexes {
			newRow[index] = byte('x')
		}
		rows[i] = string(newRow)
	}
	return sum, rows
}
