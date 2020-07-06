package s494

import "testing"

func TestFindTargetSumWays(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	correct := 5
	result := findTargetSumWays(nums, target)
	if result != correct {
		t.Fatalf("failed, correct: %d, your: %d\n", correct, result)
	}

}
