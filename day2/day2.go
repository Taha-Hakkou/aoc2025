package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum int
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ranges := strings.Split(string(bytes), ",")
	for _, rg := range ranges {
		edges := strings.Split(rg, "-")
		start, _ := strconv.Atoi(edges[0])
		end, _ := strconv.Atoi(edges[1])
		for n := start; n <= end; n++ {
			s := strconv.Itoa(n)
			if len(s)%2 == 0 && s[:len(s)/2] == s[len(s)/2:] {
				sum += n
			}
		}
	}
	fmt.Printf("Sum of invalid IDs: %d\n", sum)
}
