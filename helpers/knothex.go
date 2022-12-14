package helpers

import (
	"reflect"
	"strconv"
	"strings"
)

func KnotHex(input string) string {
	list := IntRange(0, 256)
	key := []int{17, 31, 73, 47, 23}
	rotations := make([]int, len(input))
	for idx, char := range input {
		rotations[idx] = int(char)
	}
	rotations = append(rotations, key...)
	ApplyReversions(list, rotations, 64)
	xor := applyXOR(list)
	return strings.Join(ToHex(xor), "")
}

func ToHex(list []int) []string {
	output := make([]string, len(list))
	for idx, value := range list {
		temp := strconv.FormatInt(int64(value), 16)
		if len(temp) < 2 {
			temp = "0" + temp
		}
		output[idx] = temp
	}
	return output
}

func applyXOR(list []int) []int {
	numBlocks := len(list) / 16
	output := make([]int, numBlocks)
	for start := 0; start < numBlocks; start++ {
		first := list[16*start]
		for _, other := range list[(16*start)+1 : 16*(start+1)] {
			first = first ^ other
		}
		output[start] = first
	}
	return output
}

func ApplyReversions(list []int, reversions []int, rounds int) {
	pos := 0
	skip := 0
	for i := 0; i < rounds; i++ {
		for _, selLen := range reversions {
			reverseCircular(list, pos, selLen)
			pos += skip + selLen
			skip += 1
		}
	}
}

func reverseCircular(list []int, pos int, selLen int) {
	listLen := len(list)
	idxs := make([]int, selLen)
	buff := make([]int, selLen)
	swap := reflect.Swapper(buff)
	for i := 0; i < selLen; i++ {
		idxs[i] = (pos + i) % listLen
		buff[i] = list[(pos+i)%listLen]
	}
	for i, j := 0, len(buff)-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
	for i := 0; i < len(idxs); i++ {
		list[idxs[i]] = buff[i]
	}
}
