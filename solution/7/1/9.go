// 719
package s71

import "sort"

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	lower := 0                           // 最小数对差最小就是两个数想等, 为0
	upper := nums[len(nums)-1] - nums[0] // 最大数对差 = 最大值-最小值

	for lower < upper {
		mid := lower + (upper-lower)>>1 // 处于中间的一个数对差
		// 看看数对差 <=mid 的有多少个
		var count int
		for left, right := 0, 0; right < len(nums); right++ { // 优先移动右边界
			// 如果两个边界的数对差大于 mid, 则需要移动左边界才能使数据差再次 <=mid.
			for nums[right]-nums[left] > mid {
				left++
			}
			count += right - left // 索引相减求的是数对差 <=mid 的个数
		}
		// 如果数对差 <=mid 的个数 >k, 则说明数对差 <=mid 的个数过多,
		// 需要缩小右边界, 来排除数对差 >mid 的数对.
		if count >= k {
			upper = mid
		} else { // 反之, 数对差 <=mid 的个数不足, 需要扩大范围.
			lower = mid + 1
		}
	}
	return lower
}
