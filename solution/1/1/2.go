// 112
package s11

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return helper(root, sum, 0)
}

func helper(root *TreeNode, sum, part int) bool {
	if root.Left == nil && root.Right == nil {
		if sum == part+root.Val {
			return true
		}
		return false
	}

	var left, right bool
	if root.Left != nil {
		left = helper(root.Left, sum, part+root.Val)
	}
	if root.Right != nil {
		right = helper(root.Right, sum, part+root.Val)
	}
	return left || right
}
