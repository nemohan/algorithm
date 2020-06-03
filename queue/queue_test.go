package queue

import (
	"fmt"
	"testing"
)

var q = NewQueue(5)

func addElement(t *testing.T, queue *Queue, num int) {
	for i := 0; i < num; i++ {
		err := q.Enqueue(i)
		if err != nil && q.Size() != 5 {
			q.Dump()
			t.Fatalf("%d size:%d %v \n", i, q.Size(), err)
		}
	}
}

func removeElement(t *testing.T, queue *Queue, num int) {
	for i := 0; i < num; i++ {
		_, err := q.Dequeue()
		if err != nil && q.Size() != 0 {
			t.Fatalf("remove %d size:%d\n", num, q.Size())
		}
	}
}

func TestEnqueue(t *testing.T) {
	addElement(t, q, 5)
}

func TestDequeue(t *testing.T) {
	removeElement(t, q, 5)
	addElement(t, q, 5)
	if q.Size() != 5 {
		q.Dump()
		t.Fatalf("expected size:5 got:%d front:%d rear:%d\n", q.Size(), q.front, q.rear)
	}
	fmt.Printf("front:%d rear:%d\n", q.front, q.rear)
	removeElement(t, q, 3)
	addElement(t, q, 3)
	if q.Size() != 5 {
		t.Fatalf("expected size:5 got:%d front:%d rear:%d\n", q.Size(), q.front, q.rear)
	}
	q.Dump()
}
