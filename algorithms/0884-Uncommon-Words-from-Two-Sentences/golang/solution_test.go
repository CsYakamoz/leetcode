package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	A      string
	B      string
	output []string
}{
	{
		A:      "this apple is sweet",
		B:      "this apple is sour",
		output: []string{"sweet", "sour"},
	},
	{
		A:      "apple apple",
		B:      "banana",
		output: []string{"banana"},
	},
}

func TestUncommonFromSentences(t *testing.T) {
	for idx, test := range tests {
		actual := uncommonFromSentences(test.A, test.B)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: A = \"%s\" B = \"%s\", expected %v but get %v",
				idx,
				test.A,
				test.B,
				test.output,
				actual,
			)
		}
	}
}
