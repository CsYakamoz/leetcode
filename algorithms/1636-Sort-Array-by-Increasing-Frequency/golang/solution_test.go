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
		input:  []int{1, 1, 2, 2, 2, 3},
		output: []int{3, 1, 1, 2, 2, 2},
	},
	{
		input:  []int{2, 3, 1, 3, 2},
		output: []int{1, 3, 3, 2, 2},
	},
	{
		input:  []int{-1, 1, -6, 4, 5, -6, 1, 4, 1},
		output: []int{5, -1, 4, 4, -6, -6, 1, 1, 1},
	},
}

func TestFrequencySort(t *testing.T) {
	for idx, test := range tests {
		actual := frequencySort(test.input)
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
