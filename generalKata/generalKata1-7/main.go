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

// 6. Create a program that finds the largest element in an array of integers (note: you could use the sort method above as well and just take the last element in the array)
func findLargest(arr []int) int {
	largest := arr[0]
	/* In go, the discard operator is represented by an underscore. We use it here to discard the index because we only need the value in this case.
	This improves are code readability and avoids unused variable errors (go does not allow unused variables)
	*/
	for _, v := range arr {
		if v > largest {
			largest = v
		}
	}
	return largest
}

// 7. Implement a function that removes duplicates from an array/slice
/*
Wanted to try making this function generic so I used a type parameter T. This is a little more complicated than using generics in C# because Go does not have built-in support for generics.
[T comparable] specifies that T must be a comparable type. This is necessary because the function uses a map to track unique values.
*/
func removeDuplicates[T comparable](arr []T) []T {
	// create a map to hold the unique values
	unique := make(map[T]struct{})
	// create a slice to hold the unique values
	uniqueArr := make([]T, 0)

	// iterate over the input array
	for _, v := range arr {
		// check if the value is not in the map
		if _, ok := unique[v]; !ok {
			// add the value to the map
			unique[v] = struct{}{}
			// add the value to the slice
			uniqueArr = append(uniqueArr, v)
		}
	}

	return uniqueArr
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

	fmt.Println("Finding the largest element in the array:", findLargest(arr))
	
	uniqueArr := removeDuplicates(arr)
	fmt.Println("Array with duplicates removed:", uniqueArr)

	strArr := []string{"apple", "banana", "apple", "orange", "banana"}
	uniqueStrArr := removeDuplicates(strArr)
	fmt.Println("String array with duplicates removed:", uniqueStrArr)
}