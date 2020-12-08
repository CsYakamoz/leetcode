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
func getDecimalValue(head *ListNode) int {
	result, _ := helper(head)

	return result
}

func helper(curr *ListNode) (int, int) {
	if curr.Next == nil {
		return curr.Val * 1, 1
	}

	prevVal, prevMultiplier := helper(curr.Next)
	currMultiplier := prevMultiplier * 2

	return prevVal + curr.Val*currMultiplier, currMultiplier
}
