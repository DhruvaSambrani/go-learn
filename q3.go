package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rect struct {
	a float64
	b float64
}

func (c Circle) area() float64 {
	return circleArea(c.radius)
}

func (r Rect) area() float64 {
	return rectArea(r.a, r.b)
}

func circleArea(r float64) float64 {
	return math.Pi * math.Pow(r, 2)
}
func rectArea(a float64, b float64) float64 {
	return a * b
}

func main() {
	fmt.Println(Circle{1}.area())
	fmt.Println(Rect{3, 4}.area())
}
