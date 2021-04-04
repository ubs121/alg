package alg

// https://en.wikipedia.org/wiki/Partial_sorting
func Partial_sort_left(a []int, i, j, k int) {
	if i < j {
		p := (i + j) / 2 //i + rand.Intn(j-i+1)
		p = partition(a, i, j, p)

		Partial_sort_left(a, i, p-1, k)

		if p < k {
			Partial_sort_left(a, p+1, j, k)
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
