package main

import (
	"aoc/helpers"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("Task 01: ", Task01())
	fmt.Println("Task 02: ", Task02())
}

func Task01() int {
	data := helpers.GetLines("Data/Day04.txt")
	invalid := 0
	for _, pass := range data {
		wordSet := mapset.NewSet[string]()
		for _, word := range strings.Split(pass, " ") {
			if wordSet.Contains(word) {
				invalid += 1
				break
			} else {
				wordSet.Add(word)
			}
		}
	}
	return len(data) - invalid
}

func Task02() int {
	data := helpers.GetLines("Data/Day04.txt")
	invalid := 0
LineLoop:
	for _, pass := range data {
		words := strings.Split(pass, " ")
		wordSets := make([]map[string]int, len(words))
		for idx, word := range words {
			buffMap := make(map[string]int)
			for _, code := range word {
				char := string(code)
				_, ok := buffMap[char]
				if ok {
					buffMap[char] += 1
				} else {
					buffMap[char] = 1
				}
			}
			for oidx := 0; oidx < idx; oidx++ {
				omap := wordSets[oidx]
				if reflect.DeepEqual(buffMap, omap) {
					invalid += 1
					continue LineLoop
				}
			}
			wordSets[idx] = buffMap
		}
	}
	return len(data) - invalid
}
