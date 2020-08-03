// 53
package s05

// DP[i] 存储的是: 已 nums[i] 结尾的最大子序和. 注意: 不是 "nums[i] 内的最大子序和".
// DP[i] = max(DP[i-1] + nums[i], nums[i])
// 迷惑点在于: 在整个 nums 中, 最大连续子序和的位置可能是 nums 的中间位置,
// 但是只是记录 max 就可以了.
// 更简单的理解这道题: 如果要寻找最大子序和, 那么这个子序列中包含的正数要尽可能多,
// 如果在 "前一个正数" 加到 "后一个正数" 的过程中, 负数过多, 那么 "后一个正数"
// 可以舍弃它之前的序列, 自己作为新的 "前一个正数" 来开始新的寻找.
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		// 如果 "之前的最大子序和+当前值" 还不如当前值大, 那么重新开一个子序列好了.
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
