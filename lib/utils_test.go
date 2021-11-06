package alg

import (
	"fmt"
	"testing"
)

func TestParseArray(t *testing.T) {
	a := ParseIntArray("")
	if len(a) > 0 {
		t.Errorf("non-empty array: %v", a)
	}

	a = []int{1}
	a = a[0:0]
	fmt.Printf("%v", a)
}
