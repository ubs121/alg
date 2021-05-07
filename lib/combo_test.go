// N тооноос К-аар зохиосон бүх боломжит хэсэглэл

package alg

import (
	"fmt"
	"testing"
)

func combine(n int, k int) [][]int {
	/* хэсэглэлд зориулсан массив */
	a := make([]int, k+1)

	/* эхний байрлал */
	for i := 1; i <= k; i++ {
		a[i] = i
	}

	var ret [][]int

	p := k
	for p > 0 {
		/* шинэ байрлал */
		cmb := make([]int, k)
		copy(cmb, a[1:])
		ret = append(ret, cmb[:])

		/* дараагийн хэсэглэлийг зохиох */
		if a[k] == n {
			p--
		} else {
			p = k
		}

		if p > 0 {
			for i := k; i >= p; i-- {
				a[i] = a[p] + (i - p + 1)
			}
		}
	}

	return ret
}

func TestCombo(t *testing.T) {
	n := 20
	k := 5
	ret := combine(n, k)
	if len(ret) != fact(n)/(fact(n-k)*fact(k)) {
		t.Errorf("can't match")
	}
	fmt.Printf("%v", ret)
}

func fact(n int) int {
	p := 1
	for i := 2; i <= n; i++ {
		p = p * i
	}
	return p
}
