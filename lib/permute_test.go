// N ширхэг тоогоор зохиох бүх боломжит сэлгэмэл
package alg

import (
	"testing"
)

func TestPermute(t *testing.T) {
	a := []int{1, 2, 3} /* сэлгэх утгуудыг агуулах массив */
	N := len(a)         /* сэлгэх элементийн тоо */

	display(a)

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
			display(a)

			p[i]++
			i = 1
		} else {
			p[i] = 0
			i++
		}
	}
}

// туслах функц
func display(arr []int) {
	for i := 0; i < len(arr); i++ {
		print(arr[i], " ")
	}
	println()
}
