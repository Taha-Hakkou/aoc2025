package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(bytes), "\n")
	var password1 int = 0
	var password2 int = 0
	var count int = 50
	var nextcount int
	for _, line := range lines {
		if len(line) < 2 {
			continue
		}
		n, _ := strconv.Atoi(line[1:])
		switch line[0] {
		case 'R':
			nextcount = count + n
			password2 += nextcount / 100
		case 'L':
			nextcount = count - n
			password2 -= nextcount / 100
			if count > 0 && nextcount <= 0 {
				password2++
			}
		}
		count = nextcount % 100
		if count < 0 {
			count += 100
		}
		if count == 0 {
			password1++
		}
	}
	fmt.Println("part1's password: " + strconv.Itoa(password1))
	fmt.Println("part2's password: " + strconv.Itoa(password2))
	fmt.Println(-523 / 100)
}
