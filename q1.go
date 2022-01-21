package main

import "fmt"

func main() {
	var P, R, T float64
	fmt.Scanf("%g", &P)
	fmt.Scanf("%g", &R)
	fmt.Scanf("%g", &T)
	fmt.Println(P * R * T / 100)
}
