package main

import (
	"fmt"
	"sort"
	"strings"
)

// 1. Write a program that calculates the factorial of a number
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 2. Implement a function that reverses a string
func reverse(s string) string {
	if len(s) == 0 {
		return s
	}
	return reverse(s[1:]) + string(s[0])
}

// 3. Create a function that checks if a given string is a palindrome
func isPalindrome(s string) bool {
	// use strings package to convert to lowercase for case insensitive comparison
	s = strings.ToLower(s)
	return s == reverse(s)
}

// 4. Write a program that calculates the nth Fibonacci number
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 5. Implement a function that sorts an array of integers in ascending order
func sortArr(arr []int) []int {
	sort.Ints(arr)
	return arr
}

func main() {
	fmt.Println("Calculating the factorial of 5:", factorial(5))

	fmt.Println("Reversing the string 'hello':", reverse("hello"))
	fmt.Println("Reversing the string 'w':", reverse("w"))
	fmt.Println("Reversing the string '':", reverse(""))

	fmt.Println("Checking if 'hello' is a palindrome:", isPalindrome("hello"))
	fmt.Println("Checking if 'hannah' is a palindrome:", isPalindrome("hannah"))

	fmt.Println("Calculating the 5th Fibonacci number:", fibonacci(5))

	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fmt.Println("Original array:", arr)
	fmt.Println("Sorted array:", sortArr(arr))
}