package s74

import (
	"fmt"
	"sort"
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{3, 4, 4, 5}
	idx := sort.Search(len(nums), func(i int) bool {
		return nums[i] > 4
	})
	fmt.Printf("idx: %d\n", idx)
}
