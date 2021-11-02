package main

import (
	"fmt"
)

type node struct {
	data  int
	left  *node
	right *node
}

//insert
func (n *node) insert(value int) {
	if n.data < value {
		//move right
		if n.right == nil {
			n.right = &node{data: value}
		} else {
			n.right.insert(value)
		}
	} else if n.data > value {
		//move left
		if n.left == nil {
			n.left = &node{data: value}
		} else {
			n.left.insert(value)
		}
	}
}

//search
var searchCount int

func (n *node) search(value int) (bool, int) {
	searchCount++

	if n == nil {
		return false, searchCount
	}

	if n.data < value {
		//move right
		return n.right.search(value)
	} else if n.data > value {
		//move left
		return n.left.search(value)
	}

	return true, searchCount
}

func main() {
	root := &node{data: 100}
	root.insert(50)
	root.insert(300)
	root.insert(450)
	root.insert(210)
	root.insert(90)
	root.insert(400)

	var data = 400
	found, count := root.search(data)
	if found {
		fmt.Printf("Found %d after %d nodes\n", data, count)
	}
}
