package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scanf("%g", &a)
	fmt.Scanf("%g", &b)
	fmt.Scanf("%g", &c)
	var D = math.Pow(b, 2) - 4*a*c
	if D < 0 {
		fmt.Println("No real roots")
	} else if D == 0 {
		fmt.Println("Only one root = %g", -b/(2*a))

	} else {
		fmt.Println("Two roots = %g and %g", (-b+math.Sqrt(D))/(2*a), (-b - math.Sqrt(D))/(2*a))
	}
}
