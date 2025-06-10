package queue

import "errors"

// Queue implements a FIFO linear data structure.
type Queue struct {
	front *QueueNode
	back  *QueueNode
}

func (q *Queue) IsEmpty() bool {
	return q.front == nil
}

func (q *Queue) Peek() (any, error) {
	if q.front == nil {
		return nil, errors.New("queue is empty; cannot peek")
	}
	return q.front.value, nil
}

func (q *Queue) Add(v any) {
	node := NewQueueNode(v)
	if q.back != nil {
		q.back.next = node
	}
	q.back = node
	if q.front == nil {
		q.front = node
	}
}

func (q *Queue) Remove() (any, error) {
	if q.front == nil {
		return nil, errors.New("queue is empty; cannot remove")
	}
	v := q.front.value
	q.front = q.front.next
	if q.front == nil {
		q.back = nil
	}
	return v, nil
}

type QueueNode struct {
	value any
	next  *QueueNode
}

func NewQueueNode(v any) *QueueNode {
	return &QueueNode{value: v}
}
