package alg

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortSlice(t *testing.T) {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	// sort 'people' in ascending order by name
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	// sort 'people' in ascending order by age
	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age:", people)
}

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
