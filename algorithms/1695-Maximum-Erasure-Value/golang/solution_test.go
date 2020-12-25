package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  []int
	output int
}{
	{input: []int{4, 2, 4, 5, 6}, output: 17},
	{input: []int{5, 2, 1, 2, 5, 2, 1, 2, 5}, output: 8},
	{input: []int{7, 5, 4, 4, 3, 1, 1, 2, 3, 4, 5}, output: 16},
}

func TestMaximumUniqueSubarray(t *testing.T) {
	for idx, test := range tests {
		actual := maximumUniqueSubarray(test.input)
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
