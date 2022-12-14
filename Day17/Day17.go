package main

import "fmt"

func main() {
	orig := []int{0}
	pos := 0
	steps := 386
	for i := 1; i <= 2017; i++ {
		orig, pos = spin(orig, i, pos, steps)
	}
	fmt.Println("Task 01:", orig[pos+1])
	afterZero := 0
	for i := 1; i <= 50000000; i++ {
		pos = (pos+steps)%i + 1
		if pos == 1 {
			afterZero = i
		}
	}
	fmt.Println("Task 02:", afterZero)
}

func spin(arr []int, toInsert int, pos int, steps int) ([]int, int) {
	newPos := (pos + steps) % len(arr)
	front := make([]int, newPos+1)
	back := make([]int, len(arr)-newPos-1)
	copy(front, arr[:newPos+1])
	copy(back, arr[newPos+1:])
	newArr := append(append(front, toInsert), back...)
	return newArr, newPos + 1
}
