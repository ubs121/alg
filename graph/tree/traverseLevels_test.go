package graph

import (
	"fmt"
	"testing"
)

type msg struct {
	elem int
	lvl  int
}

// Flatten tree by levels
func traverseLevels(root []any) {
	out := make(chan *msg)

	go func() {
		defer close(out)
		_traverse(root, 0, out)
	}()

	lvls := map[int][]int{}

	for {
		m := <-out
		if m == nil {
			break
		}

		lvls[m.lvl] = append(lvls[m.lvl], m.elem)
	}

	fmt.Printf("Answer %v", lvls)
}

func _traverse(root []any, lvl int, out chan *msg) {

	for _, e := range root {
		switch v := e.(type) {
		case int:
			out <- &msg{v, lvl}
		case []any:
			_traverse(v, lvl+1, out)
		default:
			// interface, so go deeper???
		}
	}

}

func TestTraverseLevels(t *testing.T) {
	root := []any{5, 12, []any{1, []any{8, 10, 2}, 1, 100}, []any{15}, 7}
	traverseLevels(root)
}
