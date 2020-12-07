package golang

import (
	"github.com/CsYakamoz/leetcode/lib/golang/utils"
	"reflect"
	"testing"
)

var tests = []struct {
	list1  []int
	a      int
	b      int
	list2  []int
	output []int
}{

	{
		list1:  []int{0, 1, 2, 3, 4, 5},
		a:      3,
		b:      4,
		list2:  []int{1000000, 1000001, 1000002},
		output: []int{0, 1, 2, 1000000, 1000001, 1000002, 5},
	},
	{
		list1:  []int{0, 1, 2, 3, 4, 5, 6},
		a:      2,
		b:      5,
		list2:  []int{1000000, 1000001, 1000002, 1000003, 1000004},
		output: []int{0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6},
	},
}

func TestMergeInBetween(t *testing.T) {
	for idx, test := range tests {
		actual := utils.LinkedListToArray(mergeInBetween(
			utils.ArrayToLinkedList(test.list1),
			test.a,
			test.b,
			utils.ArrayToLinkedList(test.list2),
		))
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: list1 = %v, a = %d, b = %d, list2 = %v, expected %v but get %v",
				idx,
				test.list1,
				test.a,
				test.b,
				test.list2,
				test.output,
				actual,
			)
		}
	}
}
