package main

import (
	"aoc/helpers"
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	target        string
	action        string
	value         int
	firstOperand  string
	condition     string
	secondOperand int
}

func main() {
	data := helpers.GetLines("Data/Day08.txt")
	instructions := make([]Instruction, len(data))
	for idx, line := range data {
		parts := strings.Split(line, " ")
		target := parts[0]
		action := parts[1]
		value, _ := strconv.Atoi(parts[2])
		first := parts[4]
		cond := parts[5]
		second, _ := strconv.Atoi(parts[6])
		instructions[idx] = Instruction{target, action, value, first, cond, second}
	}
	execute(instructions)
}

func execute(instructions []Instruction) {
	regs := make(map[string]int)
	max := 0
	for _, instruction := range instructions {
		_, ok := regs[instruction.target]
		if !ok {
			regs[instruction.target] = 0
		}
		_, ok = regs[instruction.firstOperand]
		if !ok {
			regs[instruction.firstOperand] = 0
		}

		condition := false
		if instruction.condition == "==" {
			condition = regs[instruction.firstOperand] == instruction.secondOperand
		} else if instruction.condition == "!=" {
			condition = regs[instruction.firstOperand] != instruction.secondOperand
		} else if instruction.condition == ">" {
			condition = regs[instruction.firstOperand] > instruction.secondOperand
		} else if instruction.condition == "<" {
			condition = regs[instruction.firstOperand] < instruction.secondOperand
		} else if instruction.condition == ">=" {
			condition = regs[instruction.firstOperand] >= instruction.secondOperand
		} else if instruction.condition == "<=" {
			condition = regs[instruction.firstOperand] <= instruction.secondOperand
		}
		if condition {
			if instruction.action == "inc" {
				regs[instruction.target] += instruction.value
			} else {
				regs[instruction.target] -= instruction.value
			}
			if regs[instruction.target] > max {
				max = regs[instruction.target]
			}
		}
	}
	largestReg := largestRegVal(regs)
	fmt.Println("Task 01: ", largestReg)
	fmt.Println("Task 02: ", max)
}

func largestRegVal(regs map[string]int) int {
	max := 0
	for _, value := range regs {
		if value > max {
			max = value
		}
	}
	return max
}
