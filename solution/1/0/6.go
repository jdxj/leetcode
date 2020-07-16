// 106
package s10

//       1
//   2       3
// 4   5   6   7
//
// inorder:    4 2 5 1 6 3 7
// postorder:  4 5 2 6 7 3 1

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == len(postorder) &&
		len(inorder) == 0 {
		return nil
	}

	// "后序" 中最后一个元素是 root
	postRoot := len(postorder) - 1
	// 构建 root
	root := &TreeNode{
		Val: postorder[postRoot],
	}
	// 在 "中序" 中寻找 root 的索引
	var inRoot int
	for i, v := range inorder {
		if v == postorder[postRoot] {
			inRoot = i
			break
		}
	}
	// 计算左子树的元素的个数
	leftAmount := len(inorder[:inRoot])

	// 构建左子树, 传递 root 的左子树的 "中序" 与 "后序"
	root.Left = buildTree(inorder[:inRoot], postorder[:leftAmount])
	// 构建右子树, 传递 root 的右子树的 "中序" 与 "后序"
	root.Right = buildTree(inorder[inRoot+1:], postorder[leftAmount:postRoot])
	return root
}
