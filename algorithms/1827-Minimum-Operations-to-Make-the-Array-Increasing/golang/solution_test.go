package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	nums   []int
	output int
}{
	{nums: []int{1, 1, 1}, output: 3},
	{nums: []int{1, 5, 2, 4, 1}, output: 14},
	{nums: []int{8}, output: 0},
}

func TestMinOperations(t *testing.T) {
	for idx, test := range tests {
		actual := minOperations(test.nums)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: nums = %v, expected %d but get %d",
				idx,
				test.nums,
				test.output,
				actual,
			)
		}
	}
}
