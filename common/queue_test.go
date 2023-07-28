package common

import (
	"testing"
)

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
