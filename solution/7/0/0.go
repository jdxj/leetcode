// 700
package s704

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		// 应该向右寻找
		if root.Val < val {
			root = root.Right
		} else if root.Val > val { // 向左查找
			root = root.Left
		} else { // 找到了
			return root
		}
	}
	return nil
}
