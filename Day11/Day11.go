package main

import (
	"aoc/helpers"
	"fmt"
	"strings"
)

type Coord struct {
	q int
	r int
	s int
}

func main() {
	coord := Coord{0, 0, 0}
	data := strings.Split(helpers.GetLines("Data/Day11.txt")[0], ",")
	distances := make([]int, len(data))
	for idx, step := range data {
		walk(&coord, step)
		distances[idx] = distance(coord)
	}
	fmt.Println("Task 01:", distances[len(distances)-1])
	max := helpers.IntSliceArgMax(distances).Val
	fmt.Println("Task 02:", max)
}

func walk(coordinate *Coord, direction string) {
	if direction == "n" {
		coordinate.q -= 1
		coordinate.s += 1
	} else if direction == "ne" {
		coordinate.q -= 1
		coordinate.r += 1
	} else if direction == "nw" {
		coordinate.s += 1
		coordinate.r -= 1
	} else if direction == "s" {
		coordinate.q += 1
		coordinate.s -= 1
	} else if direction == "se" {
		coordinate.r += 1
		coordinate.s -= 1
	} else if direction == "sw" {
		coordinate.q += 1
		coordinate.r -= 1
	}
}

func distance(coordinate Coord) int {
	return (helpers.AbsInt(coordinate.q) + helpers.AbsInt(coordinate.r) + helpers.AbsInt(coordinate.s)) / 2
}
