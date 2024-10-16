package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// 9. Implement a function that finds the longest word in a sentence - note: if there are multiple words with the maximum length this method will return the first one.
func findLongest(sentence string) string {
	words := strings.Fields(sentence)
	longest := ""
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

// 10. Create a program converts a decimal number to binary
func decToBinary(n float64) string {
    // Convert the float to an integer by truncating the fractional part
    intPart := int(n)

    if intPart == 0 {
        return "0"
    }

    isNegative := intPart < 0
    if isNegative {
        intPart = -intPart
    }

    binary := ""
    for intPart > 0 {
        // similar to the Printf function, the Sprintf function formats the string but returns it instead of printing it
        binary = fmt.Sprintf("%d", intPart%2) + binary
        intPart = intPart / 2
    }

	// For simplicity, I am simply prefixing the binary rep. with a negative sign if the number is negative vs. using two's complement
    if isNegative {
        binary = "-" + binary
    }

    return binary
}

// 11. Write a function that checks if a number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 12. Implement a prgram that generates a random number within a given range
func generateRandom(min, max int) int {
    return rand.Intn(max-min+1) + min
}

// 13. Write a function that counts the number of vowels in a string
func countVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, c := range s {
		if strings.ContainsRune(vowels, c) {
			count++
		}
	}
	return count
}

// 14. Create a function that finds the second smallest element in an array - Note: trying out some error handling in this one
func secondSmallest(arr []int) (int, error) {
	/*
		Using a map as a set to store unique elements in the array. Using true as a placeholder to indicate key exists in the map.
	*/
    uniqueElements := make(map[int]bool)
    for _, num := range arr {
        uniqueElements[num] = true
    }

    uniqueArr := make([]int, 0, len(uniqueElements))
    for num := range uniqueElements {
        uniqueArr = append(uniqueArr, num)
    }

    sort.Ints(uniqueArr)
    if len(uniqueArr) < 2 {
		return 0, fmt.Errorf("array %v must have at least 2 unique elements", arr)
    }
    return uniqueArr[1], nil
}

// 15. Implement a function that checks if two strings are anagrams of each other
func areAnagrams(s1, s2 string) bool {
	return sortString(s1) == sortString(s2)
}

// Helper function to sort a string for anagram comparison
func sortString(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func main() {
	/*
		Standard practice to put the defer function at the top of the main function. This ensures that the
		deferred function is registered as early as possible, allowing it to handle any panics that might occur during the execution of the main function.

		Handling panic:
			- Defer and Recover: Use defer and recover to handle panics and allow the program to continue executing. Recover is a built-in function that regains
			control of a panicking goroutine.
			- Graceful Error Handling: Instead of using panic, you can retrun an error from the function (ex: secondSmallest) and handle it appropriately in
			the calling code.
	*/
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

	/*
	quick notes about := operator:
		- The := operator is a shorthand for declaring and initializing variables. 
		- It is only available inside functions.
		- The type of variable is inferred from the value on the right-hand side.
	*/
	sentence := "The quickest brown fox jumps over the lazy dog"
	fmt.Printf("The longest word in the sentence: '%s' is %s\n",sentence, findLongest(sentence)) // expected output: "quick"

	fmt.Println("Decimal to Binary 0:", decToBinary(0)) // expected output: "0"
	fmt.Println("Decimal to Binary 5.6:", decToBinary(5.6))  // expected output: "101"
    fmt.Println("Decimal to Binary -5.6:", decToBinary(-5.6)) // expected output: "-101"
	fmt.Println("Decimal to Binary 198.6:", decToBinary(198.6)) // expected output: "11000110"

	fmt.Println("Is 0 prime?", isPrime(0))   // expected output: false
    fmt.Println("Is 2 prime?", isPrime(2))   // expected output: true
    fmt.Println("Is 3 prime?", isPrime(3))   // expected output: true
    fmt.Println("Is 24 prime?", isPrime(24))   // expected output: false
    fmt.Println("Is 29 prime?", isPrime(29)) // expected output: true

	min := 10
    max := 20
    randomNumber := generateRandom(min, max)
    fmt.Printf("Random number between %d and %d(inclusive): %d\n", min, max, randomNumber)

	fmt.Printf("Number of vowels in the sentence: '%s' is %d\n", sentence, countVowels(sentence)) // expected output: 12

    // Test secondSmallest function with an array that works
    arr1 := []int{4, 2, 2, 3, 1, 4, 5}
    printSecondSmallest(arr1)

    // Test secondSmallest function with an array that returns an error
    arr2 := []int{1, 1}
    printSecondSmallest(arr2)

	str1 := "listen"
    str2 := "silent"
    fmt.Printf("Are '%s' and '%s' anagrams? %v\n", str1, str2, areAnagrams(str1, str2)) // expected output: true
    str5 := "anagram"
    str6 := "nag a ram"
    fmt.Printf("Are '%s' and '%s' anagrams? %v\n", str5, str6, areAnagrams(str5, str6)) // expected output: true
    str7 := "hello"
    str8 := "world"
    fmt.Printf("Are '%s' and '%s' anagrams? %v\n", str7, str8, areAnagrams(str7, str8)) // expected output: false
}

// Helper function to print the second smallest element or error
func printSecondSmallest(arr []int) {
    secondSmallestElement, err := secondSmallest(arr)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("The second smallest element in the array %v is %d\n", arr, secondSmallestElement)
    }
}