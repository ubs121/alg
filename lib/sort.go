package alg

// QuickSort - range [low..high], O(nlogn), worst O(n^2)
func QuickSort(arr []int, low int, high int) {
	i := low
	j := high
	pivot := arr[(low+high)/2]
	var temp int

	for i <= j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}

		if i <= j {
			temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
			i++
			j--
		}
	}

	if low < j {
		QuickSort(arr, low, j)
	} /* зүүн хэсгийг эрэмбэлэх */

	if i < high {
		QuickSort(arr, i, high)
	} /* баруун хэсгийг эрэмбэлэх */
}

// Counting sort is not a comparison sort algorithm, O(n)
// assumes each element is in range [1..maxVal)
func CountingSort(arr []int, maxVal int) {
	n := len(arr) // number of elements

	// count each elements
	countArr := make([]int, maxVal)
	for i := 0; i < n; i++ {
		countArr[arr[i]]++
	}

	// do prefix sum for each element
	for i := 1; i < maxVal; i++ {
		countArr[i] += countArr[i-1]
	}

	// place each elements in order
	sortedArr := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		sortedArr[countArr[arr[i]]-1] = arr[i]
		countArr[arr[i]]--
	}

	// copy back to 'arr'
	copy(arr, sortedArr)
}
