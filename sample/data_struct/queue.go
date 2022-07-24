package datastruct

import (
	"errors"
	"fmt"
)

var queueIsEmpty = errors.New("queue is empty")

type Queue struct {
	data []any
}

func newQueue() *Queue {
	return &Queue{}
}

func (q *Queue) push(value any) {
	q.data = append(q.data, value)
}

func (q *Queue) pop() (any, error) {
	length := len(q.data)
	if length == 0 {
		return nil, queueIsEmpty
	}

	value := q.data[0]
	q.data = removeArrayByIndex(q.data, 0)
	return value, nil
}

func (q *Queue) peek() (any, error) {
	length := len(q.data)
	if length == 0 {
		return nil, queueIsEmpty
	}

	return q.data[0], nil
}

func QueueMain() {
	queue := newQueue()
	queue.push(1)
	queue.push(2)
	queue.push(3)
	fmt.Println(queue.data...)

	v, _ := queue.pop()
	fmt.Println("pop:", v, "\tqueue:", queue)
	v, _ = queue.pop()
	fmt.Println("pop:", v, "\tqueue:", queue)
	v, _ = queue.pop()
	fmt.Println("pop:", v, "\tqueue:", queue)
	v, _ = queue.pop()
	fmt.Println("pop:", v, "\tqueue:", queue)

	queue.push(1)
	queue.push(2)
	queue.push(3)
	v, _ = queue.peek()
	fmt.Println("pop:", v, "\tqueue:", queue)
}
