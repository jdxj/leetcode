// 14
package s01

import "sort"

func longestCommonPrefix(strs []string) string {
	amount := len(strs)
	if amount == 0 {
		return ""
	}

	// 排序的目的是找出最小的字符串.
	// 最小的字符串应该是: 1.长度小 2.ascii 小
	sort.Strings(strs)
	result := make([]uint8, 0, len(strs[0]))
	for i := range strs[0] {
		if strs[0][i] == strs[amount-1][i] {
			result = append(result, strs[0][i])
		} else {
			// 只要不相同说明已找完公共前缀
			break
		}
	}
	return string(result)
}
