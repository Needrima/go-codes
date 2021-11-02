package main

import "fmt"

func merge(arr, leftPart, rightPart []int) {
	newLeftPart := make([]int, len(leftPart))
	newRightPart := make([]int, len(rightPart))

	copy(leftPart, newLeftPart)
	copy(rightPart, newRightPart)

	for i,j,k := 0,0,0; i<len(newLeftPart) && j<len(newRightPart); k++ {
		if newLeftPart[i] < newRightPart[j] {
			arr[k] = newLeftPart[i]
			i++
		}else {
			arr[k] = newRightPart[j]
			j++
		}
	}
}

func mergeSort(arr []int) {
	if len(arr) > 1 {
		center := len(arr)/2
		leftPart := arr[:center]
		rightPart := arr[center:]

		mergeSort(leftPart)
		mergeSort(rightPart)

		merge(arr, leftPart, rightPart)
	}
}

func main() {
	arr := []int{3,-1,30,34,5,7,4,10,8,15}

	fmt.Println(arr)
	mergeSort(arr)

	fmt.Println(arr)
}