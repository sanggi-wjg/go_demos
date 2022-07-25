package channel

import (
	"fmt"
	"math/rand"
)

const QUEUE_SIZE = 20

var queue = newChannelQueue(QUEUE_SIZE)

type ChannelQueue struct {
	data chan any
}

func newChannelQueue(size int) *ChannelQueue {
	return &ChannelQueue{
		data: make(chan any, size),
	}
}

func (q *ChannelQueue) get() any {
	return <-q.data
}

func (q *ChannelQueue) set(value any) {
	q.data <- value
}

func setTest() {
	queue.set(rand.Intn(10))
}

func getTest() any {
	return queue.get()
}

func ChannelQueueMain() {
	setTest()
	fmt.Println(getTest())

	setTest()
	fmt.Println(getTest())

	setTest()
	fmt.Println(getTest())

	setTest()
	setTest()
	setTest()
	setTest()
	setTest()
	setTest()
	setTest()
	fmt.Println(getTest())
}
