package s27

import (
	"github.com/jdxj/leetcode/util"
)

// 279
// 类似322题.

// 定义 DP[n]: 组成完全平方数 n 所需数字的最少个数
// DP[n] = min(DP[n-k]) + 1, k -> {1, 4, 9, 16, 25, ...}
//   即: DP[n] = min(DP[n-1], DP[n-4], ...) + 1
func numSquares(n int) int {
	// 初始化 dp 数组
	dp := make([]int, n+1)
	for i := range dp {
		// 这里不使用 "dp[i] = i" 的原因是:
		// 避免 n 刚好是所需最少个数, 所以加个1, 也可以直接使用 math.MaxXXX
		dp[i] = n + 1
	}
	dp[0] = 0

	// 初始化 k 可取的集合
	var kSet []int
	for i := 1; i*i <= n; i++ {
		kSet = append(kSet, i*i)
	}

	for i := 1; i < len(dp); i++ {
		// 不断的在 kSet 中寻找最少个数
		for j := 0; j < len(kSet); j++ {
			if i-kSet[j] >= 0 {
				dp[i] = util.Min2(dp[i], dp[i-kSet[j]]+1)
			}
		}
	}

	// 没找到
	if dp[n] > n {
		return -1
	}
	return dp[n]
}
