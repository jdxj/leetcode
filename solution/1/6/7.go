// 167
package s16

import "sort"

func twoSum(numbers []int, target int) []int {
	var result []int
	for i, part1 := range numbers {
		part2 := target - part1
		idx := sort.SearchInts(numbers, part2)
		// 没找到 part2 或者使用了 i 本身都算没找到
		if idx == len(numbers) || numbers[idx] != part2 || idx == i {
			continue
		}
		result = append(result, i+1, idx+1)
		break
	}
	sort.Ints(result)
	return result
}
