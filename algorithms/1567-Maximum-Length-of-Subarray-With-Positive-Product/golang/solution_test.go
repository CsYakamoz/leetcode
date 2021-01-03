package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  []int
	output int
}{
	{input: []int{1, -2, -3, 4}, output: 4},
	{input: []int{0, 1, -2, -3, -4}, output: 3},
	{input: []int{-1, -2, -3, 0, 1}, output: 2},
	{input: []int{-1, 2}, output: 1},
	{input: []int{1, 2, 3, 5, -6, 4, 0, 10}, output: 4},
}

func TestGetMaxLen(t *testing.T) {
	for idx, test := range tests {
		actual := getMaxLen(test.input)
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
