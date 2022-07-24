package datastruct

import (
	"errors"
	"fmt"
)

var stackIsEmpty = errors.New("stack is empty")

type Stack struct {
	data []any
}

func newStack() *Stack {
	return &Stack{}
}

func (s *Stack) push(value any) {
	s.data = append(s.data, value)
}

func (s *Stack) pop() (any, error) {
	length := len(s.data)
	if length == 0 {
		return nil, stackIsEmpty
	}

	value := s.data[length-1]
	s.data = removeArrayByIndex(s.data, length-1)
	return value, nil
}

func (s *Stack) peek() (any, error) {
	length := len(s.data)
	if length == 0 {
		return nil, stackIsEmpty
	}

	return s.data[length-1], nil
}

func (s *Stack) isEmpty() bool {
	if len(s.data) == 0 {
		return true
	}
	return false
}

func StackMain() {
	stack := newStack()
	stack.push(1)
	stack.push(2)
	stack.push(3)

	v, _ := stack.pop()
	fmt.Println("pop:", v)
	fmt.Println(stack.data...)

	v, _ = stack.peek()
	fmt.Println("peek:", v)

	fmt.Println("isEmpty:", stack.isEmpty())
	stack.pop()
	stack.pop()
	fmt.Println("isEmpty:", stack.isEmpty())
}
