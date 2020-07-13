// 658
package s65

import (
	"fmt"
	"sort"
)

// 放弃
func findClosestElements(arr []int, k int, x int) []int {
	var result []int
	diff := make([]int, len(arr))
	for i := range arr {
		diff[i] = x - arr[i]
		if diff[i] < 0 {
			diff[i] = -diff[i]
		}
	}
	fmt.Printf("diff: %v\n", diff)

	diff2 := make([]int, len(diff))
	copy(diff2, diff)
	sort.Ints(diff2)
	fmt.Printf("diff2: %v\n", diff2)

	kMin := diff2[:k]

	fmt.Printf("kMin: %v\n", kMin)
	for _, v := range kMin {
		for idx, d := range diff {
			if v == d {
				result = append(result, arr[idx])
			}
		}
	}
	record := make(map[int]struct{})
	for _, v := range result {
		record[v] = struct{}{}
	}
	result = result[:0]
	for k := range record {
		result = append(result, k)
	}
	sort.Ints(result)
	return result[:k]
}
