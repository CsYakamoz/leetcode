package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  []int
	output []int
}{
	{
		input:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		output: []int{0, 1, 2, 4, 8, 3, 5, 6, 7},
	},
	{
		input:  []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1},
		output: []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
	},
	{
		input:  []int{10000, 10000},
		output: []int{10000, 10000},
	},
	{
		input:  []int{2, 3, 5, 7, 11, 13, 17, 19},
		output: []int{2, 3, 5, 17, 7, 11, 13, 19},
	},
	{
		input:  []int{10, 100, 1000, 10000},
		output: []int{10, 100, 10000, 1000},
	},
}

func TestSortByBits(t *testing.T) {
	for idx, test := range tests {
		actual := sortByBits(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: arr = %v, expected %v but get %v",
				idx,
				test.input,
				test.output,
				actual,
			)
		}
	}
}
