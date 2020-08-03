// 70
package s07

// 因为到达最顶层之前, 那么迈1步, 要么迈2步, 所以 DP[n] = DP[n-1] + DP[n-2].
// 这题时倒着想的, 正着想要寻找所有情况.
func climbStairs(n int) int {
	dp := []int{1, 1}
	for i := 2; i <= n; i++ {
		//dp[i] = dp[i-1] + dp[i-2]
		dp[i%2] = dp[(i-1)%2] + dp[(i-2)%2]
	}
	return dp[n%2]
}
