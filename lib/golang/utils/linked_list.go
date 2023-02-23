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

func LinkedListToArray(head *ListNode) []int {
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}

func CopyLinkList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := &ListNode{}
	ptr := newHead
	for ; head != nil; head = head.Next {
		ptr.Next = &ListNode{Val: head.Val}
		ptr = ptr.Next
	}

	return newHead.Next
}
