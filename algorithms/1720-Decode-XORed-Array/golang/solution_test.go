package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	encoded []int
	first   int
	output  []int
}{
	{
		encoded: []int{1, 2, 3},
		first:   1,
		output:  []int{1, 0, 2, 1},
	},
	{
		encoded: []int{6, 2, 7, 3},
		first:   4,
		output:  []int{4, 2, 0, 7, 4},
	},
}

func TestDecode(t *testing.T) {
	for idx, test := range tests {
		actual := decode(test.encoded, test.first)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: encoded = %v first = %d, expected %v but get %v",
				idx,
				test.encoded,
				test.first,
				actual,
				test.output,
			)
		}
	}
}
