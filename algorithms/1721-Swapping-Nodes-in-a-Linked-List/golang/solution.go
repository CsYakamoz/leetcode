package golang

import (
	"github.com/CsYakamoz/leetcode/lib/golang/structure"
)

type ListNode = structure.ListNode

func swapNodes(head *ListNode, k int) *ListNode {
	dump := &ListNode{Val: 0, Next: head}
	p1 := dump
	trulyK := k - 1

	for trulyK != 0 {
		p1 = p1.Next
	}

	first := dump
	for p1 != nil {
		first = first.Next
		p1 = p1.Next
	}

	return nil
}

