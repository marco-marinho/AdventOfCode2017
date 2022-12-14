package main

import (
	"aoc/helpers"
	"fmt"
	"strconv"
	"strings"
)

type Instruction2 struct {
	code string
	arg1 string
	arg2 string
}

type CPU2 struct {
	registers   map[string]int
	instPointer int
	queue       []int
	rom         []Instruction2
	sent        int
}

func main() {
	data := helpers.GetLines("Data/Day18.txt")
	rom := make([]Instruction2, len(data))
	for idx, line := range data {
		pieces := strings.Split(line, " ")
		if len(pieces) < 3 {
			pieces = append(pieces, "")
		}
		rom[idx] = Instruction2{pieces[0], pieces[1], pieces[2]}
	}
	cpu1 := CPU2{make(map[string]int), 0, make([]int, 0), rom, 0}
	cpu1.registers["p"] = 0
	cpu2 := CPU2{make(map[string]int), 0, make([]int, 0), rom, 0}
	cpu2.registers["p"] = 1
	cpus := []CPU2{cpu1, cpu2}
	last1 := -1
	last2 := -1
	for last1 != cpus[0].instPointer || last2 != cpus[1].instPointer {
		last1 = cpus[0].instPointer
		last2 = cpus[1].instPointer
		executeInstruction2(cpus, 0)
		executeInstruction2(cpus, 1)
	}
	fmt.Println("Task 02:", cpus[1].sent)
}

func executeInstruction2(cpus []CPU2, executeID int) {
	cpu := &cpus[executeID]
	instruction := cpu.rom[cpu.instPointer]
	var other *CPU2

	if executeID == 0 {
		other = &cpus[1]
	} else {
		other = &cpus[0]
	}
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
		cpu.sent += 1
		arg1Buff, err := strconv.Atoi(instruction.arg1)
		if err == nil {
			other.queue = append(other.queue, arg1Buff)
		} else {
			val, ok := cpu.registers[instruction.arg1]
			if !ok {
				cpu.registers[instruction.arg1] = 0
				other.queue = append(other.queue, 0)
			} else {
				other.queue = append(other.queue, val)
			}
		}
	} else if instruction.code == "set" {
		cpu.registers[instruction.arg1] = arg2Buff
	} else if instruction.code == "add" {
		cpu.registers[instruction.arg1] += arg2Buff
	} else if instruction.code == "mul" {
		cpu.registers[instruction.arg1] *= arg2Buff
	} else if instruction.code == "mod" {
		cpu.registers[instruction.arg1] %= arg2Buff
	} else if instruction.code == "rcv" {
		if len(cpu.queue) == 0 {
			return
		}
		cpu.registers[instruction.arg1] = cpu.queue[0]
		cpu.queue = cpu.queue[1:]
	} else if instruction.code == "jgz" {
		arg1Buff, err := strconv.Atoi(instruction.arg1)
		if err == nil {
			if arg1Buff > 0 {
				cpu.instPointer += arg2Buff
				return
			}
		}
		val, ok := cpu.registers[instruction.arg1]
		if !ok {
			cpu.registers[instruction.arg1] = 0
		} else if val > 0 {
			cpu.instPointer += arg2Buff
			return
		}
	}
	cpu.instPointer += 1
}
