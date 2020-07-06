package s739

import (
	"fmt"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	result := dailyTemperatures(T)
	fmt.Printf("%v\n", result)
}
