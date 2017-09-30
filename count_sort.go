package main

/**
* 	A -  array of integers
* 	l - left boundary
* 	r - right boundary
 */
func countingSort(A []int, l, r int) {
	n := r - l + 1
	B := make([]int, n)

	C := make(map[int]int)

	// [l,r] интервалд i хэдэн удаа байгааг тоолох
	for i := 0; i < n; i++ {
		C[A[l+i]] += 1
	}
	fmt.Println(C)

	// i-1 тоо хэдэн удаа байгааг тоолох
	for i := 1; i < len(C); i++ {
		C[i] = C[i] + C[i-1] // FIXME: энд алдаа байна !!!
	}

	fmt.Println("DONE", len(C))

	// тоонуудыг байрлуулах
	for i := n - 1; i >= 0; i-- {
		B[C[A[l+i]]-1] = A[l+i]
		C[A[l+i]] = C[A[l+i]] - 1
	}

	// B -> A
	for i := 0; i < n; i++ {
		A[l+i] = B[i]
	}

	// DEBUG: fmt.Printf("(%d, %d): %v\n", l, r, A)
}
