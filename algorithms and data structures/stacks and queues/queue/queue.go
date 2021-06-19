package main

import (
	"fmt"
)

type Queue struct {
	data []int
}

func (s *Queue) Enqueue(value int) {
	s.data = append(s.data, value)
}

func (s *Queue) Dequeue() {
	s.data = s.data[1:]
}

func main() {
	s := &Queue{}

	s.Enqueue(2)
	s.Enqueue(3)
	s.Enqueue(4)
	s.Enqueue(5)
	s.Enqueue(6)
	s.Enqueue(7)
	fmt.Println(s)

	s.Dequeue()
	fmt.Println(s)

	s.Dequeue()
	fmt.Println(s)
}
