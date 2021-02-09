package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  []int
	output int
}{
	{input: []int{1, -3, 2, 3, -4}, output: 5},
	{input: []int{2, -5, 1, -4, 3, -2}, output: 8},
}

func TestMaxAbsoluteSum(t *testing.T) {
	for idx, test := range tests {
		actual := maxAbsoluteSum(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: nums = %v, expected %d but get %d",
				idx,
				test.input,
				test.output,
				actual,
			)
		}
	}
}
