// 98
package s09

import "math"

func isValidBST(root *TreeNode) bool {
	var stack []*TreeNode
	preVal := math.MinInt64
	for root != nil || len(stack) != 0 {
		for root != nil {
			// push
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= preVal {
			return false
		}
		preVal = root.Val
		root = root.Right
	}
	return true
}
