package main

import (
	"aoc/helpers"
	"fmt"
)

func main() {
	fmt.Println("Task 01: ", Tasks(1))
	fmt.Println("Task 02: ", Tasks(2))
}

func Tasks(task int) int {
	data := helpers.GetLines("Data/Day05.txt")
	jumps := make([]int, len(data))
	for idx, value := range data {
		jumps[idx] = helpers.StringToInts(value, " ")[0]
	}
	idx := 0
	steps := 0
	for 0 <= idx && idx < len(jumps) {
		var diff int
		steps += 1
		if task == 1 {
			diff = 1
		} else {
			if jumps[idx] >= 3 {
				diff = -1
			} else {
				diff = 1
			}
		}
		jumps[idx] += diff
		idx += jumps[idx] - diff
	}
	return steps
}
