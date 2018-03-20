package main

import (
	"fmt"
)

func main() {
	slice := []int{4, 3, 6, 6, 7, 5, 8}
	printMagicPair(slice, 10)
}

func printMagicPair(slice []int, sum int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			if slice[i] == sum-slice[j] && i != j {
				fmt.Println(slice[i], "-", slice[j])
			}
		}

	}
}
