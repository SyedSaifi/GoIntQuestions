package main

import (
	"fmt"
)

func main() {
	array1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	array2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Original array :%v", array1)
	leftShiftedArray := leftShiftArray(array1, 3)
	rightShiftedArray := rightShiftArray(array2, 3)
	fmt.Println("")
	fmt.Printf("Left shifted :%v", leftShiftedArray)
	fmt.Println("")
	fmt.Printf("Right shifted :%v", rightShiftedArray)
}

func leftShiftArray(array []int, shift int) []int {
	for index := 0; index < shift; index++ {
		first := array[0]

		for index := 0; index < len(array)-1; index++ {
			array[index] = array[index+1]
		}

		array[len(array)-1] = first
	}
	return array
}

func rightShiftArray(array []int, shift int) []int {
	for index := 0; index < shift; index++ {
		last := array[len(array)-1]

		for index := len(array) - 1; index > 0; index-- {
			array[index] = array[index-1]
		}

		array[0] = last
	}
	return array
}
