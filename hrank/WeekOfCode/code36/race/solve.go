// https://www.hackerrank.com/contests/w36/challenges/a-race-against-time
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// fast short path algorithm
func minCost(sts []*Student) int64 {
	n := len(sts)
	mason := sts[0]
	madison := sts[n-1]

	// initiliaze
	Q := make([]*Student, 1)
	Q[0] = mason // start with Mason

	inQueue := make(map[int]bool, n)
	inQueue[mason.id] = true

	distTo := make([]int64, n)
	prev := make([]int, n)

	for i := 1; i < n; i++ {
		distTo[i] = math.MaxInt64 // infinity
		prev[i] = -1              // undefined
	}

	relax := func(u, v *Student) {
		alt := distTo[u.id] + int64(u.costTo(v))

		//printPath(v.id, prev, sts)

		if alt < distTo[v.id] {
			distTo[v.id] = alt
			prev[v.id] = u.id

			// put in the queue & mark
			if !inQueue[v.id] {
				Q = append(Q, v)
				inQueue[v.id] = true
			}
		}

	}

	for len(Q) > 0 {
		// dequeue minimum item
		u := Q[0]
		Q = Q[1:]
		delete(inQueue, u.id)

		//printPath(u.id, prev, sts)

		// TODO: optimize between mandatory points

		// for each neighbor of 'u' until the mandatory point
		j := u.id + 1
		for j < n && u.height >= sts[j].height {
			relax(u, sts[j])
			j++
		}

		if j < n { // the mandatory point
			relax(u, sts[j])
		}
	}

	return distTo[madison.id]
}

func main() {
	var n, H int
	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%d\n", &H)

	sts := make([]*Student, n+1)
	mason := &Student{id: 0, height: H, price: 0}                 // mason
	madison := &Student{id: n, height: 0, price: 0, target: true} // madison
	sts[0] = mason
	sts[n] = madison

	reader := bufio.NewReader(os.Stdin)
	buf, _ := reader.ReadBytes('\n')
	p := 0
	// height
	for i := 1; i < n; i++ {
		sts[i] = &Student{id: i}
		sts[i].height, p = readInt(buf)
		buf = buf[p:]
	}

	buf, _ = reader.ReadBytes('\n')
	p = 0
	// price
	for i := 1; i < n; i++ {
		sts[i].price, p = readInt(buf)
		buf = buf[p:]
	}

	fmt.Printf("%d\n", minCost(sts))
}

func readInt(buf []byte) (int, int) {
	i := 0
	l := len(buf)
	for i < l && ' ' == buf[i] {
		i++
	}

	d := 0
	neg := false

	if buf[i] == '-' {
		neg = true
		i++
	}

	for i < l && '0' <= buf[i] && buf[i] <= '9' {
		d = d*10 + int(buf[i]-'0')
		i++
	}
	if neg {
		d = -1 * d
	}

	return d, i
}

func printPath(final int, prev []int, sts []*Student) {
	p := final
	fmt.Printf("[ ")
	for p > 0 {
		fmt.Printf("%d (%d)<- ", sts[p].height, p)
		p = prev[p]
	}
	fmt.Println("*]")
}
