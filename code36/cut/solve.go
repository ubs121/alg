// https://www.hackerrank.com/contests/w36/challenges/cut-a-strip
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func kadane2d(arr [][]int, rows, cols int) (maxSum int64, maxRect *Rect) {
	maxSum = math.MinInt64

	rects := make([]*Rect, rows/2) // random size
	n := 0

	// O(n^3) ?
	// TODO: choose (l,r) randomly ???
	for l := 0; l < cols; l++ { // left column
		// build temp 1d array for Kadane's algorithm
		temp := make([]int64, rows)

		for r := l; r < cols; r++ { // right column
			// TODO: put strip (cut)

			// special case 1: all negative => same as other cases
			// special case 2: all positive => cut minimum cell if full size, else outside cut/no cut

			/* start of Kadane's algorithm */
			temp[0] = int64(arr[0][r])
			sum := temp[0]
			tempMaxSum := sum
			tempRect := &Rect{left: l, right: r, top: 0, bottom: 0}
			start := 0

			// kadane algorithm
			for i := 1; i < rows; i++ {
				// accumulate arr[l:r] into 'temp'
				temp[i] += int64(arr[i][r])

				sum += temp[i]
				if sum < temp[i] {
					sum = temp[i]
					start = i
				}

				if tempMaxSum < sum {
					tempMaxSum = sum
					tempRect.top = start
					tempRect.bottom = i
					tempRect.sum = tempMaxSum
				}

			}

			if maxSum < tempMaxSum {
				maxSum = tempMaxSum
				maxRect = tempRect

				// put in the queue
				rects[n] = maxRect
				n++

				// resize array if needed
				if n == len(rects) {
					tmp := make([]*Rect, n)
					rects = append(rects, tmp...)
				}
			}

		}
	}

	// TODO: remove contained rects
	rects = rects[:n]
	sort.Sort(BySum(rects))

	// TODO: merge rects
	fmt.Printf("%v\n", rects)

	return
}

func main() {
	var n, m, k int
	fmt.Scanf("%d %d %d", &n, &m, &k)

	arr := make([][]int, n)

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		buf, _ := reader.ReadBytes('\n')
		arr[i] = readInts(buf, m)
	}

	max, rect := kadane2d(arr, n, m)

	fmt.Printf("maxSum = %d %v\n", max, rect)

}
