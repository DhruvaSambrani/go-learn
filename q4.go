package main

import "fmt"

func main() {
	numbers := []int{ 1,2,3,4,5,12,23,123,4,2,12,12,4,6,3,2 }
	search := 12
	count := 0 
	for i := 0; i < len(numbers); i++ {
		if search == numbers[i] {
			count++
		}
	}
	fmt.Println(count, numbers)
	fmt.Println(numbers  numbers)
}
