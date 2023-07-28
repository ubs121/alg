package common

import (
	"fmt"
	"testing"
)

func TestRandomizedQueue(t *testing.T) {
	rq := NewRandomizedQueue()
	rq.Enqueue(1)
	rq.Enqueue(2)
	rq.Enqueue(3)
	rq.Enqueue(4)
	rq.Enqueue(5)

	exp := 5
	if rq.Size() != exp {
		t.Errorf("exp %d, got %d", exp, rq.Size())
	}

	for !rq.IsEmpty() {
		e := rq.Dequeue()
		fmt.Println(e)
	}
}
