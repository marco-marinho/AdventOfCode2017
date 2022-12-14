package main

import "fmt"

func main() {
	Task01()
	Task02()
}

func Task01() {
	b, c, d, e, f, h := 0, 0, 0, 0, 0, 0
	muls := 0
	for true {
		b = 84
		c = b
		f = 1
		d = 2
		for ok := true; ok; ok = d != b {
			e = 2
			for ok2 := true; ok2; ok2 = e != b {
				if d*e == b {
					f = 0
				}
				e++
				muls += 1
			}
			d++
		}
		if f == 0 {
			h++
		}
		if b == c {
			break
		}
		b -= -17
	}
	fmt.Println("Task 01:", muls)
}

func Task02() {
	b, c, d, f, g, h := 0, 0, 0, 0, 0, 0
	b = 84
	c = b
	b = b*100 + 100000
	c = b + 17000
	for ok := true; ok; ok = g != 0 {
		f = 1
		d = 2
		for d = 2; d*d <= b; d++ {
			if b%d == 0 {
				f = 0
				break
			}
		}
		if f == 0 {
			h++
		}
		g = b - c
		b += 17
	}
	fmt.Println("Task 02:", h)
}
