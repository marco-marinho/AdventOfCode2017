package main

import "fmt"

func main() {
	state := 'A'
	steps := 12586542
	vals := make(map[int]int)
	idx := 0
	for i := 0; i < steps; i++ {
		value, ok := vals[idx]
		if !ok {
			value = 0
			vals[idx] = 0
		}
		switch state {
		case 'A':
			if value == 0 {
				vals[idx] = 1
				idx += 1
				state = 'B'
			} else {
				vals[idx] = 0
				idx -= 1
				state = 'B'
			}
			break
		case 'B':
			if value == 0 {
				vals[idx] = 0
				idx += 1
				state = 'C'
			} else {
				vals[idx] = 1
				idx -= 1
				state = 'B'
			}
			break
		case 'C':
			if value == 0 {
				vals[idx] = 1
				idx += 1
				state = 'D'
			} else {
				vals[idx] = 0
				idx -= 1
				state = 'A'
			}
			break
		case 'D':
			if value == 0 {
				vals[idx] = 1
				idx -= 1
				state = 'E'
			} else {
				vals[idx] = 1
				idx -= 1
				state = 'F'
			}
			break
		case 'E':
			if value == 0 {
				vals[idx] = 1
				idx -= 1
				state = 'A'
			} else {
				vals[idx] = 0
				idx -= 1
				state = 'D'
			}
			break
		case 'F':
			if value == 0 {
				vals[idx] = 1
				idx += 1
				state = 'A'
			} else {
				vals[idx] = 1
				idx -= 1
				state = 'E'
			}
		}

	}
	sum := 0
	for _, val := range vals {
		sum += val
	}
	fmt.Println("Task 01:", sum)
}
