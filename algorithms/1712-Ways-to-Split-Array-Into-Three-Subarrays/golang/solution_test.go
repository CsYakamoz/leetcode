package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  []int
	output int
}{
	{input: []int{1, 1, 1}, output: 1},
	{input: []int{1, 2, 2, 2, 5, 0}, output: 3},
	{input: []int{3, 2, 1}, output: 0},
	{input: []int{0, 0, 0, 0, 0, 0, 0, 0, 0}, output: 28},
	{input: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, output: 36},
	{input: []int{2, 2, 0, 0, 0, 0, 0, 0}, output: 0},
	{input: []int{7, 2, 5, 5, 6, 2, 10, 9}, output: 6},
}

func TestWaysToSplit(t *testing.T) {
	for idx, test := range tests {
		actual := waysToSplit(test.input)
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
