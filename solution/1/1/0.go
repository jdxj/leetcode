// 110
package s11

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 耗时较高

// 如果树是平衡的, 那么 "当前结点" 与 "左右子树" 也是平衡的.
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	cache = make(map[*TreeNode]int)

	left := helper3(root.Left)
	right := helper3(root.Right)
	if left < right {
		left, right = right, left
	}
	// 左右子树不平衡
	if left-right > 1 {
		return false
	}
	// 继续递归检测
	return isBalanced(root.Left) && isBalanced(root.Right)
}

// cache 用于记录 root 的高度.
var cache map[*TreeNode]int

// 返回一个结点的高度.
func helper3(root *TreeNode) int {
	if root == nil {
		// 这里定义叶子结点的高度为0,
		// 所以 nil 结点的高度为-1.
		return -1
	}
	// 从缓存中找到 root 的高度
	if height, ok := cache[root]; ok {
		return height
	}
	// 没找到就计算左右子树的高度
	left := helper3(root.Left)
	right := helper3(root.Right)
	// 取最大的一个
	if left < right {
		left = right
	}
	cache[root] = left + 1
	return left + 1
}
