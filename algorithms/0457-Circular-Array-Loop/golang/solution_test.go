package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  []int
	output bool
}{
	{input: []int{2, -1, 1, 2, 2}, output: true},
	{input: []int{-1, 2}, output: false},
	{input: []int{1, 2}, output: false},
	{input: []int{-2, 1, -1, -2, -2}, output: false},
	// test-case 12
	{input: []int{2, 2, 2, 2, 2, 4, 7}, output: false},
	// test-case 33
	{input: []int{-2, -3, -9}, output: false},
}

func TestCircularArrayLoop(t *testing.T) {
	for idx, test := range tests {
		copyData := make([]int, len(test.input), len(test.input))
		copy(copyData, test.input)

		actual := circularArrayLoop(copyData)
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
