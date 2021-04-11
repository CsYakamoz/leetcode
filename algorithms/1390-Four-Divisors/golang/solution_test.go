package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	nums   []int
	output int
}{
	{nums: []int{21, 4, 7}, output: 32},
	{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 16}, output: 27},
}

func TestSumFourDivisors(t *testing.T) {
	for idx, test := range tests {
		actual := sumFourDivisors(test.nums)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf("TestCase[%d]: nums = %v, expected %v but get %v",
				idx,
				test.nums,
				test.output,
				actual,
			)
		}
	}
}
