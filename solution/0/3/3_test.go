package s03

import (
	"fmt"
	"sort"
	"testing"
)

func TestExtremeInsertionIndex(t *testing.T) {
	//nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	nums := []int{0, 1, 2, 3, 3, 3, 3, 3, 3, 3, 4, 5, 6, 7, 8, 9}
	idx := extremeInsertionIndex(nums, 3, true)
	fmt.Printf("left: %d\n", idx)
	//idx = extremeInsertionIndex(nums, 3, false)
	//fmt.Printf("right: %d\n", idx)

	idx = sort.Search(len(nums), func(i int) bool {
		return nums[i] >= 3
	})
	fmt.Printf("idx1: %d\n", idx)
	idx = sort.Search(len(nums), func(i int) bool {
		return nums[i] <= 3
	})
	fmt.Printf("idx2: %d\n", idx)
}

func TestNext(t *testing.T) {
	res := ""
	for i := 0; i < 30; i++ {
		res = next(i, res)
		fmt.Printf("i: %2d, len: %4d, res: %s\n", i+1, len(res), res)
	}
}
