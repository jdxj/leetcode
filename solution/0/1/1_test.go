package s01

import (
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	nums := []int{3, 2, 1}
	res := marshal(nums)
	fmt.Printf("%s\n", res)
}

func TestThreeSum(t *testing.T) {
	nums := []int{0, 0}
	res := threeSum(nums)
	fmt.Printf("%v\n", res)
}
