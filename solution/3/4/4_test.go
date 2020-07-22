package s34

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	nums := []int{1, 2}
	result := topKFrequent(nums, 2)
	fmt.Printf("%v\n", result)
}
