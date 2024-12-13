package main

import "errors"

type Queue struct {
	head, tail *Qnode
}
type Qnode struct {
	data any
	next *Qnode
}

func (q *Queue) Enqueue(data any) {
	qn := &Qnode{data, nil}

	if q.tail != nil {
		q.tail.next = qn
	}
	q.tail = qn

	if q.head == nil {
		q.head = q.tail
	}
}

func (q *Queue) Dequeue() (any, error) {
	if q.head == nil {
		return nil, errors.New("dequeuing an Empty Queue")
	}

	dq := q.head.data
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}
	
	return dq, nil
}

func(q *Queue) Peek() (any, error) {
	if q.head == nil {
		return nil, errors.New("dequeuing an Empty Queue")
	}
	return q.head.data, nil
}

func(q *Queue) IsEmpty() bool {
	return q.head == nil
}