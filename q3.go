package main

import (
	"fmt"
	"math"
)

func circleArea(r float64) float64 {
	return math.Pi * math.Pow(r, 2)
}
func rectArea(a float64, b float64) float64 {
	return a * b
}

func area(r float64) float64 {
	return circleArea(r)
}

func area(a float64, b float64) float64 {
	return rectArea(a, b)
}

func main() {
	fmt.Println(circleArea(3))
	fmt.Println(rectArea(3, 4))
}
