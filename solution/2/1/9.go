// 219
package s21

func containsNearbyDuplicate(nums []int, k int) bool {
	numsMap := make(map[int]struct{})
	for i, v := range nums {
		if _, ok := numsMap[v]; ok {
			return true
		}
		numsMap[v] = struct{}{}
		if len(numsMap) > k { // 类似滑动窗口
			delete(numsMap, nums[i-k])
		}
	}
	return false
}
