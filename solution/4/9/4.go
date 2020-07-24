package s49

//
func findTargetSumWays(nums []int, S int) int {
	return 0
}

// 超出时间限制, 139个用例都通过
//func findTargetSumWays(nums []int, S int) int {
//	return helper(nums, -1, S)
//}
//
//func helper(nums []int, index, S int) int {
//	// 结算
//	if index == len(nums)-1 {
//		sum := 0
//		for _, v := range nums {
//			sum += v
//		}
//		if sum == S {
//			return 1
//		}
//		return 0
//	}
//
//	numPos := helper(nums, index+1, S) // 整数情况
//	nums[index+1] = -nums[index+1]
//	numNeg := helper(nums, index+1, S) // 负数情况
//	nums[index+1] = -nums[index+1]     // 修正
//	return numPos + numNeg
//}
