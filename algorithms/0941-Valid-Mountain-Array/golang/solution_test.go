package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arr    []int
	output bool
}{
	// example
	{arr: []int{2, 1}, output: false},
	{arr: []int{3, 5, 5}, output: false},
	{arr: []int{0, 3, 2, 1}, output: true},

	{arr: []int{3, 4, 5}, output: false},
	{arr: []int{3, 4, 0}, output: true},
	{arr: []int{3, 4, 3}, output: true},
	{arr: []int{3, 4, 3, 3}, output: false},
	{arr: []int{3, 4, 3, 4}, output: false},
}

func TestValidMountainArray(t *testing.T) {
	for idx, test := range tests {
		actual := validMountainArray(test.arr)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: arr = %v, expected %v but get %v",
				idx,
				test.arr,
				test.output,
				actual,
			)
		}
	}
}
