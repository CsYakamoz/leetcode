package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  string
	output int64
}{
	{input: "aba", output: 4},
	{input: "aabb", output: 9},
	{input: "he", output: 2},
	{input: "abcdefghijjihabcdefg", output: 32},
	// {input: json.Un, output: 37911166},
}

func TestWonderfulSubstrings(t *testing.T) {
	for idx, test := range tests {
		actual := wonderfulSubstrings(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: word = %s, expected %d but get %d",
				idx,
				test.input,
				test.output,
				actual,
			)
		}
	}
}
