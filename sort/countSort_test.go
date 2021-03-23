package sort

import (
	"fmt"
	"testing"
)

func TestCountSort(t *testing.T) {

	// sample array
	array := []int{4, 5, 1, 2, 4, 1, 3}

	/* массивыг эрэмбэлэх */
	countSort(array)

	/* эрэмбэлэгдсэн  массивыг хэвлэж харуулах */
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d ", array[i])
	}
}
