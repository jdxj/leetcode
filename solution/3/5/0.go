// 350
package s35

func intersect(nums1 []int, nums2 []int) []int {
	var result []int

	// <某个数字>:<出现次数>
	filter := make(map[int][]int)
	for _, v := range nums1 {
		if _, ok := filter[v]; !ok {
			filter[v] = make([]int, 2)
		}
		filter[v][0]++
	}

	for _, v := range nums2 {
		if stat, ok := filter[v]; ok {
			stat[1]++
			if stat[0] >= stat[1] {
				result = append(result, v)
			}
		}
	}
	return result
}
