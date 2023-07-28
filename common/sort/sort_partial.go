package sort

// https://en.wikipedia.org/wiki/Partial_sorting
func PartialSortLeft(arr []int, i, j, k int) {
	if i < j {
		p := (i + j) / 2 //i + rand.Intn(j-i+1)
		p = partition(arr, i, j, p)

		PartialSortLeft(arr, i, p-1, k)

		if p < k {
			PartialSortLeft(arr, p+1, j, k)
		}
	}
}

// https://en.wikipedia.org/wiki/Partial_sorting
func PartialSortRight(arr []int, i, j, k int) {
	if i < j {
		p := (i + j) / 2
		p = partition(arr, i, j, p)

		PartialSortRight(arr, p+1, j, k)

		if p > k-1 {
			PartialSortRight(arr, i, p-1, k)
		}
	}
}

// https://en.wikipedia.org/wiki/Quickselect
func partition(arr []int, l, r, p int) int {
	v := arr[p]

	// Move pivot to end
	arr[p], arr[r] = arr[r], arr[p]

	i := l
	for j := l; j < r; j++ {
		if v > arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// Move pivot to its final place
	arr[r], arr[i] = arr[i], arr[r]

	return i
}
