package main

import (
	"fmt"
)

var sum = 0
var jump = 0
var loop = false

func main() {
	arr := []int{2, 1, 4, 2, 6, 7}

	for true {
		sum = arr[sum] + sum
		if sum > len(arr)-1 || sum < 0 {
			break
		}
		arr[sum] = sum
		jump++

		if jump > len(arr)-1 {
			loop = true
			break
		}
	}

	if loop {
		fmt.Println("Looping pawn")
	} else {
		fmt.Println("No of Jumps is => ", jump)
	}
}
