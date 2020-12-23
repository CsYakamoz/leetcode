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
func removeZeroSumSublists(head *ListNode) *ListNode {
	dict := make(map[int]int)
	dict[0] = 0

	prefixSum := make([]*ListNode, 0)
	prefixSum = append(
		prefixSum,
		&ListNode{Val: 0, Next: head},
	)

	for i, sum := 1, 0; head != nil; {
		sum += head.Val

		idx, exist := dict[sum]
		if exist {
			prefixSum[idx].Next = head.Next

			for j, key := idx+1, sum; j < len(prefixSum) && j < i; j++ {
				key += prefixSum[j].Val
				delete(dict, key)
			}

			i = idx
			prefixSum = prefixSum[0 : idx+1]
		} else {
			dict[sum] = i
			prefixSum = append(
				prefixSum,
				head,
			)
		}

		head = head.Next
		i++
	}

	return prefixSum[0].Next
}
