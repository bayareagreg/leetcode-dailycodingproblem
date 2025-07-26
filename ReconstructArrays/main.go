package main

import (
	"fmt"
)

type Stack struct {
	arr []byte
}

func NewStack() Stack {
	return Stack{
		arr: make([]byte, 0),
	}
}

func printStack(s Stack) {
	fmt.Printf("Stack: ")
	for i := range s.arr {
		fmt.Printf(" %v", s.arr[i])
	}
	fmt.Println()
}

func (s *Stack) push(val byte) {
	s.arr = append(s.arr, val)
}

func (s *Stack) peek() byte {
	return s.arr[len(s.arr)-1]
}

func (s *Stack) pop() byte {
	val := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return val
}

func (s *Stack) isEmpty() bool {
	return len(s.arr) == 0
}

func reconstruct(array []byte) []int {
	result := make([]int, 0)
	stack := NewStack()

	for i := 0; i < len(array)-1; i++ {
		if array[i+1] == '-' {
			stack.push(byte(i))
		} else {
			result = append(result, i)
			for !stack.isEmpty() {
				result = append(result, int(stack.pop()))
			}
		}
	}
	stack.push(byte(len(array) - 1))
	for !stack.isEmpty() {
		result = append(result, int(stack.pop()))
	}

	return result
}

func main() {
	//input0 := []byte{0, '+', '+', '+', '+'}
	//fmt.Printf("Reconstructed array: %v\n", reconstruct(input0))

	//	input1 := []byte{0, '+', '+', '-', '+'}
	//	fmt.Printf("Reconstructed array: %v\n", reconstruct(input1))

	input2 := []byte{0, '-', '-', '-', '+'}
	fmt.Printf("Reconstructed array: %v\n", reconstruct(input2))

}
