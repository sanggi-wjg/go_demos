package datastruct

import (
	"fmt"
)

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

type Node struct {
	Data any
	Next *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

func newNode(value any) *Node {
	return &Node{
		Data: value,
		Next: nil,
	}
}

func (list *LinkedList) Lpush(value any) {
	node := newNode(value)

	if list.Head == nil {
		list.Head = node
		list.Tail = node

	} else {
		node.Next = list.Head
		list.Head = node
	}
	list.Size += 1
}

func (list *LinkedList) Rpush(value any) {
	node := newNode(value)

	if list.Tail == nil {
		list.Head = node
		list.Tail = node

	} else {
		list.Tail.Next = node
		list.Tail = node
	}
	list.Size += 1
}

func (list *LinkedList) ShowNodes() {
	if list.Size == 0 {
		fmt.Println("LinkedList is empty")

	} else {
		node := list.Head
		res := fmt.Sprintf("%v → ", node.Data)

		for i := 1; i < list.Size; i++ {
			node = node.Next
			res += fmt.Sprintf("%v → ", node.Data)
		}
		fmt.Println(res)
	}
}

func (list *LinkedList) Lpop() *Node {
	if list.Size == 0 {
		return nil

	} else if list.Size == 1 {
		head := list.Head
		list.Head = nil
		list.Tail = nil
		list.Size -= 1
		return head

	} else {
		head := list.Head
		list.Head = head.Next
		list.Size -= 1
		return head
	}
}

func (list *LinkedList) Rpop() *Node {
	if list.Size == 0 {
		return nil

	} else if list.Size == 1 {
		tail := list.Tail
		list.Head = nil
		list.Tail = nil

		list.Size -= 1
		return tail

	} else {
		tail := list.Tail

		node := list.Head
		for i := 1; i < list.Size; i++ {
			node = node.Next
			if i == list.Size-2 {
				list.Tail = node
				list.Tail.Next = nil
			}
		}

		list.Size -= 1
		return tail
	}
}

func LinkedListMain() {
	linkedList := NewLinkedList()

	// lpush
	linkedList.Lpush(1)
	linkedList.Lpush(2)
	linkedList.Lpush(3)

	// rpush
	linkedList.Rpush(4)
	linkedList.Rpush(5)
	linkedList.Rpush(6)

	// show
	linkedList.ShowNodes()

	// lpop
	linkedList.Lpop()
	linkedList.Lpop()
	linkedList.ShowNodes()

	// rpop
	linkedList.Rpop()
	linkedList.Rpop()
	linkedList.ShowNodes()

	// pop
	linkedList.Lpop()
	linkedList.ShowNodes()
	linkedList.Rpop()
	linkedList.ShowNodes()

	// push
	linkedList.Rpush(4)
	linkedList.Rpush(5)
	linkedList.ShowNodes()

	// pop
	linkedList.Rpop()
	linkedList.Rpop()
	linkedList.ShowNodes()
}
