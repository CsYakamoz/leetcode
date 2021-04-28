package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  string
	output string
}{
	{input: "aaaabbbbcccc", output: "abccbaabccba"},
	{input: "rat", output: "art"},
	{input: "leetcode", output: "cdelotee"},
	{input: "ggggggg", output: "ggggggg"},
	{input: "spo", output: "ops"},
}

func TestSortString(t *testing.T) {
	for idx, test := range tests {
		actual := sortString(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: s = %s, expected %s but get %s",
				idx,
				test.input,
				test.output,
				actual,
			)
		}
	}
}
