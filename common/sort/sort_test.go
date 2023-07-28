package sort

import (
	"fmt"
	"reflect"
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

func TestCountingSort(t *testing.T) {

	// test input
	arr := []int{4, 5, 1, 2, 4, 1, 3}

	// expected result
	arrExp := make([]int, len(arr))
	copy(arrExp, arr)
	sort.Ints(arrExp) // sort using standard func

	/* массивыг эрэмбэлэх */
	maxVal := arrExp[len(arrExp)-1] + 1
	CountingSort(arr, maxVal)

	if !reflect.DeepEqual(arr, arrExp) {
		t.Error("not equal")
	}

	/* эрэмбэлэгдсэн  массивыг хэвлэж харуулах */
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
}
