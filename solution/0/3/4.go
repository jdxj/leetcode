// 034
package s03

func searchRange(nums []int, target int) []int {
	reslut := []int{-1, -1}

	leftIdx := extremeInsertionIndex(nums, target, true)
	if leftIdx == len(nums) || nums[leftIdx] != target {
		return reslut
	}

	reslut[0] = leftIdx
	reslut[1] = extremeInsertionIndex(nums, target, false) - 1
	return reslut
}

// extremeInsertionIndex 用于在 nums 中查找 target 所在位置,
// findLeft 用于控制查找左端或者右端.
func extremeInsertionIndex(nums []int, target int, findLeft bool) int {
	lower, upper := 0, len(nums)
	for lower < upper {
		mid := lower + (upper-lower)>>1
		if nums[mid] > target || (findLeft && nums[mid] == target) {
			upper = mid
		} else {
			lower = mid + 1
		}
	}
	return lower
}
