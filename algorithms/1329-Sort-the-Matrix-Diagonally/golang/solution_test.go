package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  [][]int
	output [][]int
}{
	{
		input: [][]int{
			{3, 3, 1, 1},
			{2, 2, 1, 2},
			{1, 1, 1, 2},
		},
		output: [][]int{
			{1, 1, 1, 1},
			{1, 2, 2, 2},
			{1, 2, 3, 3},
		},
	},
}

func TestDiagonalSort(t *testing.T) {
	for idx, test := range tests {
		copyInput := make([][]int, len(test.input), len(test.input))
		copy(copyInput, test.input)

		actual := diagonalSort(copyInput)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: mat = %v, expected %v but get %v",
				idx,
				test.input,
				test.output,
				actual,
			)
		}
	}
}

