package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  []int
	output int
}{
	{input: []int{2, 1, 4, 7, 3, 2, 5}, output: 5},
	{input: []int{2, 2, 2}, output: 0},
	{input: []int{1, 2, 3, 2, 1}, output: 5},
	{input: []int{1, 2, 3}, output: 0},
	{input: []int{1, 2, 3, 2}, output: 4},
	{input: []int{0, 2, 2}, output: 0},
	{input: []int{875, 884, 239, 731, 723, 685}, output: 4},
}

func TestLongestMountain(t *testing.T) {
	for idx, test := range tests {
		actual := longestMountain(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: arr = %v, expected %d but get %d",
				idx,
				test.input,
				test.output,
				actual,
			)
		}
	}
}
