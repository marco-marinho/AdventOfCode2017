package main

import (
	"aoc/helpers"
	"fmt"
)

func main() {
	results := Tasks()
	fmt.Println("Task 01: ", results[0])
	fmt.Println("Task 02: ", results[1])
}

func Tasks() []int {
	banks := []int{10, 3, 15, 10, 5, 15, 5, 15, 9, 2, 5, 8, 5, 2, 3, 6}
	seen := make(map[string]int)
	curr := helpers.IntSliceToString(banks)
	steps := 0
	ok := false
	var val int
	var diff int
	for !ok {
		steps += 1
		sliceMax := helpers.IntSliceArgMax(banks)
		idx := sliceMax.Arg
		toDistribute := sliceMax.Val
		banks[idx] = 0
		for toDistribute > 0 {
			idx += 1
			banks[idx%len(banks)] += 1
			toDistribute -= 1
		}
		curr = helpers.IntSliceToString(banks)
		val, ok = seen[curr]
		if ok {
			diff = steps - val
		}
		seen[curr] = steps
	}
	return []int{steps, diff}
}
