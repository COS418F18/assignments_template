package chandy_lamport

import "container/list"

// Define a queue -- simple implementation over List
type Queue struct {
	elements *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) Empty() bool {
	return (q.elements.Len() == 0)
}

func (q *Queue) Push(v interface{}) {
	q.elements.PushFront(v)
}

func (q *Queue) Pop() interface{} {
	return q.elements.Remove(q.elements.Back())
}

func (q *Queue) Peek() interface{} {
	return q.elements.Back().Value
}
