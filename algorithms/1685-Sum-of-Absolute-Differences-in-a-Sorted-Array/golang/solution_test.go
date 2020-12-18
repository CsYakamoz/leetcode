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
		input: []int{2, 3, 5}, output: []int{4, 3, 5},
	},
	{
		input: []int{1, 4, 6, 8, 10}, output: []int{24, 15, 13, 15, 21},
	},
}

func TestGetSumAbsoluteDifferences(t *testing.T) {
	for idx, test := range tests {
		actual := getSumAbsoluteDifferences(test.input)
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
