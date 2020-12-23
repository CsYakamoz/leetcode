package golang

import (
	"github.com/CsYakamoz/leetcode/lib/golang/utils"
	"reflect"
	"testing"
)

var tests = []struct {
	input  []int
	output []int
}{
	{
		input:  []int{1, 2, -3, 3, 1},
		output: []int{3, 1},
	},
	{
		input:  []int{1, 2, 3, -3, 4},
		output: []int{1, 2, 4},
	},
	{
		input:  []int{1, 2, 3, -3, -2},
		output: []int{1},
	},
	{
		input:  []int{1, -1, 1, -1},
		output: []int{},
	},
	{
		input:  []int{1, 3, -2, 2, 1, -1, 3, -3, 2},
		output: []int{1, 3, 2},
	},
	{
		input:  []int{1, 3, 2, -3, -2, 5, 5, -5, 1},
		output: []int{1, 5, 1},
	},
}

func TestRemoveZeroSumSublists(t *testing.T) {
	for idx, test := range tests {
		actual := utils.LinkedListToArray(
			removeZeroSumSublists(
				utils.ArrayToLinkedList(test.input),
			),
		)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: head = %v, expected %v but get %v",
				idx,
				test.input,
				test.output,
				actual,
			)
		}
	}
}
