package s542

import (
	"testing"

	"github.com/jdxj/leetcode/util"
)

// 输入:
//   0 0 0
//   0 1 0
//   1 1 1
// 输出:
//   0 0 0
//   0 1 0
//   1 2 1
func TestUpdateMatrix(t *testing.T) {
	matrix := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	updateMatrix(matrix)

	util.PrintTwoDimensionalArray(matrix)
}
