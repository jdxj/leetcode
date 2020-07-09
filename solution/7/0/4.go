package s704

import "sort"

func search(nums []int, target int) int {
	idx := sort.SearchInts(nums, target)
	if idx >= len(nums) || nums[idx] != target {
		return -1
	}
	return idx
}
