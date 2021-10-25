package main

import (
	"fmt"
)

func bubbleSort(x []int) []int {
	swapped := true
	for swapped {
		swapped = false

		for i := 0; i < len(x)-1; i++ {
			if x[i] > x[i+1] {
				x[i], x[i+1] = x[i+1], x[i]
				swapped = true
			}
		}
	}

	return x
}

func main() {
	unSorted := []int{5, 1, 4, 6, 9}
	fmt.Println(bubbleSort(unSorted))
}
