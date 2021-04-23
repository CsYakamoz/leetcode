package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	sentence string
	output   bool
}{
	{sentence: "thequickbrownfoxjumpsoverthelazydog", output: true},
	{sentence: "leetcode", output: false},
}

func TestCheckIfPangram(t *testing.T) {
	for idx, test := range tests {
		actual := checkIfPangram(test.sentence)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: sentence = %s, expected %v but get %v",
				idx,
				test.sentence,
				test.output,
				actual,
			)
		}
	}
}
