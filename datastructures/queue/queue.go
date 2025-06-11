package queue

import "errors"

var ErrQueueIsEmpty = errors.New("queue is empty")

// Queue implements a FIFO linear data structure.
type Queue struct {
	front *QueueItem
	back  *QueueItem
}

// IsEmpty returns 'true' when there are no items in the Queue.
func (q *Queue) IsEmpty() bool {
	return q.front == nil
}

// Peek returns the first item in the Queue, or an error if the Queue is empty.
func (q *Queue) Peek() (any, error) {
	if q.IsEmpty() {
		return nil, ErrQueueIsEmpty
	}
	return q.front.value, nil
}

// Add inserts the given item into the Queue.
func (q *Queue) Add(v any) {
	item := NewQueueItem(v)
	if q.back != nil {
		q.back.next = item
	}
	q.back = item
	if q.IsEmpty() {
		q.front = item
	}
}

// Remove returns the first item from the Queue, or an error if the Queue is empty.
func (q *Queue) Remove() (any, error) {
	if q.IsEmpty() {
		return nil, ErrQueueIsEmpty
	}
	v := q.front.value
	q.front = q.front.next
	if q.IsEmpty() {
		q.back = nil
	}
	return v, nil
}

// QueueItem represents an item in the Queue. It contains a value, as well as a reference to the next item.
type QueueItem struct {
	value any
	next  *QueueItem
}

// NewQueueItem stores the given value in a QueueItem before returning its reference.
func NewQueueItem(v any) *QueueItem {
	return &QueueItem{value: v}
}
