// 198
package s19

// 这道题干想就是: 要么都选奇数的房间, 要么都选偶数的房间, 选一个比较大的 (不完全正确, 可能在中途空了两个没选, 那么奇就调换了).
// DP:
//   定义 DP[i] 为前 i 个房间中 (包括 i) 可偷取的最大值.
//   dp[0] = nums[0], 只有一个房间那么就是该房间;
//   dp[1] = max(nums[0], nums[1]), 有两个房间时选择其中一个较大的;
//   dp[i]: 对于 dp[i] 来说, 有两种情况:
//     1. 选择 nums[i], 那么 nums[i-1] 就不能选, dp[i] = dp[i-2] + nums[i]
//     2. 不选择 nums[i], 那么 dp[i] = dp[i-1]
//     3. 在情况1, 2中选择一个较大的值
func rob(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	} else if numsLen == 1 {
		return nums[0]
	}

	dp := make([]int, 2)
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < numsLen; i++ {
		dp[i%2] = max(dp[(i-2)%2]+nums[i], dp[(i-1)%2])
	}
	return dp[(numsLen-1)%2]
}

func max(p1, p2 int) int {
	if p1 < p2 {
		p1 = p2
	}
	return p1
}
