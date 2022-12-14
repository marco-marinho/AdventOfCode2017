package main

import (
	"aoc/helpers"
	"fmt"
	"strconv"
	"strings"
)

type Command struct {
	command string
	arg1    string
	arg2    string
}

func main() {
	data := helpers.GetLines("Data/Day16.txt")
	commands := parseCommands(data)
	input := "abcdefghijklmnop"
	originalInput := input
	repeatLen := 0
	cache := make(map[int]string)
	for i := 0; i < 1000000000; i++ {
		for _, command := range commands {
			input = execCommand(input, command)
		}
		cache[i+1] = input
		if input == originalInput {
			repeatLen = i + 1
			break
		}
	}
	remaining := 1000000000 % repeatLen
	fmt.Println("Task 01:", cache[1])
	fmt.Println("Task 02:", cache[remaining])
}

func parseCommands(data []string) []Command {
	rawCommands := strings.Split(data[0], ",")
	commands := make([]Command, len(rawCommands))
	for idx, input := range rawCommands {
		command := string(input[0])
		args := make([]string, 2)
		if strings.Contains(input[1:], "/") {
			argsBuff := strings.Split(input[1:], "/")
			args[0] = argsBuff[0]
			args[1] = argsBuff[1]
		} else {
			args[0] = input[1:]
			args[1] = ""
		}
		commands[idx] = Command{command, args[0], args[1]}
	}
	return commands
}

func execCommand(input string, command Command) string {
	if command.command == "s" {
		size, _ := strconv.Atoi(command.arg1)
		stop := len(input) - size
		buffEnd := input[stop:]
		buffStart := input[:stop]
		output := buffEnd + buffStart
		return output
	} else {
		var first int
		var second int
		if command.command == "x" {
			first, _ = strconv.Atoi(command.arg1)
			second, _ = strconv.Atoi(command.arg2)
		} else {
			first = strings.Index(input, command.arg1)
			second = strings.Index(input, command.arg2)
		}
		buff1 := rune(input[first])
		buff2 := rune(input[second])
		output := []rune(input)
		output[first] = buff2
		output[second] = buff1
		return string(output)
	}
}
