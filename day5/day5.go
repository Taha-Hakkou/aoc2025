package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start, End int
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	input := strings.Split(string(bytes), "\n\n")
	ranges := strings.Split(input[0], "\n")
	ids := strings.Split(input[1], "\n")

	// Part 1
	var sum1 int
	for _, id := range ids {
		n, _ := strconv.Atoi(id)
		for _, rg := range ranges {
			bds := strings.Split(rg, "-")
			start, _ := strconv.Atoi(bds[0])
			end, _ := strconv.Atoi(bds[1])
			if n >= start && n <= end {
				sum1++
				break
			}
		}
	}

	// Part 2
	var sum2 int
	rangeObjs := make([]Range, 0)
	for _, rg := range ranges {
		bds := strings.Split(rg, "-")
		start, _ := strconv.Atoi(bds[0])
		end, _ := strconv.Atoi(bds[1])
		rangeObj := Range{
			Start: start,
			End:   end,
		}
		var i int
		for ; i < len(rangeObjs); i++ {
			obj := rangeObjs[i]
			if rangeObj.Start <= obj.End+1 && rangeObj.End >= obj.Start-1 {
				obj.Start = min(obj.Start, rangeObj.Start)
				obj.End = max(obj.End, rangeObj.End)
				break
			}
		}
		if i == len(rangeObjs) {
			rangeObjs = append(rangeObjs, rangeObj)
		}
	}
	/*for i := 1; i < len(rangeObjs); i++ {
		obj := rangeObjs[i]
		for j := 0; j < i; j++ {
			o := rangeObjs[j]
			if obj.Start <= o.End+1 && obj.End >= o.Start-1 {
				o.Start = min(o.Start, obj.Start)
				o.End = max(o.End, obj.End)
				//
				newRangeObjs := make([]Range, len(rangeObjs)-1)
				copy(newRangeObjs, rangeObjs[:i])
				for k, rg := range rangeObjs[i+1:] {
					newRangeObjs[i+k] = rg
				}
				rangeObjs = newRangeObjs
				break
			}
		}
	}*/
	for _, rg := range rangeObjs {
		sum2 += rg.End - rg.Start + 1
	}

	fmt.Printf("Sum of available fresh IDs: %d\n", sum1)
	fmt.Printf("Sum of fresh IDs: %d\n", sum2)
}
