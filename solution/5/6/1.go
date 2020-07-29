// 561
package s56

func arrayPairSum(nums []int) int {
	offset := 10000
	amounts := make([]int, 20001)
	for _, v := range nums {
		amounts[v+offset]++
	}
	var d, sum int
	for i, amount := range amounts {
		sum += (amount + 1 - d) / 2 * (i - offset) // 2. 那么下一个元素要-1个
		d = (2 + amount - d) % 2                   // 1. 如果当前元素是奇数
	}
	return sum
}
