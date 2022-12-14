package main

import (
	"aoc/helpers"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := "xlqgujun-"
	squares := 0
	grid := make([]string, 128)
	for i := 0; i < 128; i++ {
		buff := data + strconv.Itoa(i)
		hex := helpers.KnotHex(buff)
		bits := hexToBits(hex)
		squares += countOnes(bits)
		grid[i] = bits
	}
	fmt.Println("Task 01:", squares)
	visited := make(map[string]bool)
	groups := 0
	for row := 0; row < 128; row++ {
		for col := 0; col < 128; col++ {
			localCoordinate := []int{row, col}
			localCoordStr := helpers.IntSliceToString(localCoordinate)
			_, ok := visited[localCoordStr]
			if !ok {
				group := getConnected(grid, localCoordinate)
				if len(group) > 0 {
					groups += 1
					for _, element := range group {
						visited[helpers.IntSliceToString(element)] = true
					}
				}
			}
		}
	}
	fmt.Println("Task 02:", groups)
}

func countOnes(input string) int {
	output := 0
	for _, char := range input {
		if string(char) == "1" {
			output += 1
		}
	}
	return output
}

func getConnected(grid []string, coordinate []int) [][]int {
	var connected [][]int
	toCheck := [][]int{coordinate}
	visited := make(map[string]bool)
	for len(toCheck) > 0 {
		localCoordinate := toCheck[0]
		localCoordStr := helpers.IntSliceToString(localCoordinate)
		_, ok := visited[localCoordStr]
		if ok || localCoordinate[0] < 0 || localCoordinate[0] > 127 || localCoordinate[1] < 0 || localCoordinate[1] > 127 {
			toCheck = toCheck[1:]
			continue
		}
		if string(grid[localCoordinate[0]][localCoordinate[1]]) == "1" {
			connected = append(connected, localCoordinate)
			toCheck = append(toCheck, []int{localCoordinate[0] + 1, localCoordinate[1]})
			toCheck = append(toCheck, []int{localCoordinate[0] - 1, localCoordinate[1]})
			toCheck = append(toCheck, []int{localCoordinate[0], localCoordinate[1] + 1})
			toCheck = append(toCheck, []int{localCoordinate[0], localCoordinate[1] - 1})
		}
		visited[localCoordStr] = true
	}
	return connected
}

func hexToBits(input string) string {
	table := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"a": "1010",
		"b": "1011",
		"c": "1100",
		"d": "1101",
		"e": "1110",
		"f": "1111"}
	var sb strings.Builder
	for _, letter := range input {
		sb.WriteString(table[string(letter)])
	}
	return sb.String()
}
