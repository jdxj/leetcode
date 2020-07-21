// 599
package s59

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	var result []string
	idxSum := math.MaxInt64

	list2Map := make(map[string]int)
	for i, v := range list2 {
		list2Map[v] = i
	}

	for idx1, v := range list1 {
		if idx2, ok := list2Map[v]; ok {
			sum := idx1 + idx2
			if sum < idxSum {
				// 发现更小的 "和" 后清空原来的结果, 并添加
				result = append([]string{}, v)
				idxSum = sum
			} else if sum == idxSum {
				result = append(result, v)
			}
		}
	}
	return result
}
