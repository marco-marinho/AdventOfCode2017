package main

import (
	"aoc/helpers"
	"fmt"
)

func main() {
	Task01()
	Task02()
}

func Task02() {
	data := helpers.GetLines("Data/Day10.txt")[0]
	fmt.Println("Task 02:", helpers.KnotHex(data))
}

func Task01() {
	list := helpers.IntRange(0, 256)
	data := helpers.GetLines("Data/Day10.txt")
	reversions := helpers.StringToInts(data[0], ",")
	helpers.ApplyReversions(list, reversions, 1)
	fmt.Println("Task 01:", list[0]*list[1])
}
