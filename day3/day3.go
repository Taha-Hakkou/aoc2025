package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input.txt")
	banks := strings.Split(string(bytes), "\n")
	var sum1, sum2 int
	for _, bank := range banks {
		// Part 1
		var maxi_index, maxj_index int
		for i := 0; i < len(bank)-1; i++ {
			if bank[i] > bank[maxi_index] {
				maxi_index = i
				if bank[i] == '9' {
					break
				}
			}
		}
		maxj_index = maxi_index + 1
		for i := maxi_index + 1; i < len(bank); i++ {
			if bank[i] > bank[maxj_index] {
				maxj_index = i
				if bank[i] == '9' {
					break
				}
			}
		}
		n1 := (bank[maxi_index]-byte('0'))*10 + bank[maxj_index] - byte('0')
		sum1 += int(n1)

		// Part 2
		var indexes []int = []int{}
		for n := range 12 {
			var maxi int
			if len(indexes) > 0 {
				maxi = indexes[len(indexes)-1] + 1
			}
			for i := maxi + 1; i < len(bank)-(12-n-1); i++ {
				if bank[i] > bank[maxi] {
					maxi = i
					if bank[i] == '9' {
						break
					}
				}
			}
			indexes = append(indexes, maxi)
		}
		var n2 int
		for _, index := range indexes {
			n2 *= 10
			n2 += int(bank[index] - byte('0'))
		}
		sum2 += n2
	}
	fmt.Printf("Sum of maximum joltages(2 batteries): %d\n", sum1)
	fmt.Printf("Sum of maximum joltages(12 batteries): %d\n", sum2)
}
