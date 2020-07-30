// 209
package s20

func minSubArrayLen(s int, nums []int) int {
	numsLen := len(nums)
	minCount := numsLen + 1
	var start, end, sum int
	for end < numsLen {
		sum += nums[end]
		for sum >= s {
			minCount = min(minCount, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if minCount == numsLen+1 {
		return 0
	}
	return minCount
}

func min(p1, p2 int) int {
	if p1 > p2 {
		p1 = p2
	}
	return p1
}
