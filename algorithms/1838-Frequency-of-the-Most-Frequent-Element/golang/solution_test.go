package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	nums   []int
	k      int
	output int
}{
	{nums: []int{1, 2, 4}, k: 5, output: 3},
	{nums: []int{1, 4, 8, 13}, k: 5, output: 2},
	{nums: []int{3, 9, 6}, k: 2, output: 1},
}

func TestMaxFrequency(t *testing.T) {
	for idx, test := range tests {
		actual := maxFrequency(test.nums, test.k)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: nums = %v, k = %d, expected %d but get %d",
				idx,
				test.nums,
				test.k,
				test.output,
				actual,
			)
		}
	}
}
