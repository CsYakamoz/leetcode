package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  []int
	output int
}{
	{input: []int{4, 14, 2}, output: 6},
}

func TestTotalHammingDistance(t *testing.T) {
	for idx, test := range tests {
		actual := totalHammingDistance(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: nums = %v, expected %d but get %d",
				idx,
				test.input,
				test.output,
				actual,
			)
		}
	}
}
