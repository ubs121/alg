package main

import (
	"fmt"
)

type Rect struct {
	left, top, right, bottom int
	sum                      int64
	index                    int // The index of the item in the heap.
}

func (r *Rect) contains(o *Rect) bool {
	if r.left <= o.left && r.top <= o.top && o.right <= r.right && o.bottom < r.bottom {
		return true
	}

	return false
}

// stringify
func (r *Rect) String() string {
	return fmt.Sprintf("(%d,%d)/(%d,%d)", r.top, r.left, r.bottom, r.right)
}

type BySum []*Rect

func (a BySum) Len() int           { return len(a) }
func (a BySum) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySum) Less(i, j int) bool { return a[i].sum > a[j].sum }
