package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  string
	output int
}{
	{input: "1101", output: 6},
	{input: "10", output: 1},
	{input: "1", output: 0},
}

func TestNumSteps(t *testing.T) {
	for idx, test := range tests {
		actual := numSteps(test.input)

		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: s = %s, expected %d but get %d",
				idx,
				test.input,
				test.output,
				actual,
			)
		}
	}
}
