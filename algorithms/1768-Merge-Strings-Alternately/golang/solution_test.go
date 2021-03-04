package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	word1  string
	word2  string
	output string
}{
	{word1: "abc", word2: "pqr", output: "apbqcr"},
	{word1: "ab", word2: "pqrs", output: "apbqrs"},
	{word1: "abcd", word2: "pq", output: "apbqcd"},
}

func TestMergeAlternately(t *testing.T) {
	for idx, test := range tests {
		actual := mergeAlternately(test.word1, test.word2)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: word1 = %s word2 = %s, expected %s but get %s",
				idx,
				test.word1,
				test.word2,
				test.output,
				actual,
			)
		}
	}
}
