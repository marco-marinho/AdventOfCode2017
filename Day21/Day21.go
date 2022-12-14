package main

import (
	"aoc/helpers"
	"fmt"
	"strings"
)

func main() {
	Task01()
	Task02()
}

func Task01() {
	rules := getRules()
	a := [][]string{{".", "#", "."},
		{".", ".", "#"},
		{"#", "#", "#"}}
	for i := 0; i < 5; i++ {
		a = runIteration(a, rules)
	}
	fmt.Println("Task 01:", countOn(a))
}

func Task02() {
	rules := getRules()
	a := [][]string{{".", "#", "."},
		{".", ".", "#"},
		{"#", "#", "#"}}
	for i := 0; i < 18; i++ {
		a = runIteration(a, rules)
	}
	fmt.Println("Task 02:", countOn(a))
}

func countOn(input [][]string) int {
	count := 0
	for _, row := range input {
		for _, char := range row {
			if char == "#" {
				count += 1
			}
		}
	}
	return count
}

func runIteration(input [][]string, rules map[string][][]string) [][]string {
	blocks := getBlocks(input)
	output := getNextMat(input)
	for idx, block := range blocks {
		toWrite, ok := rules[matToString(block)]
		if !ok {
			panic("No rule found!")
		}
		writeBlock(idx, output, toWrite)
	}
	return output
}

func getRules() map[string][][]string {
	data := helpers.GetLines("Data/Day21.txt")
	changeMap := make(map[string][][]string)
	for _, line := range data {
		splitLine := strings.Split(strings.ReplaceAll(line, "\r", ""), " => ")
		reference := inputStrToMat(splitLine[0])
		transform := inputStrToMat(splitLine[1])
		references := getAllTransformations(reference)
		for _, ref := range references {
			changeMap[matToString(ref)] = transform
		}
	}
	return changeMap
}

func writeBlock(idx int, input [][]string, block [][]string) {
	blockPerRow := len(input) / len(block)
	coords := [2]int{(idx / blockPerRow) * len(block), (idx % blockPerRow) * len(block)}
	size := len(block)
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			input[row+coords[0]][col+coords[1]] = block[row][col]
		}
	}
}

func getBlocks(input [][]string) [][][]string {
	blockSize := 0
	if len(input)%2 == 0 {
		blockSize = 2
	} else if len(input)%3 == 0 {
		blockSize = 3
	}
	if blockSize == 0 {
		panic("Invalid block size")
	}
	numBlocks := (len(input) * len(input)) / (blockSize * blockSize)
	output := make([][][]string, numBlocks)
	idx := 0
	for rowStart := 0; rowStart < len(input); rowStart += blockSize {
		for colStart := 0; colStart < len(input); colStart += blockSize {
			output[idx] = make([][]string, blockSize)
			for row := 0; row < blockSize; row++ {
				output[idx][row] = make([]string, blockSize)
				for col := 0; col < blockSize; col++ {
					output[idx][row][col] = input[rowStart+row][colStart+col]
				}
			}
			idx += 1
		}
	}
	return output
}

func getNextMat(input [][]string) [][]string {
	nextLen := 0
	if len(input)%2 == 0 {
		nextLen = len(input) + len(input)/2
	} else if len(input)%3 == 0 {
		nextLen = len(input) + len(input)/3
	}
	if nextLen == 0 {
		panic("Invalid next size")
	}
	output := make([][]string, nextLen)
	for i := 0; i < nextLen; i++ {
		output[i] = make([]string, nextLen)
	}
	return output
}

func inputStrToMat(input string) [][]string {
	pieces := strings.Split(input, "/")
	nrows := len(pieces)
	ncols := len(pieces[1])
	output := make([][]string, nrows)
	for i := 0; i < ncols; i++ {
		output[i] = make([]string, ncols)
	}
	for i, row := range pieces {
		for j, char := range row {
			output[i][j] = string(char)
		}
	}
	return output
}

func getAllTransformations(input [][]string) [][][]string {
	output := make([][][]string, 8)
	A := input
	output[0] = A
	A = transpose(A)
	output[1] = A
	A = fliplr(A)
	output[2] = A
	A = transpose(A)
	output[3] = A
	A = fliplr(A)
	output[4] = A
	A = transpose(A)
	output[5] = A
	A = fliplr(A)
	output[6] = A
	A = transpose(A)
	output[7] = A
	return output
}

func fliplr(input [][]string) [][]string {
	nrows := len(input)
	ncols := len(input[0])
	output := make([][]string, nrows)
	for i := 0; i < ncols; i++ {
		output[i] = make([]string, ncols)
	}

	for i := 0; i < nrows; i++ {
		for j := 0; j < ncols/2+1; j++ {
			output[i][j] = input[i][ncols-1-j]
			output[i][ncols-1-j] = input[i][j]
		}
	}
	return output
}

func flipud(input [][]string) [][]string {
	nrows := len(input)
	ncols := len(input[0])
	output := make([][]string, nrows)
	for i := 0; i < ncols; i++ {
		output[i] = make([]string, ncols)
	}

	for i := 0; i < nrows/2+1; i++ {
		for j := 0; j < ncols; j++ {
			output[i][j] = input[nrows-1-i][j]
			output[nrows-1-i][j] = input[i][j]
		}
	}
	return output
}

func transpose(input [][]string) [][]string {
	nrows := len(input[0])
	ncols := len(input)
	output := make([][]string, nrows)
	for i := range output {
		output[i] = make([]string, ncols)
	}
	for i := 0; i < nrows; i++ {
		for j := 0; j < ncols; j++ {
			output[i][j] = input[j][i]
		}
	}
	return output
}

func matToString(input [][]string) string {
	var sb strings.Builder
	for _, line := range input {
		sb.WriteString(strings.Join(line, ""))
		sb.WriteString("\n")
	}
	return sb.String()
}
