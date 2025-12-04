package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum1, sum2 int
	bytes, _ := os.ReadFile("input.txt")
	ranges := strings.Split(string(bytes), ",")
	for _, rg := range ranges {
		// fmt.Println(rg)
		edges := strings.Split(rg, "-")
		start, _ := strconv.Atoi(edges[0])
		end, _ := strconv.Atoi(edges[1])
		for n := start; n <= end; n++ {
			s := strconv.Itoa(n)
			for i := 2; i < len(s)-1; i++ { // not working properly
				if len(s)%i == 0 {
					var j int = 1
					for ; j < i; j++ {
						if s[:len(s)/i] != s[len(s)/i*j:len(s)/i*(j+1)] {
							break
						}
					}
					if j == len(s)/i {
						if i == 2 {
							sum1 += n
						}
						sum2 += n
						break
					}
				}
			}

		}
	}
	fmt.Printf("Sum of invalid IDs(Repeated twice): %d\n", sum1)
	fmt.Printf("Sum of invalid IDs(Repeated at least twice): %d\n", sum2)
}
