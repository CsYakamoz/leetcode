package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	nums   []int
	x      int
	output int
}{
	{nums: []int{1, 1, 4, 2, 3}, x: 5, output: 2},
	{nums: []int{5, 6, 7, 8, 9}, x: 4, output: -1},
	{nums: []int{3, 2, 20, 1, 1, 3}, x: 10, output: 5},
	{nums: []int{1, 1}, x: 3, output: -1},
}

func TestMinOperations(t *testing.T) {
	for idx, test := range tests {
		actual := minOperations(test.nums, test.x)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: nums = %v x = %d, expected %d but get %d",
				idx,
				test.nums,
				test.x,
				test.output,
				actual,
			)
		}
	}
}
