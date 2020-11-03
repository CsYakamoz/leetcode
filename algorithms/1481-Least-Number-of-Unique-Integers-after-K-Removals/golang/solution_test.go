package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arr    []int
	k      int
	output int
}{
	{
		arr:    []int{5, 5, 4},
		k:      1,
		output: 1,
	},
	{
		arr:    []int{4, 3, 1, 1, 3, 3, 2},
		k:      3,
		output: 2,
	},
}

func TestFindLeastNumOfUniqueInts(t *testing.T) {
	for idx, test := range tests {
		actual := findLeastNumOfUniqueInts(test.arr, test.k)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: arr = %v k = %d, expected %d but gets %d",
				idx,
				test.arr,
				test.k,
				test.output,
				actual,
			)
		}
	}
}
