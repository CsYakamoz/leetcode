package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  int
	output int
}{
	{input: 1, output: 0},
	{input: 2, output: 1},
	{input: 3, output: 2},
	{input: 5, output: 3},
	{input: 509, output: 11},
	{input: 8, output: 3},
	{input: 7, output: 4},
	{input: 4, output: 2},
	{input: 33, output: 6},
	{input: 35, output: 7},
	{input: 100000000, output: 31},
	{input: 2147483647, output: 32},
}

func TestIntegerReplacement(t *testing.T) {
	for idx, test := range tests {
		actual := integerReplacement(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: n = %d, expected %d but get %d",
				idx,
				test.input,
				test.output,
				actual,
			)
		}
	}
}
