// N тооноос К-аар зохиосон бүх боломжит хэсэглэл

package alg

import (
	"fmt"
	"testing"
)

var (
	n int /* нийт элементийн тоо */
	k int /* хэсэглэж авах элементийн тоо */
)

func TestCombo(t *testing.T) {
	fmt.Printf("n k ? ")
	fmt.Scanf("%d %d", &n, &k)

	/* хэсэглэлд зориулсан массив */
	a := make([]int, k+1)

	p := k
	/* эхний байрлал */
	for i := 1; i <= k; i++ {
		a[i] = i
	}

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
