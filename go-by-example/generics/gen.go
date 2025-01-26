package main

import (
	"fmt"
)

// func Max[T int | float32](a, b T) T {
// 	return max(a, b)
// }

// func Map[S comparable](input []S, ele S) (bool, error) {
// 	for _, value := range input {
// 		if value == ele {
// 			return true, nil
// 		}
// 	}
// 	return false, nil
// }

type Stack[T comparable] struct {
	stack []T
}

func (s *Stack[T]) push(ele T) {
	s.stack = append(s.stack, ele)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.stack) == 0 {
		var zero T
		return zero, true
	}
	size := len(s.stack)
	ele := s.stack[size-1]
	s.stack = s.stack[:size-1]
	return ele, false
}

func main() {
	s := Stack[int]{}
	s.push(1)
	s.push(3)
	s.push(5)
	fmt.Println("Stack after pushing -->", s.stack)

	ele, err := s.pop()
	if err {
		fmt.Println("Error: Stack is empty")
	} else {
		fmt.Println("Popped element -->", ele)
		fmt.Println("Stack after popping -->", s.stack)
	}

	_, _ = s.pop()
	_, _ = s.pop()
	ele, err = s.pop() // Attempt to pop from an empty stack
	if err {
		fmt.Println("Error: Stack is empty after popping everything")
	} else {
		fmt.Println("Popped element -->", ele)
	}
}
