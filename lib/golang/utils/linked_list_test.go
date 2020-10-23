package utils

import (
	"reflect"
	"testing"
)

var linkedListTests = []struct {
	input  []int
	output *ListNode
}{
	{
		input:  []int{},
		output: nil,
	},
	{
		input:  []int{1},
		output: &ListNode{Val: 1},
	},
	{
		input: []int{1, 2, 3, 4, 5},
		output: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{Val: 5},
					},
				},
			},
		},
	},
}

func TestArrayToLinkedList(t *testing.T) {
	for idx, test := range linkedListTests {
		if !reflect.DeepEqual(ArrayToLinkedList(test.input), test.output) {
			t.Fatalf("TestCase[%d]: linked lists are not same", idx)
		}
	}
}
