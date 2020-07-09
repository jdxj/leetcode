package s704

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	result := search(nums, 4)
	fmt.Printf("result: %d\n", result)

	nums = []int{-1, 0, 3, 5, 9, 12}
	result = search(nums, 12)
	fmt.Printf("result: %d\n", result)
}
