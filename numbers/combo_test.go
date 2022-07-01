// N тооноос К-аар зохиосон бүх боломжит хэсэглэл

package numbers

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

	// collect it
	cmb := make([]int, k)
	copy(cmb, a[1:])
	ret = append(ret, cmb)

	p := k
	for a[1] < n-k+1 {

		/* дараагийн хэсэглэлийг зохиох */
		if a[k] == n {
			p--
		} else {
			p = k
		}

		for i := k; i >= p; i-- {
			a[i] = a[p] + (i - p + 1)
		}

		/* шинэ байрлал */
		cmb := make([]int, k)
		copy(cmb, a[1:])
		ret = append(ret, cmb)
	}

	return ret
}

func TestCombo(t *testing.T) {
	n := 10
	k := 10
	ret := combine(n, k)
	if len(ret) != fact(n)/(fact(n-k)*fact(k)) {
		t.Errorf("can't match")
	}
	fmt.Printf("%v", ret)
}

func TestComboN(t *testing.T) {
	n := 20
	facts := map[int]int{}
	facts[1] = 1
	facts[2] = 2
	for i := 3; i <= n; i++ {
		facts[i] = fact(i)
	}

	fmt.Printf("%v", facts)
}
func fact(n int) int {
	p := 1
	for i := 2; i <= n; i++ {
		p = p * i
	}
	return p
}

func combine2(n int, k int) [][]int {
	a := make([]int, k)

	// initialize first combination
	for i := 0; i < k; i++ {
		a[i] = i
	}

	i := k - 1

	var ret [][]int

	for a[0] < n-k+1 {
		for i > 0 && a[i] == n-k+i {
			i--
		}

		// collect it
		cmb := make([]int, k)
		copy(cmb, a)
		ret = append(ret, cmb)

		a[i]++

		for i < k-1 {
			a[i+1] = a[i] + 1
			i++
		}
	}
	return ret
}
