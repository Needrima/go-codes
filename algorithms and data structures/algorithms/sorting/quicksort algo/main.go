package main

import (
	"math/rand"
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left]) 
	quickSort(arr[left+1:])

	return arr
}

func main() {
	arr := make([]int, 10, 10)

	for i := range arr {
		arr[i] = rand.Intn(len(arr))
	}


	fmt.Printf("Unsorted: %v\n", arr)

	fmt.Printf("Sorted: %v\n", quickSort(arr))
}