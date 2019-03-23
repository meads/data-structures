package queue

// QAble ...
type QAble interface {
	Enqueue(interface{})
	Dequeue() interface{}
	Empty() bool
}

// Queue is First In First Out (FIFO)
type Queue struct {
	items []interface{}
}

// NewQueue constructs an instance of Queue structure
func NewQueue() QAble {
	items := []interface{}{}
	return &Queue{
		items: items,
	}
}

// Enqueue given an item will add it to the rear of the queue
func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item found at the front of the queue
func (q *Queue) Dequeue() interface{} {
	var item interface{}
	item, q.items = q.items[0], q.items[1:]
	return item
}

// Empty returns a boolean value indicating if the queue has zero length currently
func (q *Queue) Empty() bool {
	return len(q.items) == 0
}
