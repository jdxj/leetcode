// 15
package s01

import (
	"sort"
	"strconv"
	"strings"
)

// 时间/空间复杂度极高
func threeSum(nums []int) [][]int {
	numsLen := len(nums)
	if numsLen < 3 {
		return nil
	}
	cache := make(map[int]map[int]struct{})
	for i, v := range nums {
		if _, ok := cache[v]; !ok {
			cache[v] = make(map[int]struct{})
		}
		cache[v][i] = struct{}{}
	}

	result := make(map[string][]int)
	for i := 0; i < numsLen-1; i++ {
		for j := i + 1; j < numsLen; j++ {
			if ids, ok := cache[-(nums[i] + nums[j])]; !ok {
				continue
			} else {
				// 判断是否重复使用
				_, ok1 := ids[i]
				_, ok2 := ids[j]
				if ok1 && len(ids) == 1 || ok2 && len(ids) == 1 ||
					nums[i] == nums[j] && len(ids) == 2 {
					continue
				}
			}
			tmp := []int{nums[i], nums[j], -(nums[i] + nums[j])}
			result[marshal(tmp)] = tmp
		}
	}
	res := make([][]int, 0, len(result))
	for _, v := range result {
		res = append(res, v)
	}
	return res
}

func marshal(st []int) string {
	sort.Ints(st)
	stStr := make([]string, len(st))
	for i, e := range st {
		stStr[i] = strconv.Itoa(e)
	}
	return strings.Join(stStr, ",")
}
