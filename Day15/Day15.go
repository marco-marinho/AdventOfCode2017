package main

import "fmt"

func main() {

	var FactorA uint64 = 16807
	var FactorB uint64 = 48271
	var Denominator uint64 = 2147483647
	var filter uint64 = 65535
	var A uint64 = 116
	var B uint64 = 299
	match := 0
	for i := 0; i < 40000000; i++ {
		A = (A * FactorA) % Denominator
		B = (B * FactorB) % Denominator
		if A&filter == B&filter {
			match += 1
		}
	}
	fmt.Println("Task 01:", match)
	match = 0
	A = 116
	B = 299
	for i := 0; i < 5000000; i++ {
		A = (A * FactorA) % Denominator
		for A%4 != 0 {
			A = (A * FactorA) % Denominator
		}
		B = (B * FactorB) % Denominator
		for B%8 != 0 {
			B = (B * FactorB) % Denominator
		}
		if A&filter == B&filter {
			match += 1
		}
	}
	fmt.Println("Task 02:", match)
}
