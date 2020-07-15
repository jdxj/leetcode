package s11

import "testing"

func TestHasPathSum(t *testing.T) {
	if hasPathSum(getTree(), 2) {

	}
}

//     1
//   /
//  2
func getTree() *TreeNode {
	node1 := &TreeNode{
		Val: 1,
	}
	node2 := &TreeNode{
		Val: 1,
	}
	node1.Left = node2
}
