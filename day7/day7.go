package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(bytes), "\n")
	var index, numOfSplits, numOfTimelines int
	for i, r := range lines[0] {
		if r == 'S' {
			index = i
			break
		}
	}
	if len(lines) > 1 {
		fmt.Println("a")
		numOfSplits, numOfTimelines = splitBeam(index, lines[1:])
	}
	numOfTimelines++
	fmt.Printf("Number of beam splits: %d\n", numOfSplits)
	fmt.Printf("Number of timelines: %d\n", numOfTimelines)
}

func splitBeam(index int, lines []string) (int, int) {
	var splits int
	if lines[0][index] == '^' {
		// count++
		// beams[i] = false
		splits = 1
		if len(lines) > 1 {
			if index > 0 { //&& lines[0][i-1] != '^' {
				n, _ := splitBeam(index-1, lines[1:])
				splits += n
			}
			if index < len(lines[0])-1 {
				n, _ := splitBeam(index+1, lines[1:])
				splits += n
			}
		}
		return splits, 0
	}
	if len(splits) > 1 {
	}
	return splitBeam(index, lines[1:])
}
