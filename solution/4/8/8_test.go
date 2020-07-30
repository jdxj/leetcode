package s48

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	nums := []int{1, 0, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 1, 1}
	t.Logf("count: %d\n", len(nums))
	res := findMaxConsecutiveOnes(nums)
	t.Logf("len: %d\n", res)
}
