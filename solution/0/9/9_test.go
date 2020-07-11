package s09

import (
	"fmt"
	"testing"

	"github.com/jdxj/leetcode/util"
)

func TestInorderTraversal(t *testing.T) {
	result := inorderTraversal(getBinaryTree())
	fmt.Printf("%v\n", result)
}

func getBinaryTree() *TreeNode {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 7}

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7
	return node1
}

func TestInorderTraversalRecursive(t *testing.T) {
	result := InorderTraversalRecursive(getBinaryTree())
	if !util.IsEqualSliceInt(result, []int{4, 2, 5, 1, 6, 3, 7}) {
		t.Fatalf("failed\n")
	}
}
