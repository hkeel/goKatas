package main

import (
	"fmt"
	"os"
	"strconv"
)

func sum(a, b float64) float64 {
    return a + b
}

func sub(a, b float64) float64 {
    return a - b
}

func mult(a, b float64) float64 {
    return a * b
}

func div(a, b float64) float64 {
    return a / b
}

func main() {
    if len(os.Args) < 4 {
        fmt.Println("Usage: go run main.go <num1> <num2> <operation>")
        return
    }

	// common practice in go to use a pattern where a function returns two values: the result and an error (typically seen when parsing, file operations, network requests, etc.)
    num1, err1 := strconv.ParseFloat(os.Args[1], 64)
    num2, err2 := strconv.ParseFloat(os.Args[2], 64)
    operation := os.Args[3]

    if err1 != nil || err2 != nil {
        fmt.Println("Please provide valid numbers for num1 and num2")
        return
    }

    var result float64
    switch operation {
    case "sum":
        result = sum(num1, num2)
    case "sub":
        result = sub(num1, num2)
    case "mult":
        result = mult(num1, num2)
    case "div":
        result = div(num1, num2)
    default:
        fmt.Println("Invalid operation. Please use one of: sum, sub, mult, div")
        return
    }

    fmt.Printf("Result: %f\n", result)
}