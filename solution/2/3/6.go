// 236
package s23

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil ||
		root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// p, q 在当前结点的左右两侧时, 说明当前结点就是最近
	if left != nil && right != nil {
		return root
	}
	// p, q 同时在一边的情况需要继续递归
	if left != nil {
		return lowestCommonAncestor(left, p, q)
	}
	if right != nil {
		return lowestCommonAncestor(right, p, q)
	}
	return nil
}
