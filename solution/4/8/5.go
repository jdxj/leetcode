// 485
package s48

// 内存占用大
func findMaxConsecutiveOnes(nums []int) int {
	numsLen := len(nums)
	var maxCount, count int
	for i := 0; i < numsLen; i++ {
		if nums[i] == 1 {
			count++
			continue
		}
		// nums[i] == 0
		maxCount = max(maxCount, count)
		count = 0
	}
	return max(maxCount, count)
}

func max(p1, p2 int) int {
	if p1 < p2 {
		p1 = p2
	}
	return p1
}
