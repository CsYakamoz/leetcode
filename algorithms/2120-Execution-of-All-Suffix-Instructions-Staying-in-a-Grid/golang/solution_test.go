package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	n        int
	startPos []int
	s        string
	output   []int
}{
	/*
	 * {n: 3, startPos: []int{0, 1}, s: "RRDDLU", output: []int{1, 5, 4, 3, 1, 0}},
	 * {n: 2, startPos: []int{1, 1}, s: "LURD", output: []int{4, 1, 0, 0}},
	 * {n: 1, startPos: []int{0, 0}, s: "LRUD", output: []int{0, 0, 0, 0}},
	 */
	{n: 2, startPos: []int{0, 0}, s: "RDLURDLURDLURLDU", output: []int{16, 1, 0, 0, 12, 1, 0, 0, 8, 1, 0, 0, 4, 0, 2, 0}},
}

func TestExecuteInstructions(t *testing.T) {
	for idx, test := range tests {
		actual := executeInstructions(test.n, test.startPos, test.s)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: n = %d startPos = %v s = %s, expected %v but get %v",
				idx,
				test.n,
				test.startPos,
				test.s,
				test.output,
				actual,
			)
		}
	}
}
