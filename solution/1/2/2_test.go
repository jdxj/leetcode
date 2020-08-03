package s12

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(nums)
	fmt.Printf("%d\n", res)
}
