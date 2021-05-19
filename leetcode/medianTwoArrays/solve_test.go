package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := map[string]float64{
		"1,3 2":                    2.0,
		"1,2 3,4":                  2.5,
		"0,0 0,0":                  0.0,
		" 1":                       1.0,
		"2 ":                       2.0,
		" 1,2,3,4":                 2.5,
		"1,3,3,4,4 1,1,2,2,3":      2.5,
		"1 1,2,2,2,2,3,4":          2.0,
		"0,0,0,0,0 -1,0,0,0,0,0,1": 0.0,
		"1,3 2,7":                  2.5,
		"2,3,4 1":                  2.5,
		"5,6 1,2,3,4,7,8,9":        5.0,
		"1,2 -1,3":                 1.5,
		"1,4 2,3,5,6,7,8":          4.5,
		"1,3,5 2,4,6,7,8":          4.5,
		"5,6,7 1,2,3,4,8,9,10":     5.5,
		"1,2,2 1,2,3":              2.0,
		"1,2,3 1,2,2":              2.0,
	}

	for tc, exp := range testCases {
		arrays := strings.Split(tc, " ")
		a := parseArray(arrays[0])
		b := parseArray(arrays[1])

		got := findMedianSortedArrays(a, b)
		if got != exp {
			t.Errorf("%s: exp: %f, got %f", tc, exp, got)
		}
	}
}

func parseArray(strArr string) []int {
	items := strings.Split(strArr, ",")
	var arr []int
	for i := 0; i < len(items); i++ {
		if len(items[i]) > 0 {
			n, _ := strconv.Atoi(items[i])
			arr = append(arr, n)
		}
	}
	return arr
}

func TestParseArray(t *testing.T) {
	a := parseArray("")
	if len(a) > 0 {
		t.Errorf("non-empty array: %v", a)
	}

	a = []int{1}
	a = a[0:0]
	fmt.Printf("%v", a)
}

func TestBinSearch(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	am := len(a) / 2
	bm := len(b) / 2
	amInd := sort.SearchInts(b, a[am])
	bmInd := sort.SearchInts(a, b[bm])

	exp := 2
	if amInd != exp {
		t.Errorf("exp %d, got %d", exp, amInd)
	}

	exp = 1
	if bmInd != exp {
		t.Errorf("exp %d, got %d", exp, bmInd)
	}
}
