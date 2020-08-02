// 234
package s23

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	mid := findMid(head)
	l := head
	r := reverseLink(mid.Next)

	for r != nil {
		if l.Val != r.Val {
			return false
		}
		l = l.Next
		r = r.Next
	}
	return true
}

func isPalindrome2(head *ListNode) bool {
	l, r := head, reverseLink(findRight(head, countLink(head)))
	lCur, rCur := l, r
	for rCur != nil {
		if lCur.Val != rCur.Val {
			return false
		}
		// len(lCur) >= len(rCur)
		lCur = lCur.Next
		rCur = rCur.Next
	}
	return true
}

func findRight(head *ListNode, count int) *ListNode {
	mid := (count + 1) / 2
	i, cur := 0, head
	for i = 1; i <= mid; i++ {
		cur = cur.Next
	}
	return cur
}

func countLink(head *ListNode) int {
	cur := head
	var count int
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

func reverseLink(head *ListNode) *ListNode {
	var newHead, float, cur *ListNode = nil, nil, head
	for cur != nil {
		// 弹出一个 node
		float = cur
		cur = cur.Next
		// 将该 node 指向头
		float.Next = newHead
		// 更新 newHead
		newHead = float
	}
	return newHead
}

func findMid(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
