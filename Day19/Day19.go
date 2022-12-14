package main

import (
	"aoc/helpers"
	"fmt"
	"strings"
)

func main() {
	data := helpers.GetLines("Data/Day19.txt")
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var sb strings.Builder
	puzzle := make([][]string, len(data))
	for idx, line := range data {
		puzzle[idx] = strings.Split(line, "")
	}
	rows := len(puzzle)
	cols := len(puzzle[0])
	coords := []int{0, 0}
	for i, char := range puzzle[0] {
		if char == "|" {
			coords[1] = i
		}
	}
	direction, change := findDirection("l", coords, puzzle)
	steps := 0
	for coords[0] >= 0 && coords[0] < rows && coords[1] >= 0 && coords[1] < cols {
		steps += 1
		char := puzzle[coords[0]][coords[1]]
		if char == "+" {
			direction, change = findDirection(direction, coords, puzzle)
		} else if strings.Index(letters, char) != -1 {
			sb.WriteString(char)
		}
		coords = walk(coords, change)
	}
	fmt.Println("Task 01:", sb.String())
	fmt.Println("Task 02:", steps-1)
}

func walk(coords []int, change []int) []int {
	return []int{coords[0] + change[0], coords[1] + change[1]}
}

func findDirection(direction string, coords []int, puzzle [][]string) (string, []int) {
	rows := len(puzzle)
	cols := len(puzzle[0])
	if direction != "d" {
		if coords[0]-1 > 0 {
			temp := puzzle[coords[0]-1][coords[1]]
			if temp != " " && temp != "+" {
				return "u", []int{-1, 0}
			}
		}
	}
	if direction != "u" {
		if coords[0]+1 < rows {
			temp := puzzle[coords[0]+1][coords[1]]
			if temp != " " && temp != "+" {
				return "d", []int{+1, 0}
			}
		}
	}
	if direction != "r" {
		if coords[1]-1 > 0 {
			temp := puzzle[coords[0]][coords[1]-1]
			if temp != " " && temp != "+" {
				return "l", []int{0, -1}
			}
		}
	}
	if direction != "l" {
		if coords[1]+1 < cols {
			temp := puzzle[coords[0]][coords[1]+1]
			if temp != " " && temp != "+" {
				return "r", []int{0, +1}
			}
		}
	}
	return "err", []int{0, 0}
}
