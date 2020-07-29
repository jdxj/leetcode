// 27
package s02

func removeElement(nums []int, val int) int {
	var i int
	for j := len(nums); i < j; {
		if nums[i] == val {
			// 与最后标记位交换
			j--
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}
	return i
}
