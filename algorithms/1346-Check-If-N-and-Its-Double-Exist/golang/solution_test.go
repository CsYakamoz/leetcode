package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  []int
	output bool
}{
	{input: []int{10, 2, 5, 3}, output: true},
	{input: []int{7, 1, 14, 11}, output: true},
	{input: []int{3, 1, 7, 11}, output: false},
}

func TestCheckIfExist(t *testing.T) {
	for idx, test := range tests {
		actual := checkIfExist(test.input)
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

