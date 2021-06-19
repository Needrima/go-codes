package main

import (
	"fmt"
)

type Stack struct {
	data []int
}

func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() {
	lastValue := s.data[len(s.data)-1]

	s.data = s.data[:lastValue-2]
}

func main() {
	s := &Stack{}

	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Push(7)
	fmt.Println(s)

	s.Pop()
	fmt.Println(s)
	
	s.Pop()
	fmt.Println(s)
}
