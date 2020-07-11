// 094
package s09

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//        1
//      /   \
//    2      3
//  / \    /  \
// 4  5   6   7
//
// 4 2 5 1 6 3 7

func inorderTraversal(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode

	for root != nil || len(stack) != 0 {
		for root != nil {
			// push
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, root.Val)
		root = root.Right
	}
	return result
}

func InorderTraversalRecursive(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	leftInorder := InorderTraversalRecursive(root.Left)
	rightInorder := InorderTraversalRecursive(root.Right)

	result := append([]int{}, leftInorder...)
	result = append(result, root.Val)
	result = append(result, rightInorder...)
	return result
}
