// https://www.hackerrank.com/contests/w36/challenges/a-race-against-time
package main

import (
	"bufio"
	"fmt"
	"os"
)

// cr - carrier, sts - other students
func relay(cr *Student, sts []*Student) int {
	n := len(sts)
	fmt.Printf("%d -> %v\n", cr.id, sts)

	//fmt.Printf("cr=%v -> sts=%v\n", cr, sts)

	// hand over directly if only one student left
	if n == 1 {
		return cr.costTo(sts[n-1])
	}

	// find mandatory handovers, and [0:k1] + [k1:k2] + ... + [kj:n]
	sum := 0
	res := sts[:] // all

	for len(res) > 0 {
		j := 0
		for j < len(res) && cr.height >= res[j].height {
			j++
		}

		ho := j + 1 // mandatory handover student + 1
		if j >= len(res) {
			ho = len(res) // all last segment
		}

		target := res[ho-1]
		min := cr.costTo(target) // direct cost

		list := res[:ho]
		for len(list) > 0 {
			c := cr.costTo(list[0]) + relay(list[0], list[1:])

			if min > c {
				min = c
			}

			// remove
			list = list[1:]
		}

		//fmt.Printf("cost=%v\n", min)

		sum += min
		cr = target
		res = res[ho:]

	}

	return sum
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

	// calculate
	sum := relay(mason, sts[1:])
	fmt.Printf("%d\n", sum)
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
