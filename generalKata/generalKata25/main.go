package main

import "fmt"

func calcCircleArea(r float64) float64 {
	return 3.14 * r * r
}

func calcRectangleArea(l, w float64) float64 {
	return l * w
}

func calcTriangleArea(b, h float64) float64 {
	return 0.5 * b * h
}

func main() {
	fmt.Println("Circle with radius of 3:", calcCircleArea(3))
	fmt.Println("Rectangle with length of 3, width of 4:", calcRectangleArea(3, 4))
	fmt.Println("Triangle with base of 3, height of 4:", calcTriangleArea(3, 4))
}