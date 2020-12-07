package golang

import (
	"github.com/CsYakamoz/leetcode/lib/golang/structure"
)

type ListNode = structure.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	head := &ListNode{Next: list1}

	first := head
	for i := 0; i < a; i++ {
		first = first.Next
	}

	last := first
	for i := 0; i < b-a+1; i++ {
		last = last.Next
	}

	first.Next = list2

	endOfList2 := list2
	for endOfList2.Next != nil {
		endOfList2 = endOfList2.Next
	}

	endOfList2.Next = last.Next

	return head.Next
}
