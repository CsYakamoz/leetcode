package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  string
	output int
}{
	{input: "abbcccaa", output: 13},
	{input: "xy", output: 2},
	{input: "zzzzz", output: 15},
}

func TestCountHomogenous(t *testing.T) {
	for idx, test := range tests {
		actual := countHomogenous(test.input)
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
