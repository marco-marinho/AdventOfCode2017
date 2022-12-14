package main

import (
	"aoc/helpers"
	"fmt"
	"math"
	"math/cmplx"
	"strconv"
)

func main() {
	Task01(347991.0)
	Task02(347991)
}

func Task01(input float64) {
	input -= 1
	coords := getCoords(input)
	dist := helpers.ManhattanDistance([]float64{0.0, 0.0}, coords)
	fmt.Println("Task 01: ", dist)
}

// https://web.archive.org/web/20141202041502/https://danpearcymaths.wordpress.com/2012/09/30/infinity-programming-in-geogebra-and-failing-miserably/
func getCoords(n float64) []float64 {
	p := math.Floor(math.Sqrt(4*n - 1))
	q := n - math.Floor(math.Pow(p, 2)/4)
	i := complex(0, 1)
	r1 := complex(q, 0) * cmplx.Pow(i, complex(p, 0))
	r2 := complex(math.Floor((p+2)/4), -math.Floor((p+1)/4)) * cmplx.Pow(i, complex(p-1, 0))
	z := r1 + r2
	coords := []float64{math.Round(real(z)), math.Round(imag(z))}
	return coords
}

func Task02(max int) int {
	grid := make(map[string]int)
	grid["0,0"] = 1
	size := 1
	current := []int{0, 1}
	curVal := 1
	for curVal < max {
		steps := getSteps(size)
		for _, val := range steps {
			adjacent := getAdjacent(current)
			acc := 0
			for _, next := range adjacent {
				if value, ok := grid[next]; ok {
					acc += value
				}
			}
			grid[coordToString(current)] = acc
			curVal = acc
			if curVal > max {
				break
			}
			current = []int{current[0] + val[0], current[1] + val[1]}
		}
		size += 2
	}
	fmt.Println("Task 02: ", curVal)
	return curVal
}

func coordToString(coords []int) string {
	s1 := strconv.Itoa(coords[0])
	s2 := strconv.Itoa(coords[1])
	return s1 + "," + s2
}

func getAdjacent(coords []int) []string {
	output := make([]string, 8)
	offsets := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
		{1, 1},
		{1, -1},
		{-1, -1},
		{-1, 1},
	}
	for idx, element := range offsets {
		s1 := strconv.Itoa(coords[0] + element[0])
		s2 := strconv.Itoa(coords[1] + element[1])
		output[idx] = s1 + "," + s2
	}
	return output
}

func getSteps(size int) [][]int {
	output := make([][]int, 4*size+4)
	idx := 0
	for i := 0; i < size; i++ {
		output[idx] = []int{1, 0}
		idx++
	}
	for i := 0; i < size+1; i++ {
		output[idx] = []int{0, -1}
		idx++
	}
	for i := 0; i < size+1; i++ {
		output[idx] = []int{-1, 0}
		idx++
	}
	for i := 0; i < size+2; i++ {
		output[idx] = []int{0, 1}
		idx++
	}
	return output
}
