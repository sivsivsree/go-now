package queue

import (
	"container/list"
)

// Q data structure
type Q struct {
	list *list.List
}

// New creates a new Queue
func New() *Q {
	return new(Q).init()
}

// Enqueue used to push a value to the Queue
func (q *Q) Enqueue(value interface{}) *Q {

	q.list.PushBack(value)

	return q
}

// Dequeue is used to Pop a value from the Queue and
// it returns Element
func (q *Q) Dequeue() *list.Element {

	if q.list.Len() > 0 {
		element := q.list.Front()
		// fmt.Println("Removed", element.Value)
		q.list.Remove(element)
		return element
	}

	return nil
}

// Size gives the length of the items in the queue
func (q *Q) Size() int {
	return q.list.Len()
}

func (q *Q) init() *Q {
	queue := list.New()
	q.list = queue
	return q
}
