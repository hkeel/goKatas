package main

import "fmt"

func createMultTable(n int) [][]int {
	table := make([][]int, n)
    for i := 1; i <= n; i++ {
        row := make([]int, n)
        for j := 1; j <= n; j++ {
            row[j-1] = i * j
        }
        table[i-1] = row
    }
    return table
}

func main() {
	n := 6
    table := createMultTable(n)
    fmt.Printf("Multiplication table up to %d:\n", n)
    for _, row := range table {
        fmt.Println(row)
    }
}