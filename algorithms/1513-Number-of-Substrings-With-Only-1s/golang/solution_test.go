package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  string
	output int
}{
	{input: "0110111", output: 9},
	{input: "101", output: 2},
	{input: "111111", output: 21},
	{input: "000", output: 0},
}

func TestNumSub(t *testing.T) {
	for idx, test := range tests {
		actual := numSub(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: s = %s, expected %d but get %d",
				idx,
				test.input,
				test.output,
				actual,
			)
		}
	}
}
