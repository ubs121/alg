// N тооноос К-аар зохиосон бүх боломжит хэсэглэл

package alg

import (
	"testing"
)

func TestCombo(t *testing.T) {
	n := 10 /* нийт элементийн тоо */
	k := 3  /* хэсэглэж авах элементийн тоо */

	/* хэсэглэлд зориулсан массив */
	a := make([]int, k+1)

	/* эхний байрлал */
	for i := 1; i <= k; i++ {
		a[i] = i
	}

	p := k
	for p > 0 {
		/* шинэ байрлал */
		for i := 1; i <= k; i++ {
			print(a[i], " ")
		}
		println()

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
}
