package datastruct

import (
	"errors"
	"fmt"
)

var errStackIsEmpty = errors.New("stack is empty")

type stack struct {
	data []any
}

func NewStack() *stack {
	return &stack{}
}

func (s *stack) Push(value any) {
	s.data = append(s.data, value)
}

func (s *stack) Pop() (any, error) {
	length := len(s.data)
	if length == 0 {
		return nil, errStackIsEmpty
	}

	value := s.data[length-1]
	s.data = removeArrayByIndex(s.data, length-1)
	return value, nil
}

func (s *stack) Peek() (any, error) {
	length := len(s.data)
	if length == 0 {
		return nil, errStackIsEmpty
	}

	return s.data[length-1], nil
}

func (s *stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *stack) Show() {
	for i, d := range s.data {
		fmt.Printf("%d %v\n", i, d)
	}
}

func StackMain() {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	v, _ := stack.Pop()
	fmt.Println("pop:", v)
	fmt.Println(stack.data...)

	v, _ = stack.Peek()
	fmt.Println("peek:", v)

	fmt.Println("isEmpty:", stack.IsEmpty())
	stack.Pop()
	stack.Pop()
	fmt.Println("isEmpty:", stack.IsEmpty())
}
