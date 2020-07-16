// 116
package s11

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 注意: 该题有 II, 相对于 I 更具有通用性

//                  1
//        2                  3
//   4         5          6      7
// 8   9    10  11     12  13  14  15

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	helper2(root, nil)
	return root
}

func helper2(root, next *Node) {
	if root == nil || root.Left == nil {
		return
	}

	root.Left.Next = root.Right
	if next != nil {
		root.Right.Next = next.Left
	}
	helper2(root.Left, root.Left.Next)
	helper2(root.Right, root.Right.Next)
}

// 以 II 的方法解决 I.
func connectI2(root *Node) *Node {
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
