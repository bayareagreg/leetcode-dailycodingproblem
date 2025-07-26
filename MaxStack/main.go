package main

import (
	"fmt"
)

type Stack struct {
	arr   []int
	maxes []int
}

func NewStack() Stack {
	return Stack{
		arr:   make([]int, 0),
		maxes: make([]int, 0),
	}
}

func printStack(s Stack) {
	fmt.Printf("Stack: ")
	for i := range s.arr {
		fmt.Printf(" %v", s.arr[i])
	}
	fmt.Println()
}

func (s *Stack) push(val int) {
	s.arr = append(s.arr, val)
	if len(s.maxes) == 0 {
		s.maxes = append(s.maxes, val)
	} else {
		s.maxes = append(s.maxes, max(val, s.maxes[len(s.maxes)-1]))
	}
}

func (s *Stack) pop() int {
	val := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	s.maxes = s.maxes[:len(s.maxes)-1]
	return val
}

func (s *Stack) max() int {
	return s.maxes[len(s.maxes)-1]
}

func main() {

	s := NewStack()
	s.push(5)
	s.push(7)
	s.push(3)

	printStack(s)

	v := s.pop()

	fmt.Printf("pop returned: %v\n", v)

	m := s.max()

	fmt.Printf("max returned: %v\n", m)

	s.pop()

	printStack(s)

	m1 := s.max()

	fmt.Printf("max returned: %v\n", m1)
}
