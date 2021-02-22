package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  []string
	output int
}{
	{
		input:  []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"},
		output: 16,
	},
	{
		input:  []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"},
		output: 4,
	},
	{
		input:  []string{"a", "aa", "aaa", "aaaa"},
		output: 0,
	},
}

func TestMaxProduct(t *testing.T) {
	for idx, test := range tests {
		actual := maxProduct(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: words = %v, expected %d but get %d",
				idx,
				test.input,
				test.output,
				actual,
			)
		}
	}
}
