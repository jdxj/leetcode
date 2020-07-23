// 193
package s19

func pivotIndex(nums []int) int {
	var sum int // 求和
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	// 每次假设当前 nums[i] 是中心索引
	for i, left := 0, 0; i < len(nums); i++ {
		// left == right
		if left == sum-nums[i]-left {
			return i
		}
		left += nums[i]
	}
	return -1
}
