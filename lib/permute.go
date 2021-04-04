// N ширхэг тоогоор зохиох бүх боломжит сэлгэмэл
package alg

import (
	"fmt"
	"testing"
)

var (
	a [100]int /* сэлгэх утгуудыг агуулах массив */
	N int      /* сэлгэх элементийн тоо */
)

func TestPermute(t *testing.T) {
	print("N=")
	fmt.Scanf("%d", &N)

	/* эхний хувилбар */
	for i := 0; i < N; i++ {
		a[i] = i + 1
	}
	display()

	p := make([]int, N) /* сэлгэмлийг удирдах массив */
	var j int

	for i := 1; i < N; {
		if p[i] < i {
			if i%2 > 0 {
				j = p[i]
			} else {
				j = 0
			}
			// swap
			a[j], a[i] = a[i], a[j]

			/* сэлгэмлийн  шинэ хувилбарыг хэвлэх */
			display()

			p[i]++
			i = 1
		} else {
			p[i] = 0
			i++
		}
	}
}

// туслах функц
func display() {
	for i := 0; i < N; i++ {
		print(a[i], " ")
	}
	println()
}
