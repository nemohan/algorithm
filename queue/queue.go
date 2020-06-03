package queue

import (
	"errors"
	"fmt"
)

var (
	errQueueIsFull  = errors.New("queue is full")
	errQueueIsEmpty = errors.New("queue is empty")
)

type Queue struct {
	size  int
	data  []interface{}
	front int
	rear  int
	max   int
}

func NewQueue(size int) *Queue {
	return &Queue{
		size:  0,
		data:  make([]interface{}, size),
		max:   size,
		front: 0,
		rear:  -1,
	}
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.isEmpty() {
		return nil, errQueueIsEmpty
	}
	q.size--
	i := q.front
	q.front = (q.front + 1) % q.max
	return q.data[i], nil
}

func (q *Queue) Enqueue(v interface{}) error {
	if q.IsFull() {
		return errQueueIsFull
	}
	q.size++
	q.rear = (q.rear + 1) % q.max
	q.data[q.rear] = v
	return nil
}

func (q *Queue) IsFull() bool {
	return q.size == q.max
}

func (q *Queue) isEmpty() bool {
	return q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Dump() {
	if q.isEmpty() {
		return
	}
	fmt.Printf("================dump begin\n")
	fmt.Printf("front:%d rear:%d size:%d\n", q.front, q.rear, q.size)
	size := q.Size()
	i := q.front
	for size > 0 {
		fmt.Printf("%v\n", q.data[i])
		i = (i + 1) % q.max
		size--
	}
}
