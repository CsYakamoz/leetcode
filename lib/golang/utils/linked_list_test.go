package utils

import (
	"reflect"
	"testing"
)

var linkedListTests = []struct {
	arr  []int
	head *ListNode
}{
	{
		arr:  []int{},
		head: nil,
	},
	{
		arr:  []int{1},
		head: &ListNode{Val: 1},
	},
	{
		arr: []int{1, 2, 3, 4, 5},
		head: &ListNode{
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
		if !reflect.DeepEqual(ArrayToLinkedList(test.arr), test.head) {
			t.Fatalf("TestCase[%d]: linked lists are not same", idx)
		}
	}
}

func TestLinkedListToArray(t *testing.T) {
	for idx, test := range linkedListTests {
		if !reflect.DeepEqual(LinkedListToArray(test.head), test.arr) {
			t.Fatalf("TestCase[%d]: linked lists are not same", idx)
		}
	}
}
