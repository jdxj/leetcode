// 287
package s28

// 规律: 只有一个元素是重复的
func findDuplicate(nums []int) int {
	lower, upper := 1, len(nums)-1
	for lower < upper {
		mid := lower + (upper-lower)>>1
		// 统计 <= mid 的个数
		var cnt int
		for _, v := range nums {
			if v <= mid {
				cnt++
			}
		}
		// 有重复, 且重复的元素在 lower~mid 中.
		// 比如: 1, 2, 2, 3, 4, 5, 6, 7, 8, 9, mid = 5, 统计 <=5 的有6个
		if cnt > mid {
			upper = mid
		} else { // 在 mid~upper 中
			// 比如: 1, 2, 3, 4, 5, 6, 6, 7, 8, 9, mid = 5, 统计 <=5 的有5个
			lower = mid + 1
		}
	}
	return lower
}
