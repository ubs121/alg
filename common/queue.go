package common

// Double queue, similar to "container/list"
type Deque struct {
	arr   []int // array items
	n     int   // number of elements
	front int
	rear  int
}

// NewDeque creates a new double queue
func NewDeque() *Deque {
	cap := 2
	dq := &Deque{}
	dq.arr = make([]int, cap)
	dq.front = cap / 2
	dq.rear = cap / 2
	return dq
}

// Is the queue empty?
func (dq *Deque) IsEmpty() bool {
	return dq.n == 0
}

// Size returns the number of items
func (dq *Deque) Size() int {
	return dq.n
}

// resize the underlying array holding the elements
func (dq *Deque) resize(cap int) {
	middle := cap / 2
	newFront := middle - dq.n/2 // center it
	newRear := newFront + dq.n

	dest := make([]int, cap)
	copy(dest[newFront:], dq.arr)
	dq.arr = dest
	dq.front = newFront
	dq.rear = newRear
}

// Add the item to the front
func (dq *Deque) AddFirst(v int) {
	if dq.front == 0 {
		dq.resize(2 * len(dq.arr))
	}

	dq.front--
	dq.arr[dq.front] = v
	dq.n++
}

// Add the item to the end
func (dq *Deque) AddLast(v int) {
	if dq.rear == len(dq.arr) {
		dq.resize(2 * len(dq.arr))
	}

	dq.arr[dq.rear] = v
	dq.rear++
	dq.n++
}

// Remove and return front item
func (dq *Deque) RemoveFirst() int {
	v := dq.arr[dq.front]
	dq.front++
	dq.n--

	// downsize
	if dq.n > 0 && dq.n == len(dq.arr)/4 {
		dq.resize(len(dq.arr) / 2)
	}

	return v
}

// Remove and return an item from the end
func (dq *Deque) RemoveLast() int {
	dq.rear--
	v := dq.arr[dq.rear]
	dq.n--

	// shrink size of array
	if dq.n > 0 && dq.n == len(dq.arr)/4 {
		dq.resize(len(dq.arr) / 2)
	}

	return v
}

// Return elements
func (dq *Deque) Elems() []int {
	return dq.arr[:dq.n]
}

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
