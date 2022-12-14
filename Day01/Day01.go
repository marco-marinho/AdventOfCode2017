package main

import (
	"aoc/helpers"
	"fmt"
)

func main() {
	data := helpers.GetLines("Data/Day01.txt")
	t01 := CalcCaptcha(data, 1)
	t02 := CalcCaptcha(data, len(data[0])/2)
	fmt.Println("Task 01: ", t01)
	fmt.Println("Task 02: ", t02)
}

func CalcCaptcha(data []string, offset int) int {
	for _, entry := range data {
		acc := 0
		for idx, val := range entry {
			var next = (idx + offset) % len(entry)
			if val == int32(entry[next]) {
				acc += int(val) - 48
			}
		}
		return acc
	}
	return 0
}
