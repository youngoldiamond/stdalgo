package queue

import (
	"errors"
)

type Queue struct {
	data       []int
	head, tail int
}

func New(capacity int) (*Queue, error) {
	if capacity < 1 {
		return nil, errors.New("invalid capacity")
	}
	return &Queue{make([]int, capacity+1), 0, 0}, nil
}

func (q *Queue) Push(key int) {
	if q.Size() == (len(q.data) - 1) {
		panic("queue owerflow")
	}
	q.data[q.tail] = key
	q.tail = (q.tail + 1) % len(q.data)
}

func (q *Queue) Pop() int {
	if q.tail == q.head {
		panic("queue is empty")
	}
	ans := q.data[q.head]
	q.head = (q.head + 1) % len(q.data)
	return ans
}

func (q *Queue) Size() int {
	if q.tail >= q.head {
		return q.tail - q.head
	} else {
		return len(q.data) - q.head + q.tail
	}
}

func (q *Queue) Empty() bool {
	return q.tail == q.head
}
