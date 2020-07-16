// 117
package s11

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connectII(root *Node) *Node {
	for cur := root; cur != nil; {
		// 用于构建下一层的链表
		head := &Node{} // 头结点
		tail := head
		// 构建过程
		for ; cur != nil; cur = cur.Next {
			if cur.Left != nil {
				tail.Next = cur.Left
				tail = tail.Next
			}
			if cur.Right != nil {
				tail.Next = cur.Right
				tail = tail.Next
			}
		}
		// 进入下一层
		cur = head.Next
	}
	return root
}
