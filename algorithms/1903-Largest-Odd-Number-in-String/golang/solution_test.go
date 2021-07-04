package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  string
	output string
}{
	{input: "52", output: "5"},
	{input: "4206", output: ""},
	{input: "35427", output: "35427"},
}

func TestLargestOddNumber(t *testing.T) {
	for idx, test := range tests {
		actual := largestOddNumber(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: num = %s, expected %s but get %s",
				idx,
				test.input,
				test.output,
				actual,
			)
		}
	}
}
