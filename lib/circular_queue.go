package alg

// Memory efficient queue implementation using circular array
type QueueCircular struct {
	data     []interface{}
	front    int
	rear     int
	capacity int
	size     int
}

// Creates a new queue
func NewQueueCircular(cap int) *QueueCircular {
	q := new(QueueCircular)
	q.capacity = cap
	q.size = 0
	q.front = -1
	q.rear = -1
	q.data = make([]interface{}, cap)
	return q
}

// Len returns the number of elements of queue
func (q *QueueCircular) Len() int {
	return q.size
}

// Checks if queue is empty
func (q *QueueCircular) IsEmpty() bool {
	return q.size == 0
}

// Checks if queue is full
func (q *QueueCircular) IsFull() bool {
	return q.size == q.capacity
}

func (q *QueueCircular) Front() interface{} {
	return q.data[q.front]
}

func (q *QueueCircular) Back() interface{} {
	return q.data[q.rear]
}

func (q *QueueCircular) EnQueue(v interface{}) {
	if q.IsFull() {
		panic("queue is full")
	}
	q.rear = (q.rear + 1) % q.capacity
	q.data[q.rear] = v
	if q.front == -1 {
		q.front = q.rear
	}
	q.size++
}

func (q *QueueCircular) DeQueue() interface{} {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	v := q.data[q.front]
	if q.front == q.rear {
		q.front = -1
		q.rear = -1
		q.size = 0
	} else {
		q.front = (q.front + 1) % q.capacity
		q.size--
	}

	return v
}
