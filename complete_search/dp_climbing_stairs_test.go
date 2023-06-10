package search

// Climbing Stairs: You are climbing a staircase with n steps. Each time you can either climb 1 or 2 steps. 
// Determine the number of distinct ways to climb to the top.
func climbStairs(n int) int {
    if n <= 2 {
        return n
    }

    dp := make([]int, n+1) // number of distinct ways to climb
    dp[1] = 1
    dp[2] = 2

    for i := 3; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }

    return dp[n]
}
