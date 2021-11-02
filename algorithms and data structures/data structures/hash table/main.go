package main

import (
	"fmt"
	"strings"
)

type bucketNode struct { // bucket node as linked list node
	data string
	next *bucketNode
}

type bucket struct { // bucket as linked list
	head *bucketNode
}

const hashTableLength = 10

type hashTable struct { // has table object
	array [hashTableLength]*bucket
}

// Insert for hash table

func (h *hashTable) Insert(data string) {
	tableIndex := hash(data) // hash data
	h.array[tableIndex].insert(data) // insert data into table
}

// Search for hash table

func (h *hashTable) Search(data string) bool {
	tableIndex := hash(data) // hash data
	return h.array[tableIndex].search(data) // search for data
}

// Delete for hash table

func (h *hashTable) Delete(data string) {
	tableIndex := hash(data) // hash data
	h.array[tableIndex].delete(data) // delete data
}

// insert into bucket
func (b *bucket) insert(data string) {
	dataLC := strings.ToLower(data)
	if !b.search(data) { // if data was not found, prepend data to bucket
		secondNode := b.head
		b.head = &bucketNode{data: dataLC}

		b.head.next = secondNode
	} else {
		fmt.Println("Data already exist")
	}
}

// search bucket
func (b *bucket) search(data string) bool {
	dataLC := strings.ToLower(data)
	firstNode := b.head

	for firstNode != nil {
		if firstNode.data == dataLC {
			return true
		}
		firstNode = firstNode.next
	}

	return false
}

//delete from bucket
func (b *bucket) delete(data string) {
	dataLC := strings.ToLower(data)

	if b.head.data == dataLC {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next.data != dataLC {
		if previousNode.next.next == nil {
			return
		}
		previousNode = previousNode.next
	}
	previousNode.next = previousNode.next.next
}

// hash function adds all the letter in data and return the remainder when divided by length of hash table
func hash(data string) int {
	dataLC := strings.ToLower(data)
	var sum int
	for _, v := range dataLC {
		sum += int(v)
	}

	return sum % hashTableLength
}

func initTable() *hashTable {
	table := &hashTable{}

	for i := range table.array {
		table.array[i] = &bucket{}
	}

	return table
}

func main() {
	h2 := hashTable{}
	h2.Insert("bruce")
	h2.Insert("clark")
	h2.Insert("diana")
	h2.Insert("barry")
	h2.Delete("arthur")
	fmt.Println(h2.Search("bruce"))
}
