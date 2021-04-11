package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	s1     string
	s2     string
	output bool
}{
	{s1: "bank", s2: "kanb", output: true},
	{s1: "attack", s2: "defend", output: false},
	{s1: "kelb", s2: "kelb", output: true},
	{s1: "abcd", s2: "dcba", output: false},
	{s1: "aa", s2: "bb", output: false},
}

func TestAreAlmostEqual(t *testing.T) {
	for idx, test := range tests {
		actual := areAlmostEqual(test.s1, test.s2)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: s1 = %s s2 = %s, expected %v but get %v",
				idx,
				test.s1,
				test.s2,
				test.output,
				actual,
			)
		}
	}
}
