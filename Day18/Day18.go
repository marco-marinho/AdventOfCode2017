package main

import (
	"aoc/helpers"
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	code string
	arg1 string
	arg2 string
}

type CPU struct {
	registers   map[string]int
	played      int
	instPointer int
	recovered   int
}

func main() {
	data := helpers.GetLines("Data/Day18.txt")
	rom := make([]Instruction, len(data))
	for idx, line := range data {
		pieces := strings.Split(line, " ")
		if len(pieces) < 3 {
			pieces = append(pieces, "")
		}
		rom[idx] = Instruction{pieces[0], pieces[1], pieces[2]}
	}
	cpu := CPU{make(map[string]int), 0, 0, 0}
	for cpu.instPointer < len(rom) {
		executeInstruction(&cpu, rom[cpu.instPointer])
		if cpu.recovered != 0 {
			fmt.Println("Task 01:", cpu.recovered)
			break
		}
	}
}

func executeInstruction(cpu *CPU, instruction Instruction) {
	arg2Buff, err := strconv.Atoi(instruction.arg2)
	var ok bool
	if err != nil {
		arg2Buff, ok = cpu.registers[instruction.arg2]
		if !ok {
			cpu.registers[instruction.arg2] = 0
			arg2Buff = 0
		}
	}
	if instruction.code == "snd" {
		val, ok := cpu.registers[instruction.arg1]
		if !ok {
			cpu.registers[instruction.arg1] = 0
			cpu.played = 0
		} else {
			cpu.played = val
		}
		cpu.instPointer += 1
	} else if instruction.code == "set" {
		cpu.registers[instruction.arg1] = arg2Buff
		cpu.instPointer += 1
	} else if instruction.code == "add" {
		cpu.registers[instruction.arg1] += arg2Buff
		cpu.instPointer += 1
	} else if instruction.code == "mul" {
		cpu.registers[instruction.arg1] *= arg2Buff
		cpu.instPointer += 1
	} else if instruction.code == "mod" {
		cpu.registers[instruction.arg1] = cpu.registers[instruction.arg1] % arg2Buff
		cpu.instPointer += 1
	} else if instruction.code == "rcv" {
		val, ok := cpu.registers[instruction.arg1]
		if !ok {
			cpu.registers[instruction.arg1] = 0
		} else if val != 0 {
			cpu.recovered = cpu.played
		}
		cpu.instPointer += 1
	} else if instruction.code == "jgz" {
		val, ok := cpu.registers[instruction.arg1]
		if !ok {
			cpu.registers[instruction.arg1] = 0
			cpu.instPointer += 1
		} else if val > 0 {
			cpu.instPointer += arg2Buff
		} else {
			cpu.instPointer += 1
		}
	}
}
