package queue

import (
	"container/list"
)

type Q struct {
	list *list.List
}

func (q *Q) Enqueue(value interface{}) *Q {
	q.list.PushBack(value)
	return q
}

func (q *Q) Dequeue() *list.Element {
	element := q.list.Front()
	// fmt.Println("Removed", element.Value)
	q.list.Remove(element)
	return element
}

func (q *Q) Size() int {
	return q.list.Len()
}

func (q *Q) init() *Q {
	queue := list.New()
	q.list = queue
	return q
}

func New() *Q {
	return new(Q).init()
}
