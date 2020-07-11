// 144
package s14

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var result []int
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}

		result = append(result, node.Val)
		// push
		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
	}
	return result
}
