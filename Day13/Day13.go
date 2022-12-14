package main

import (
	"aoc/helpers"
	"fmt"
	"strconv"
	"strings"
)

type Scanner struct {
	Depth int
	Range int
}

func main() {
	data := helpers.GetLines("Data/Day13.txt")
	scanners := make([]Scanner, len(data))
	for idx, line := range data {
		buff := strings.Split(strings.Trim(line, "\r"), ": ")
		depthBuff, _ := strconv.Atoi(buff[0])
		rangeBuff, _ := strconv.Atoi(buff[1])
		scanners[idx] = Scanner{depthBuff, rangeBuff}
	}
	Task01(scanners)
	Task02(scanners)
}

func Task01(scanners []Scanner) {
	severity := 0
	for _, scanner := range scanners {
		if scanner.Depth%((scanner.Range-1)*2) == 0 {
			severity += scanner.Depth * scanner.Range
		}
	}
	fmt.Println("Task 01:", severity)
}

func Task02(scanners []Scanner) {
	delay := -1
	detected := true
	for detected {
		delay += 1
		detected = false
		for _, scanner := range scanners {
			if (scanner.Depth+delay)%((scanner.Range-1)*2) == 0 {
				detected = true
				break
			}
		}
	}
	fmt.Println("Task 02:", delay)
}
