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

func areBracketsBalanced(s string) bool {
	stack := NewStack()
	for i := range s {
		switch s[i] {
		case '(', '[', '{':
			stack.push(s[i])
		case ')':
			ch := stack.pop()
			if ch != '(' {
				return false
			}
		case ']':
			ch := stack.pop()
			if ch != '[' {
				return false
			}
		case '}':
			ch := stack.pop()
			if ch != '{' {
				return false
			}
		}
	}
	return stack.isEmpty()
}

func main() {
	input1 := "([])[]({})"
	fmt.Printf("Input: %v, Output: %t\n", input1, areBracketsBalanced(input1))

	input2 := "([)]"
	fmt.Printf("Input: %v, Output: %t\n", input2, areBracketsBalanced(input2))

	input3 := "((()"
	fmt.Printf("Input: %v, Output: %t\n", input3, areBracketsBalanced(input3))
}
