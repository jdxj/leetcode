package s49

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

func TestFindDiagonalOrder(t *testing.T) {
	matrix := [][]int{
		{3},
		{2},
	}
	result := findDiagonalOrder(matrix)
	t.Logf("%v\n", result)
}
