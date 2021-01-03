package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	people []int
	limit  int
	output int
}{
	{people: []int{1, 2}, limit: 3, output: 1},
	{people: []int{3, 2, 2, 1}, limit: 3, output: 3},
	{people: []int{3, 5, 3, 4}, limit: 5, output: 4},
	{people: []int{1, 2, 2, 2}, limit: 3, output: 3},
}

func TestNumRescueBoats(t *testing.T) {
	for idx, test := range tests {
		copyData := make([]int, len(test.people), len(test.people))
		copy(copyData, test.people)
		actual := numRescueBoats(copyData, test.limit)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: people = %v, limit = %d, expected %d but get %d",
				idx,
				test.people,
				test.limit,
				test.output,
				actual,
			)
		}
	}
}
