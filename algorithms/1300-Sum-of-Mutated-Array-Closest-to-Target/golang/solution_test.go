package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arr    []int
	target int
	output int
}{
	{
		arr:    []int{4, 9, 3},
		target: 10,
		output: 3,
	},
	{
		arr:    []int{2, 3, 5},
		target: 10,
		output: 5,
	},
	{
		arr:    []int{60864, 25176, 27249, 21296, 20204},
		target: 56803,
		output: 11361,
	},
	{
		arr:    []int{1, 2, 3, 4, 5},
		target: 3,
		output: 1,
	},
	{
		arr:    []int{3, 4, 5, 9},
		target: 13,
		output: 3,
	},
}

func TestFindBestValue(t *testing.T) {
	for idx, test := range tests {
		actual := findBestValue(test.arr, test.target)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: arr = %v, target = %d, expected %d but get %d",
				idx,
				test.arr,
				test.target,
				test.output,
				actual,
			)
		}
	}
}
