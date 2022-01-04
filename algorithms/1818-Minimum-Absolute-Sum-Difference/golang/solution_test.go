package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	nums1  []int
	nums2  []int
	output int
}{
	{nums1: []int{1, 7, 5}, nums2: []int{2, 3, 5}, output: 3},
	{nums1: []int{2, 4, 6, 8, 10}, nums2: []int{2, 4, 6, 8, 10}, output: 0},
	{nums1: []int{1, 10, 4, 4, 2, 7}, nums2: []int{9, 3, 5, 1, 7, 4}, output: 20},
	{nums1: []int{1, 2, 3, 4, 5}, nums2: []int{1, 4, 3, 9, 1000}, output: 1000},
}

func TestMinAbsoluteSumDiff(t *testing.T) {
	for idx, test := range tests {
		cpData := make([]int, len(test.nums1), len(test.nums1))
		copy(cpData, test.nums1)

		actual := minAbsoluteSumDiff(cpData, test.nums2)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: nums1 = %v nums2 = %v, expected %d but get %d",
				idx,
				test.nums1,
				test.nums2,
				test.output,
				actual,
			)
		}
	}
}
