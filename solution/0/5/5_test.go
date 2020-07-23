package s05

import "testing"

func TestMerge(t *testing.T) {
	intervals := [][]int{
		{15, 18},
		{1, 3},
		{2, 6},
		{8, 10},
	}
	result := merge(intervals)
	t.Logf("%v\n", result)
}
