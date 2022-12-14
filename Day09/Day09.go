package main

import (
	"aoc/helpers"
	"fmt"
)

func main() {
	data := helpers.GetLines("Data/Day09.txt")
	task1, task2 := getScore(data[0])
	fmt.Println("Task 01: ", task1)
	fmt.Println("Task 02: ", task2)
}

func getScore(input string) (int, int) {
	garbageOpen := false
	skip := false
	garbageChar := 0
	value := 0
	points := 0
	for _, element := range input {
		if skip {
			skip = false
		} else if string(element) == "<" {
			if garbageOpen {
				garbageChar++
			} else {
				garbageOpen = true
			}
		} else if string(element) == "!" && garbageOpen {
			skip = true
		} else if string(element) == ">" {
			garbageOpen = false
		} else if garbageOpen {
			garbageChar++
		} else {
			if string(element) == "{" {
				value++
			}
			if string(element) == "}" {
				points += value
				value--
			}
		}
	}
	return points, garbageChar
}
