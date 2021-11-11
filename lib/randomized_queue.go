package alg

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

// RandomizedQueue
type RandomizedQueue struct {
	arr []int
	n   int
}

// NewRandomizedQueue creates a new randomized queue
func NewRandomizedQueue() *RandomizedQueue {
	rq := &RandomizedQueue{}
	rq.arr = make([]int, 2)

	// random seed with crypto/rand
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed with crypto rand")
	}
	crypto_rand.Read(b[:])
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
	return rq
}

// Is the queue empty?
func (rq *RandomizedQueue) IsEmpty() bool {
	return rq.n == 0
}

// Return the number of elements
func (rq *RandomizedQueue) Size() int {
	return rq.n
}

// Enqueue item
func (rq *RandomizedQueue) Enqueue(v int) {
	if rq.n == len(rq.arr) {
		rq.resize(2 * len(rq.arr))
	}
	rq.arr[rq.n] = v
	rq.n++
}

// Return and remove random item
func (rq *RandomizedQueue) Dequeue() int {
	if rq.IsEmpty() {
		panic("empty queue")
	}
	pick := rand.Intn(rq.n)

	v := rq.arr[pick]
	rq.arr[pick] = rq.arr[rq.n-1] // swap with last element
	rq.n--

	// shrink size of array
	if rq.n > 0 && rq.n == len(rq.arr)/4 {
		rq.resize(len(rq.arr) / 2)
	}

	return v
}

// resize the underlying array
func (rq *RandomizedQueue) resize(cap int) {
	dest := make([]int, cap)
	copy(dest, rq.arr)
	rq.arr = dest
}

// Return elements
func (rq *RandomizedQueue) Elems() []int {
	return rq.arr[:rq.n]
}
