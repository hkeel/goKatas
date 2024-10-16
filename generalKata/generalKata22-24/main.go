package main

import "fmt"

/*
Rules:
A year that is evenly divisible by 4 is a leap year.
However, if the year is also evenly divisible by 100, it is not a leap year unless it is also evenly divisible by 400.
*/

// 22. Write a program that checks if a given year is a leap year
func isLeapYear(year int) bool {
	if year % 4 == 0 {
		if year % 100 == 0 {
			if year % 400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}

// 23. Create a function that finds the GCD of two numbers
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// 24. Implement a program that performs linear search on an array to find a specifc elemment - made function generic and used some error handling
func linearSearch[T comparable](arr []T, target T) (T, error) {
	for _, v := range arr {
		if v == target {
			return v, nil
		}
	}
	var zeroVal T
	return zeroVal, fmt.Errorf("%v not found in array: %v", target, arr)
}

func main() {
	years := []int{2000, 2006, 2008, 2009, 2012, 2070, 2072, 2076, 2080, 2084, 2096, 2100}
	for _, year := range years {
		if isLeapYear(year) {
			println(year, "is a leap year")
		} else {
			println(year, "is not a leap year")
		}
	}

	println("GCD of 12 and 15 is:", gcd(12, 15))
	println("GCD of 30 and 80 is:", gcd(30, 80))
	println("GCD of 102 and 120 is:", gcd(102, 120))
	println("GCD of 100 and 10 is:", gcd(100, 10))

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 5
	fmt.Println("Linear search for", target, "in", arr)
	if val, err := linearSearch(arr, target); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Element found:", val)
	}
}