package numbers

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	m := len(matrix[0])

	d := 0

	var elems []int
	for d < n/2 && d < m/2 {
		i := d
		j := d

		// right
		for j < m-d-1 {
			elems = append(elems, matrix[i][j])
			j++
		}

		// down
		for i < n-d-1 {
			elems = append(elems, matrix[i][j])
			i++
		}

		// left
		for j > d {
			elems = append(elems, matrix[i][j])
			j--
		}

		// up
		for i > d {
			elems = append(elems, matrix[i][j])
			i--
		}

		d++
	}

	// left-over in the middle
	if len(elems) < n*m {
		if n < m && n%2 == 1 {
			for j := d; j <= m-d-1; j++ {
				elems = append(elems, matrix[d][j])
			}
		} else {
			for i := d; i <= n-d-1; i++ {
				elems = append(elems, matrix[i][d])
			}
		}
	}

	return elems
}
