package helpers

import (
	"math"
	"os"
	"strconv"
	"strings"
)

type IntMax struct {
	Arg int
	Val int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(path string) string {
	data, err := os.ReadFile(path)
	check(err)
	return string(data)
}

func GetLines(path string) []string {
	data := ReadFile(path)
	return strings.Split(data, "\n")
}

func StringToInts(input string, sep string) []int {
	input = strings.ReplaceAll(input, "	", " ")
	split := strings.Split(input, sep)
	output := make([]int, 0)
	for _, entry := range split {
		buff, err := strconv.Atoi(entry)
		check(err)
		output = append(output, buff)
	}
	return output
}

func ManhattanDistance(first []float64, second []float64) float64 {
	return math.Abs(first[0]-second[0]) + math.Abs(first[1]-second[1])
}

func IntSliceToString(slice []int) string {
	output := ""
	for _, element := range slice {
		output += strconv.Itoa(element) + ","
	}
	output = output[:len(output)-1]
	return output
}

func IntSliceArgMax(slice []int) IntMax {
	argMax := 0
	max := 0
	for idx, val := range slice {
		if val > max {
			argMax = idx
			max = val
		}
	}
	return IntMax{argMax, max}
}

func IntRange(start int, stop int) []int {
	output := make([]int, stop-start)
	for i := start; i < stop; i++ {
		output[i-start] = i
	}
	return output
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IntSliceIndexes(input []int, target int) []int {
	output := make([]int, 0)
	for idx, element := range input {
		if element == target {
			output = append(output, idx)
		}
	}
	return output
}
