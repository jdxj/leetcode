// 145
package s14

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append([]int{node.Val}, result...)
		// push
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return result
}
