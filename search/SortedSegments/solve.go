// https://www.hackerrank.com/challenges/sorted-subsegments
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// segment type
type Segment struct {
	id   int // query id
	l, r int // segment limits
}

// query list
type QueryList []Segment

func main() {
	file, _ := os.Open("input13.txt")
	defer file.Close()

	buf := make([]byte, 2<<(10*2))    // 2 Mb buffer
	scanner := bufio.NewScanner(file) // file, os.Stdin
	scanner.Buffer(buf, len(buf))

	var n, q, k int // 1 <= n,q <= 75000

	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d %d %d\n", &n, &q, &k)

	// read data array
	doneParsing := make(chan bool)
	a := make([]int, n)
	scanner.Scan()
	line := scanner.Text() // as single line

	// async parsing
	go parseArray(doneParsing, line, a, n)

	// read queries
	qlist := make(QueryList, q)
	for i := 0; i < q; i++ {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d\n", &qlist[i].l, &qlist[i].r)
		qlist[i].id = i
	}

	// clean the query list
	chain := clean(k, qlist)

	//fmt.Println(chain)

	// wait for data
	<-doneParsing

	// apply queries on the data
	var curr, next Segment
	q = len(chain) - 1 // exclude last query
	for i := 0; i < q; i++ {
		curr = chain[i]
		next = chain[i+1]

		//sort.Ints(a[curr.l : curr.r+1])

		if curr.l < next.l {
			partial_sort_left(a, curr.l, curr.r, next.l)
		} else {
			partial_sort_right(a, curr.l, curr.r, next.r)
		}
	}

	// last query
	partial_sort_left(a, chain[q].l, chain[q].r, k)

	// print value at position k
	fmt.Println(a[k])
}

// https://en.wikipedia.org/wiki/Partial_sorting
func partial_sort_left(a []int, i, j, k int) {
	if i < j {
		p := (i + j) / 2 //i + rand.Intn(j-i+1)
		p = partition(a, i, j, p)

		partial_sort_left(a, i, p-1, k)

		if p < k {
			partial_sort_left(a, p+1, j, k)
		}
	}
}

func partial_sort_right(a []int, i, j, k int) {
	if i < j {
		p := (i + j) / 2
		p = partition(a, i, j, p)

		partial_sort_right(a, p+1, j, k)

		if p > k-1 {
			partial_sort_right(a, i, p-1, k)
		}
	}
}

// https://en.wikipedia.org/wiki/Quickselect
func partition(a []int, l, r, p int) int {
	v := a[p]

	// Move pivot to end
	a[p], a[r] = a[r], a[p]

	i := l
	for j := l; j < r; j++ {
		if v > a[j] {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}

	// Move pivot to its final place
	a[r], a[i] = a[i], a[r]

	return i
}

// Clean query list (remove unrelated segments, remove duplicates etc.)
func clean(k int, qlist QueryList) QueryList {
	var chain QueryList
	last := -1

	/*
		Find related segments with 'k', and remove unrelated
	*/
	seg := Segment{0, -1, -1}
	for i := len(qlist) - 1; i >= 0; i-- {
		if last < 0 {
			// find the highest order query that contains 'k'
			if qlist[i].l <= k && k <= qlist[i].r {
				last = i
				seg = qlist[i]
				chain = append(chain, qlist[i])
			}
		} else {
			if qlist[i].r < seg.l || seg.r < qlist[i].l {
				// no intersection, so skip
				continue
			} else {
				if seg.r < qlist[i].r {
					seg.r = qlist[i].r // extend to right
				}
				if seg.l > qlist[i].l {
					seg.l = qlist[i].l // extend to left
				}

				chain = append(chain, qlist[i])
			}
		}
	}

	// sort by segment
	sort.Slice(chain, func(i, j int) bool {
		if chain[i].l == chain[j].l {
			return chain[i].r < chain[j].r
		}

		return chain[i].l < chain[j].l
	})

	/*
		Remove duplicated segments
	*/
	var chain2 QueryList

	chain2 = append(chain2, chain[0])
	seg = chain[0]
	for i := 1; i < len(chain); i++ {
		if seg.l == chain[i].l {
			chain2[len(chain2)-1] = chain[i] // replace previuos with current
			seg = chain[i]
		} else if seg.r < chain[i].r { // seg.l <= chain[i].l
			chain2 = append(chain2, chain[i])
			seg = chain[i]
		} else {
			// small segment, so skip it
		}
	}

	// put in the order
	sort.Slice(chain2, func(i, j int) bool { return chain2[i].id < chain2[j].id })

	return chain2

}

// parse integer array
func parseArray(done chan bool, line string, a []int, n int) {
	scanner2 := bufio.NewScanner(strings.NewReader(line))
	scanner2.Split(bufio.ScanWords) // now tokenize

	for i := 0; i < n; i++ {
		scanner2.Scan()
		fmt.Sscanf(scanner2.Text(), "%d", &a[i])
	}

	done <- true
}
