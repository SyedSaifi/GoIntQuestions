package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 5, 7, 2, 8, 9, 13, 2}
	maxOne := 0
	maxTwo := 0

	for _, value := range numbers {
		if value > maxOne {
			maxTwo = maxOne
			maxOne = value
		} else if value > maxTwo {
			maxTwo = value
		}
	}

	fmt.Println("The largest number is :: ", maxOne)
	fmt.Println("Second largest number is :: ", maxTwo)
}
