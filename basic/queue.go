package basic

import (
	"errors"
)

const maxQueueSize = 1000000000

type Queue struct {
	elements *LinkedList
	count    int32
}

func NewQueue() *Queue {
	return &Queue{elements: NewLinkedList(), count: 0}
}

func (q *Queue) Enqueue(data interface{}) {
	if q.count >= maxQueueSize {
		panic("queue count overflow")
	}

	q.elements.AppendToTail(data)
	q.count++
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.count == 0 {
		return nil, errors.New("no elements in queue")
	}

	head := q.elements.GetHead()
	q.elements.RemoveHead()

	return head.GetData()
}

func (q *Queue) IsEmpty() bool {
	return q.count == 0
}

func (q *Queue) String() string {
	return q.elements.String()
}
