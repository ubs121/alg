package search

import (
	"fmt"
	"testing"
)

type msg struct {
	elem int
	lvl  int
}

// Flatten tree by levels
func traverseLevels(root []interface{}) {
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

func _traverse(root []interface{}, lvl int, out chan *msg) {

	for _, e := range root {
		switch v := e.(type) {
		case int:
			out <- &msg{v, lvl}
		case []interface{}:
			_traverse(v, lvl+1, out)
		default:
			// interface, so go deeper???
		}
	}

}

func TestTraverseLevels(t *testing.T) {
	root := []interface{}{5, 12, []interface{}{1, []interface{}{8, 10, 2}, 1, 100}, []interface{}{15}, 7}
	traverseLevels(root)
}
