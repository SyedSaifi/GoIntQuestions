package main

import (
	"fmt"
)

func main() {
	petSlice := make([]string, 5)
	petMap := make(map[string]int)
	petSlice = append(petSlice, "cat", "dog", "rabbit", "dog", "cat", "cat")
	fmt.Println(petSlice)

	for _, pet := range petSlice {
		if petMap[pet] == 0 {
			petMap[pet] = 1
		} else {
			petMap[pet] = petMap[pet] + 1
		}
	}
	fmt.Println(petMap)
}
