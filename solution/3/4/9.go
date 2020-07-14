// 349
package s34

func intersection(nums1 []int, nums2 []int) []int {
	var result []int

	// <某个数字>:<是否出现过>
	filter := make(map[int]bool)
	for _, v := range nums1 {
		filter[v] = false // 初始化时都是 false
	}

	for _, v := range nums2 {
		if seen, ok := filter[v]; ok && !seen {
			filter[v] = true
			result = append(result, v)
		}
	}
	return result
}
