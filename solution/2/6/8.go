// 268
package s26

func missingNumber(nums []int) int {
	miss := len(nums)
	for i, v := range nums {
		miss = miss ^ i ^ v
	}
	return miss
}
