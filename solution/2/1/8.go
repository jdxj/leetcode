package s21

func findPeakElement(nums []int) int {
	lower, upper := 0, len(nums)-1
	for lower < upper {
		mid := (lower + upper) >> 1
		if nums[mid] > nums[mid+1] {
			upper = mid
		} else {
			lower = mid + 1
		}
	}
	return lower
}
