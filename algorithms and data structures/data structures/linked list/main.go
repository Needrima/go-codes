package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

//print data

func (l *linkedList) print() {
	for l.head != nil {
		fmt.Println(l.head.data)
		l.head = l.head.next
	}
}

//prepend
func (l *linkedList) prepend(n *node) {
	second := l.head

	l.head = n

	l.head.next = second

	l.length++
}

//append
func (l *linkedList) append(n *node) {
	start := l.head
	for start.next != nil {
		start = start.next
	}

	start.next = n

	l.length++
}

func (l *linkedList) search(value int) bool {
	start := l.head 

	for start != nil {
		if start.data == value {
			return true
		}
		start = start.next
	}

	return false
}

//delete
func (l *linkedList) deleteWIthData(value int) {
	//if list is empty
	if l.length == 0 {
		fmt.Println("Empty list")
		return
	}

	// if data is in head
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil { //check if value exist
			fmt.Println("Value not found")
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--

}

func main() {
	list := &linkedList{}

	node1 := &node{data: 100}
	node2 := &node{data: 200}
	node3 := &node{data: 300}
	node4 := &node{data: 400}
	node5 := &node{data: 500}

	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)
	list.prepend(node4)
	list.prepend(node5)
	list.deleteWIthData(500)
	list.print()

	/*
		node6 := &node{data: 600}
		node7 := &node{data: 700}
		node8 := &node{data: 800}
		node9 := &node{data: 900}
		node10 := &node{data: 1000}

		list.append(node6)
		list.append(node7)
		list.append(node8)
		list.append(node9)
		list.append(node10)
	*/
}
