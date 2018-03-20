package main

import (
	"fmt"
)

func main() {
	numbers := []int{0, 0, 1, 1, 0, 1, 0, 1, 0, 1, 1, 0, 1}

	left := 0
	right := len(numbers) - 1

	for left < right {
		for numbers[left] == 0 && left < right {
			left++
		}

		for numbers[right] == 0 && left < right {
			right--
		}

		for left < right {
			numbers[left] = 0
			numbers[right] = 1
			left++
			right--
		}
	}

	fmt.Print(numbers)
}
