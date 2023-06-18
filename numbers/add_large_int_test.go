package numbers

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestAddSimple(t *testing.T) {
	num1 := []int{}
	num2 := []int{}

	num1length := 6
	for i := 1; i <= num1length; i++ {
		num1 = append(num1, i)
	}

	num2length := 6
	for i := 1; i <= num2length; i++ {
		num2 = append(num2, i)
	}

	result := Add(num1, num2)

	fmt.Println("Result:", result)
}

func TestAdd(t *testing.T) {
	// []{num1,num2,expected}
	testCases := [][]string{
		{"123456", "123456", "777777"},
		{"", "", ""},
		{"", "1", "1"},
		{"99", "0", "99"},
		{"", "99", "99"},
		{"99", "10", "100"},
		{"99", "1", "100"},
		{"999", "100", "1000"},
		{"999", "101", "1100"},
		{"999", "110", "1010"},
		{"333", "22", "355"},
		{
			"2222222222222222222222222222222222222222222",
			"2222222222222222222222222222222222222222222",
			"4444444444444444444444444444444444444444444",
		},
		{
			"9999999999999999999999999999999999999999999",
			"100",
			"10000000000000000000000000000000000000000000",
		},
	}

	for i, tc := range testCases {
		got := Add(str2intArray(tc[0]), str2intArray(tc[1]))
		exp := tc[2]
		if got != exp {
			t.Errorf("tc %d: exp %s, got %s", i, exp, got)
		}
	}
}

func TestReverse(t *testing.T) {
	// all test cases
	testCases := map[*[]int][]int{
		{1, 2, 3, 4, 5, 6}:    {6, 5, 4, 3, 2, 1},
		{7, 7, 7, 7, 7, 7, 7}: {7, 7, 7, 7, 7, 7, 7},
		{}:                    {},
		{1}:                   {1},
	}

	tc := 1
	for inp, exp := range testCases {
		got := Reverse(*inp)
		if !reflect.DeepEqual(got, exp) {
			t.Errorf("tc %d: failed", tc)
		}
		tc++
	}
}

func TestReverseLarge(t *testing.T) {
	largeIn1, largeOut1 := generateLargeTestCaseForReverse(1000000000)

	// all test cases
	testCases := map[*[]int][]int{
		&largeIn1: largeOut1,
	}

	start := time.Now()
	tc := 1
	for inp, exp := range testCases {
		// batch size per cpu, assuming 4 core is common
		batchSize := len(*inp) / 4
		got := ReverseInBatches(*inp, batchSize)
		if !reflect.DeepEqual(got, exp) {
			t.Errorf("tc %d: failed", tc)
		}
		tc++
	}
	dur := time.Since(start)
	fmt.Printf("Duration %v", dur)
}

// helper func that generates a large input for testing
// size - number of digits
func generateLargeTestCaseForReverse(size int) ([]int, []int) {
	orig := make([]int, size)
	reversed := make([]int, size)

	for i := 0; i < size; i++ {
		orig[i] = (i + 1) % 10 // single digit
	}

	for i := 0; i < size; i++ {
		reversed[i] = (size - i) % 10
	}
	return orig, reversed
}
