// 154
package s15

func findMinII(nums []int) int {
	lower, upper := 0, len(nums)-1
	if upper < 0 {
		return -1
	}

	for lower < upper {
		mid := lower + (upper-lower)>>1 // lower/2 + upper/2
		if nums[mid] > nums[upper] {    // 最小值在右边
			lower = mid + 1
		} else if nums[mid] == nums[upper] {
			upper--
		} else { // 最小值在左边
			upper = mid
		}
	}
	return nums[lower]
}
