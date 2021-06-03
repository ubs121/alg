package main

// in-place 90 degree rotation
func rotateMatrix(matrix [][]int) {
	n := len(matrix)

	// 1. flip diagonally
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			matrix[i][j], matrix[n-j-1][n-i-1] = matrix[n-j-1][n-i-1], matrix[i][j]
		}
	}

	// 2. flip vertically
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j]
		}
	}
}
