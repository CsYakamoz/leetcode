package golang

import (
	"github.com/CsYakamoz/leetcode/lib/golang/utils"
	"reflect"
	"testing"
)

var tests = []struct {
	input  []int
	output int
}{
	{
		input:  []int{1, 0, 1},
		output: 5,
	},
	{
		input:  []int{0},
		output: 0,
	},
	{
		input:  []int{1},
		output: 1,
	},
	{
		input:  []int{1, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 0},
		output: 1143510,
	},
}

func TestGetDecimalValue(t *testing.T) {
	for idx, test := range tests {
		actual := getDecimalValue(utils.ArrayToLinkedList(test.input))
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: input is %v, expected %d but get %d",
				idx,
				test.input,
				test.output,
				actual,
			)
		}
	}
}

