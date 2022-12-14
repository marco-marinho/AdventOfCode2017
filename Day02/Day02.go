package main

import (
	"aoc/helpers"
	"fmt"
	"math"
)

func main() {
	data := helpers.GetLines("Data/Day02.txt")
	t1 := Task01(data)
	t2 := Task02(data)
	fmt.Println("Task 01:", t1)
	fmt.Println("Task 02:", t2)
}

func Task01(data []string) int {
	acc := 0
	for _, line := range data {
		nums := helpers.StringToInts(line, " ")
		min := math.MaxInt32
		max := 0
		for idx := 0; idx < len(nums); idx++ {
			if nums[idx] < min {
				min = nums[idx]
			}
			if nums[idx] > max {
				max = nums[idx]
			}
		}
		acc += max - min
	}
	return acc
}

func Task02(data []string) int {
	acc := 0
Line:
	for _, line := range data {
		nums := helpers.StringToInts(line, " ")
		for idxO := 0; idxO < len(nums); idxO++ {
			for idxI := idxO + 1; idxI < len(nums); idxI++ {
				var numerator int
				var denominator int
				if nums[idxO] >= nums[idxI] {
					numerator = nums[idxO]
					denominator = nums[idxI]
				} else {
					numerator = nums[idxI]
					denominator = nums[idxO]
				}
				if numerator%denominator == 0 {
					acc += numerator / denominator
					continue Line
				}
			}
		}
	}
	return acc
}
