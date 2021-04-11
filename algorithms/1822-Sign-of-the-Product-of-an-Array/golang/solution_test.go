package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	nums   []int
	output int
}{
	{nums: []int{-1, -2, -3, -4, 3, 2, 1}, output: 1},
	{nums: []int{1, 5, 0, 2, -3}, output: 0},
	{nums: []int{-1, 1, -1, 1, -1}, output: -1},
}

func TestArraySign(t *testing.T) {
	for idx, test := range tests {
		actual := arraySign(test.nums)
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
