package s23

import (
	"fmt"
	"testing"
)

func TestReverseLink(t *testing.T) {
	l := getLink()
	newL := reverseLink(l)

	for cur := newL; cur != nil; cur = cur.Next {
		fmt.Printf("%d\n", cur.Val)
	}
}

func getLink() *ListNode {
	l1 := &ListNode{
		Val: 1,
	}
	l2 := &ListNode{
		Val: 2,
	}
	l3 := &ListNode{
		Val: 3,
	}
	l4 := &ListNode{
		Val: 4,
	}
	l5 := &ListNode{
		Val: 5,
	}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	return l1
}

func TestCountLink(t *testing.T) {
	l := getLink()
	res := countLink(l)
	if res != 5 {
		t.Fatalf("failed, corrent: %d, your: %d", 5, res)
	}
}

func TestFindRight(t *testing.T) {
	l := getLink()
	r := findRight(l, 5)
	if r.Val != 4 {
		t.Fatalf("failed, corrent: %d, your: %d\n", 4, r.Val)
	}
}
