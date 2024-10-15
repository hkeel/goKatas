package main

import "fmt"

// 1. Write a program that calculates the factorial of a number
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("Calculating the factorial of 5:", factorial(5))
}