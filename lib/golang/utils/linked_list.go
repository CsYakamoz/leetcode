package utils

import (
	"github.com/CsYakamoz/leetcode/lib/golang/structure"
)

type ListNode = structure.ListNode

func ArrayToLinkedList(arr []int) *ListNode {
	length := len(arr)
	if length == 0 {
		return nil
	}

	head := &ListNode{}
	curr := head

	for _, v := range arr {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}

	return head.Next
}
