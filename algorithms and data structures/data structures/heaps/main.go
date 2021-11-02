package main

import (
	"fmt"
	"math/rand"
	"time"
)

// heap struct
type Heap struct {
	slice []int
}

// insert key into heap
func (h *Heap) Insert(key int) {
	h.slice = append(h.slice, key) // add new key to end of slice
	h.heapifyUP(len(h.slice) - 1)  // heapify key
}

func (h *Heap) heapifyUP(index int) {
	for h.slice[parent(index)] < h.slice[index] { // if parent is less than left child
		h.Swap(parent(index), index) // swap
		index = parent(index)        // update for loop by moving up one node
	}
}

func (h *Heap) Extract() int {
	if len(h.slice) == 0 { // cannot extract from an empty slice
		fmt.Println("Empty heap")
		return -1
	}

	extracted := h.slice[0] // extracted value is first alue in slice

	lastIndexInHeap := len(h.slice) - 1
	h.slice[0] = h.slice[lastIndexInHeap] // push bottom value to top
	h.slice = h.slice[:lastIndexInHeap] // adjust slice

	h.heapifyDown(0) // heapify down

	return extracted 
}

func (h *Heap) heapifyDown(index int) {
	lastIndexInHeap := len(h.slice) - 1

	left, right := leftChild(index), rightChild(index)

	var childToCompare int

	for left <= lastIndexInHeap { // if at least index child is present i.e left index is the only index present
		if left == lastIndexInHeap {
			childToCompare = left
		} else if h.slice[left] > h.slice[right] { // if left child is greater than right child
			childToCompare = left
		} else { // if right child is greater
			childToCompare = right
		}

		if h.slice[index] < h.slice[childToCompare] { // if parent is smaller than child
			h.Swap(index, childToCompare) // swap
			index = childToCompare // update for loop
			left, right = leftChild(index), rightChild(index) // update for loop
		} else { // if parent is greater or equals child
			return
		}
	}
}

// get parent from left child index
func parent(leftIndex int) int {
	return (leftIndex - 1) / 2
}

// get left child from parent index
func leftChild(parentIndex int) int {
	return (parentIndex * 2) + 1
}

// get left child from parent index
func rightChild(parentIndex int) int {
	return (parentIndex * 2) + 2
}

// swap two node keys based on their index
func (h *Heap) Swap(index1, index2 int) {
	h.slice[index1], h.slice[index2] = h.slice[index2], h.slice[index1]
}

func main() {
	h := &Heap{}

	keys := []int{10, 20, 30, 40, 50, 60, 70}

	for _, v := range keys {
		h.Insert(v)
		fmt.Println(h)
	}

	fmt.Println("Done inserting, starting extraction")

	for i := 0; i < 4; i++ {
		fmt.Printf("%d: Extracted %d\n", i+1, h.Extract())
		fmt.Println(h)
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
}
