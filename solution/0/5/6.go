// 056
package s05

import (
	"sort"
)

type Data [][]int

func (d Data) Len() int {
	return len(d)
}

func (d Data) Less(i, j int) bool {
	return d[i][0] < d[j][0]
}

func (d Data) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func merge(intervals [][]int) [][]int {
	// 只有一个区间或以下
	if len(intervals) <= 1 {
		return intervals
	}

	// 对各区间按照左边界排序.
	// 排序的目的是: 固定每个区间的左边界升序后就可以依次对结果集的末尾区间进行合并.
	// 也就是说, 对于新来的区间, 它只能是结果集中最后一个区间的向后交集或者向后无交集,
	// 确保新来的区间不会跑到前面去.
	data := Data(intervals)
	sort.Sort(data)
	// 添加第一个区间进行初始化,
	// 对于排好序的 intervals 来说, 第一个区间的左边界是最小的值.
	result := [][]int{
		{intervals[0][0], intervals[0][1]}, // 深拷贝
	}

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		// 新区间左边界小于等于 last 的右边界, 说明有重合.
		if intervals[i][0] <= last[1] {
			// 检查 last 右边界是否扩张
			if last[1] < intervals[i][1] {
				last[1] = intervals[i][1]
			}
		} else { // 否则表示没有重合, 需要添加新区间
			result = append(result, []int{intervals[i][0], intervals[i][1]})
		}
	}
	return result
}
