// 033
package s03

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	lower, upper := 0, len(nums)-1
	for lower <= upper {
		mid := (lower + upper) >> 1
		if nums[mid] == target { // 立马就找到了
			return mid
		}

		// 假设左半部是有序的
		if nums[lower] <= nums[mid] {
			// 假设 target 在有序部分
			if nums[lower] <= target && target < nums[mid] {
				upper = mid - 1
			} else { // 未在有序部分时, 继续切分无序部分
				lower = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[upper] {
				lower = mid + 1
			} else {
				upper = mid - 1
			}
		}
	}
	return -1
}
