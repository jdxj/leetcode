// 347
package s34

func topKFrequent(nums []int, k int) []int {
	// ele: amount, 实际上没有 len(nums) 那么多
	numsLen := len(nums)
	numsMap := make(map[int]int, numsLen)
	// 0: ele; 1: 次数
	// 用于记录最大 ele 出现的次数
	max := [2]int{}
	for i := 0; i < numsLen; i++ {
		numsMap[nums[i]]++
		if numsMap[nums[i]] > max[1] {
			max[0] = nums[i]
			max[1] = numsMap[nums[i]]
		}
	}

	result := make([][]int, max[1]+1)
	for ele, amount := range numsMap {
		result[amount] = append(result[amount], ele)
	}
	// 取出 k 个
	var preK []int
	for i := len(result) - 1; i >= 0; i-- {
		for j := 0; j < len(result[i]); j++ {
			preK = append(preK, result[i][j])
			if len(preK) == k {
				goto end
			}
		}
	}
end:
	return preK
}
