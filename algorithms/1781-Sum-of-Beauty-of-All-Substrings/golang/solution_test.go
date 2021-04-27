package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	s      string
	output int
}{
	{s: "aabcb", output: 5},
	{s: "aabcbaa", output: 17},
}

func TestBeautySum(t *testing.T) {
	for idx, test := range tests {
		actual := beautySum(test.s)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: s = %s, expected %d but get %d",
				idx,
				test.s,
				test.output,
				actual,
			)
		}
	}
}
