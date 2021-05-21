package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

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

func cmpArr(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
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
