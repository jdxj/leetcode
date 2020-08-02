// 019
package s01

type ListNode struct {
	Val  int
	Next *ListNode
}

var cache = make([]*ListNode, 0, 100)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cache = cache[:0]
	var idx int
	var cur *ListNode
	for idx, cur = 0, head; cur != nil; idx, cur = idx+1, cur.Next {
		cache = append(cache, cur)
	}
	targetPre := idx - n - 1
	if targetPre < 0 {
		return head.Next
	}
	cache[targetPre].Next = cache[targetPre].Next.Next
	return head
}
