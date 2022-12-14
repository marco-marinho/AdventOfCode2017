package main

import (
	"aoc/helpers"
	"fmt"
	"strings"
)

type Carrier struct {
	direction rune
	position  [2]int
}

func main() {
	Task01()
	Task02()
}

func Task02() {
	infected, startPos := getInfected()
	carrier := Carrier{'u', startPos}
	totalInfected := 0
	weakened := make(map[string]bool)
	flagged := make(map[string]bool)
	for i := 0; i < 10000000; i++ {
		carrierPosStr := fmt.Sprint(carrier.position)
		_, isInfected := infected[carrierPosStr]
		_, isWeakened := weakened[carrierPosStr]
		_, isFlagged := flagged[carrierPosStr]
		if isWeakened {
			totalInfected += 1
			delete(weakened, carrierPosStr)
			infected[carrierPosStr] = true
			direction, step := straight(carrier)
			walk(&carrier, direction, step)
		} else if isInfected {
			delete(infected, carrierPosStr)
			flagged[carrierPosStr] = true
			direction, step := turn(carrier, 'r')
			walk(&carrier, direction, step)
		} else if isFlagged {
			delete(flagged, carrierPosStr)
			direction, step := reverse(carrier)
			walk(&carrier, direction, step)
		} else {
			weakened[carrierPosStr] = true
			direction, step := turn(carrier, 'l')
			walk(&carrier, direction, step)
		}
	}
	fmt.Println("Task 02:", totalInfected)
}

func Task01() {
	infected, startPos := getInfected()
	carrier := Carrier{'u', startPos}
	totalInfected := 0
	for i := 0; i < 10000; i++ {
		carrierPosStr := fmt.Sprint(carrier.position)
		_, ok := infected[carrierPosStr]
		if ok {
			direction, step := turn(carrier, 'r')
			delete(infected, carrierPosStr)
			walk(&carrier, direction, step)
		} else {
			direction, step := turn(carrier, 'l')
			totalInfected += 1
			infected[carrierPosStr] = true
			walk(&carrier, direction, step)
		}
	}
	fmt.Println("Task 01:", totalInfected)
}

func walk(carrier *Carrier, direction rune, coords [2]int) {
	carrier.direction = direction
	carrier.position[0] += coords[0]
	carrier.position[1] += coords[1]
}

func getInfected() (map[string]bool, [2]int) {
	infected := make(map[string]bool)
	data := helpers.GetLines("Data/Day22.txt")
	for row, line := range data {
		line := strings.ReplaceAll(line, "\r", "")
		for col, char := range line {
			if string(char) == "#" {
				tmp := [2]int{row, col}
				infected[fmt.Sprint(tmp)] = true
			}
		}
	}
	pos := [2]int{len(data) / 2, (len(data[0]) - 1) / 2}
	return infected, pos
}

func straight(carrier Carrier) (rune, [2]int) {
	if carrier.direction == 'u' {
		return 'u', [2]int{-1, 0}
	} else if carrier.direction == 'd' {
		return 'd', [2]int{1, 0}
	} else if carrier.direction == 'l' {
		return 'l', [2]int{0, -1}
	} else if carrier.direction == 'r' {
		return 'r', [2]int{0, 1}
	} else {
		panic("Invalid direction")
	}
}

func reverse(carrier Carrier) (rune, [2]int) {
	if carrier.direction == 'u' {
		return 'd', [2]int{1, 0}
	} else if carrier.direction == 'd' {
		return 'u', [2]int{-1, 0}
	} else if carrier.direction == 'l' {
		return 'r', [2]int{0, 1}
	} else if carrier.direction == 'r' {
		return 'l', [2]int{0, -1}
	} else {
		panic("Invalid direction")
	}
}

func turn(carrier Carrier, side rune) (rune, [2]int) {
	directions := [4]rune{'l', 'u', 'r', 'd'}
	steps := [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	idx := getIndex(directions, carrier.direction)
	if side == 'l' {
		idx -= 1
		if idx < 0 {
			idx = 3
		}
		return directions[idx], steps[idx]
	} else if side == 'r' {
		idx += 1
		if idx > 3 {
			idx = 0
		}
		return directions[idx], steps[idx]
	} else {
		panic("Invalid Direction")
	}
}

func getIndex(directions [4]rune, step rune) int {
	idxOut := -1
	for idx, char := range directions {
		if step == char {
			idxOut = idx
			break
		}
	}
	return idxOut
}
