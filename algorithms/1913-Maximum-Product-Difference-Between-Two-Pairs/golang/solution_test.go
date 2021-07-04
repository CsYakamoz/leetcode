package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  []int
	output int
}{
	{input: []int{5, 6, 2, 7, 4}, output: 34},
	{input: []int{4, 2, 5, 9, 7, 4, 8}, output: 64},
}

func TestMaxProductDifference(t *testing.T) {
	for idx, test := range tests {
		copyData := make([]int, len(test.input), len(test.input))
		copy(copyData, test.input)

		actual := maxProductDifference(copyData)
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
