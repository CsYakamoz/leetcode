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
			{1, 1, 0},
			{1, 0, 1},
			{0, 0, 0},
		},
		output: [][]int{
			{1, 0, 0},
			{0, 1, 0},
			{1, 1, 1},
		},
	},
	{
		input: [][]int{
			{1, 1, 0, 0},
			{1, 0, 0, 1},
			{0, 1, 1, 1},
			{1, 0, 1, 0},
		},
		output: [][]int{
			{1, 1, 0, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 1},
			{1, 0, 1, 0},
		},
	},
}

func TestFlipAndInvertImage(t *testing.T) {
	for idx, test := range tests {
		actual := flipAndInvertImage(test.input)
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

