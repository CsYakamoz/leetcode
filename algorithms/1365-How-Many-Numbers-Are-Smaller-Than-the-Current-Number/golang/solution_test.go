package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  []int
	output []int
}{
	{input: []int{8, 1, 2, 2, 3}, output: []int{4, 0, 1, 1, 3}},
	{input: []int{6, 5, 4, 8}, output: []int{2, 1, 0, 3}},
	{input: []int{7, 7, 7, 7}, output: []int{0, 0, 0, 0}},
}

func TestSmallerNumbersThanCurrent(t *testing.T) {
	for idx, test := range tests {
		actual := smallerNumbersThanCurrent(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: nums = %v, expected %v but get %v",
				idx,
				test.input,
				test.output,
				actual,
			)
		}
	}
}
