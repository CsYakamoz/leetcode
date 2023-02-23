package golang

import (
	"reflect"
	"testing"

	"github.com/CsYakamoz/leetcode/lib/golang/utils"
)

var tests = []struct {
	input  *ListNode
	output *ListNode
}{
	{
		input:  utils.ArrayToLinkedList([]int{5, 2, 13, 3, 8}),
		output: utils.ArrayToLinkedList([]int{13, 8}),
	},
	{
		input:  utils.ArrayToLinkedList([]int{1, 1, 1, 1}),
		output: utils.ArrayToLinkedList([]int{1, 1, 1, 1}),
	},
}

func TestRemoveNodes(t *testing.T) {
	for idx, test := range tests {
		actual := removeNodes(utils.CopyLinkList(test.input))
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d] input = %v, expected = %v but get %v",
				idx,
				utils.LinkedListToArray(test.input),
				utils.LinkedListToArray(test.output),
				utils.LinkedListToArray(actual),
			)
		}
	}
}

