package container

import (
	"container/list"
	"fmt"
	"testing"
)

func TestContainerList(t *testing.T) {
	// Create a new list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func TestDeque(t *testing.T) {
	dq := NewDeque()
	dq.AddFirst(1)
	dq.AddFirst(2)
	dq.AddFirst(3)
	dq.AddLast(4)
	dq.AddLast(5)
	dq.AddLast(6)

	got := dq.RemoveLast()
	exp := 6
	if got != exp {
		t.Errorf("exp %d, got %d", exp, got)
	}

	got = dq.RemoveFirst()
	exp = 3
	if got != exp {
		t.Errorf("exp %d, got %d", exp, got)
	}

	exp = 4
	if dq.Size() != exp {
		t.Errorf("exp %d, got %d", exp, dq.Size())
	}
}

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
