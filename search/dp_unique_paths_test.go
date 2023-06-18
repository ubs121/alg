package search

// Unique Paths: A robot is located at the top-left corner of a m x n grid. 
// It can only move either down or right. 
// Calculate the number of unique paths to reach the bottom-right corner.
func uniquePaths(m int, n int) int {
    dp := make([][]int, m)

    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }

    // Set the number of ways to reach the cells in the first row and first column to 1
    for i := 0; i < m; i++ {
        dp[i][0] = 1
    }
    for j := 0; j < n; j++ {
        dp[0][j] = 1
    }

    // Calculate the number of ways for each cell based on the previous cells
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }

    return dp[m-1][n-1]
}
