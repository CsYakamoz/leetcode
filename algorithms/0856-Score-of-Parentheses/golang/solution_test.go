package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  string
	output int
}{
	{input: "()", output: 1},
	{input: "(())", output: 2},
	{input: "()()", output: 2},
	{input: "(()(()))", output: 6},
	{input: "(()()()((())(())(())(())(())(())(())((((()))))()()()))", output: 138},
}

func TestScoreOfParentheses(t *testing.T) {
	for idx, test := range tests {
		actual := scoreOfParentheses(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: S = %s, expected %d but get %d",
				idx,
				test.input,
				test.output,
				actual,
			)
		}
	}
}
