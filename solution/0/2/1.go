// 21
package s02

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	l1cur, l2cur := l1, l2
	for l1cur != nil && l2cur != nil {
		if l1cur.Val > l2cur.Val {
			l1cur, l2cur = l2cur, l1cur
		}
		cur.Next = l1cur
		cur = cur.Next
		l1cur = l1cur.Next
	}
	if l1cur == nil {
		l1cur, l2cur = l2cur, l1cur
	}
	cur.Next = l1cur
	return head.Next
}
