package alg

/******************************************************
 * QuickSort -  Массивын өгөгдсөн мужийг эрэмбэлэх    *
 *                                                    *
 * Параметрүүд                                        *
 *   low –- эрэмбэлэх мужийн доод хязгаар.            *
 *   high –- эрэмбэлэх мужийн дээд хязгаар.           *
 *                                                    *
 ******************************************************/
func QuickSort(array []int, low int, high int) {
	i := low
	j := high
	pivot := array[(low+high)/2]
	var temp int

	for i <= j {
		for array[i] < pivot {
			i++
		}
		for array[j] > pivot {
			j--
		}

		if i <= j {
			temp = array[i]
			array[i] = array[j]
			array[j] = temp
			i++
			j--
		}
	}

	if low < j {
		QuickSort(array, low, j)
	} /* зүүн хэсгийг эрэмбэлэх */

	if i < high {
		QuickSort(array, i, high)
	} /* баруун хэсгийг эрэмбэлэх */
}
