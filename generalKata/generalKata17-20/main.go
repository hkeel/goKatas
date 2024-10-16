package main

import (
	"fmt"
	"sort"
)

// 17. Create a function that merges two sorted arrays into a single sorted array (cheated a little on this one and merged first and then sorted)
func sortArrs(arr1, arr2 []int) []int {
	arr1 = append(arr1, arr2...)
	sort.Ints(arr1)
	return arr1
}

// 18. Implement a program that calculates the sum of digits in a given number
func calcSumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// 19. Write a function that converts a Roman numeral to an integer
func convertRomanToInt(roman string) int {
	romanToInt := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

	n := len(roman)
	total := 0
	for i := 0; i < n; i++ {
		value := romanToInt[roman[i]]
		if i < n-1 && value < romanToInt[roman[i+1]] {
			total -= value
		} else {
			total += value
		}
	}
	return total
}

// 20. Create a program that sorts a slice of strings based on their length
func sortStringsBasedOnLength(strs []string) []string {
	/*
	Custom sorting can be done using the sort.Slice function. The function takes a slice and a function that defines the sorting logic.
	sort.SliceStable ensures that the relative order of strings with the same length is preserved
	*/
    sort.SliceStable(strs, func(i, j int) bool {
        return len(strs[i]) < len(strs[j])
    })
    return strs
}

// Worse way of doing 20.
// func sortStringsBasedOnLength(strs []string) []string {
// 	str := make(map[int][]string)
// 	for i := 0; i < len(strs); i++ {
// 		str[len(strs[i])] = append(str[len(strs[i])], strs[i])
// 	}

// 	str = sortMap(str)
	
// 	res := make([]string, 0)
// 	for _, v := range str {
// 		res = append(res, v...)
// 	}

// 	return res
// }

// func sortMap(m map[int][]string) map[int][]string {
// 	keys := make([]int, 0, len(m))
// 	for k := range m {
// 		keys = append(keys, k)
// 	}
// 	sort.Ints(keys)

// 	sorted := make(map[int][]string)
// 	for _, k := range keys {
// 		sorted[k] = m[k]
// 	}
// 	return sorted
// }

func main() {
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{2, 4, 6, 8}
	fmt.Printf("Sorting %v and %v: %v\n", arr1, arr2, sortArrs(arr1, arr2))

	fmt.Println("Sum of digits in 123:", calcSumOfDigits(123))
	fmt.Println("Sum of digits in 3492837:", calcSumOfDigits(3492837))

	roman := "MCMXCIV"
    fmt.Printf("Roman numeral %s is %d\n", roman, convertRomanToInt(roman)) // expected output: 1994

	strs := []string{"hello", "man", "hannah"}
    sortedStrs := sortStringsBasedOnLength(strs)
    fmt.Printf("Sorted strings based on length: %v\n", sortedStrs) // expected output: [man hello hannah]

	sortedStr := sortStringsBasedOnLength([]string{"hannah"})
	fmt.Printf("Sorted string method with just one string returns: %v\n", sortedStr)
}