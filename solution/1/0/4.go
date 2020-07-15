// 104
package s10

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 0

	for len(queue) != 0 {
		depth++
		curLevelSize := len(queue)
		for i := 0; i < curLevelSize; i++ {
			// pop
			node := queue[0]
			queue = queue[1:]
			// push
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return depth
}
