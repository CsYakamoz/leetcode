package golang

import (
	"github.com/CsYakamoz/leetcode/lib/golang/structure"
)

type ListNode = structure.ListNode

func removeNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	stack := []*ListNode{}
	for head != nil {
		for len(stack) != 0 && head.Val > stack[len(stack)-1].Val {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, head)
		head = head.Next
	}

	newHead := stack[0]
	prev := newHead
	for i := 1; i < len(stack); i++ {
		prev.Next = stack[i]
		prev = stack[i]
	}

	return newHead
}
