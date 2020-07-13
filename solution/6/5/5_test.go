package s65

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	result := findClosestElements(nums, 4, 3)
	fmt.Printf("%v\n", result)
}
