package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	pushed []int
	popped []int
	output bool
}{
	{pushed: []int{1, 2, 3, 4, 5}, popped: []int{4, 5, 3, 2, 1}, output: true},
	{pushed: []int{1, 2, 3, 4, 5}, popped: []int{4, 3, 5, 1, 2}, output: true},
}

func TestValidateStackSequences(t *testing.T) {
	for idx, test := range tests {
		copyPushed := make([]int, len(test.pushed))
		copy(copyPushed, test.pushed)

		copyPopped := make([]int, len(test.popped))
		copy(copyPopped, test.popped)

		actual := validateStackSequences(copyPushed, copyPopped)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: pushed = %v popped = %v, expected = %v but get %v",
				idx,
				test.pushed,
				test.popped,
				test.output,
				actual,
			)
		}
	}
}
