package main

import (
	"aoc/helpers"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := helpers.GetLines("Data/Day24.txt")
	components := make([][2]int, len(data))
	for idx, line := range data {
		line = strings.ReplaceAll(line, "\r", "")
		pieces := strings.Split(line, "/")
		first, err := strconv.Atoi(pieces[0])
		second, err2 := strconv.Atoi(pieces[1])
		if err != nil || err2 != nil {
			panic("Parsing error")
		}
		components[idx] = [2]int{first, second}
	}
	solutions := make(chan []int, 10000)
	go dfs(0, []int{}, components, solutions)
	strengths := make([]int, 1000000)
	lengths := make([]int, 1000000)
	idxSol := 0
	for solution := range solutions {
		sum := 0
		for _, idx := range solution {
			sum += components[idx][0]
			sum += components[idx][1]
		}
		strengths[idxSol] = sum
		lengths[idxSol] = len(solution)
		idxSol += 1
	}
	maxStr := helpers.IntSliceArgMax(strengths).Val
	maxLen := helpers.IntSliceArgMax(lengths).Val
	maxLenIdxs := helpers.IntSliceIndexes(lengths, maxLen)
	fmt.Println("Task 01:", maxStr)
	maxStrLen := 0
	for _, idx := range maxLenIdxs {
		if strengths[idx] > maxStrLen {
			maxStrLen = strengths[idx]
		}
	}
	fmt.Println("Task 02:", maxStrLen)
}

func dfs(target int, used []int, components [][2]int, c chan []int) {
	for idx, component := range components {
		if contains(used, idx) {
			continue
		}
		if component[0] == target {
			newUsed := make([]int, len(used)+1)
			copy(newUsed, used)
			newUsed[len(newUsed)-1] = idx
			c <- newUsed
			dfs(component[1], newUsed, components, c)
		} else if component[1] == target {
			newUsed := make([]int, len(used)+1)
			copy(newUsed, used)
			newUsed[len(newUsed)-1] = idx
			c <- newUsed
			dfs(component[0], newUsed, components, c)
		}
	}
	if len(used) == 0 {
		close(c)
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
