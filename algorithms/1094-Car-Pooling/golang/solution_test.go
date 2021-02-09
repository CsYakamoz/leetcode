package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	trips    [][]int
	capacity int
	output   bool
}{
	{
		trips:    [][]int{{2, 1, 5}, {3, 3, 7}},
		capacity: 4,
		output:   false,
	},
	{
		trips:    [][]int{{2, 1, 5}, {3, 3, 7}},
		capacity: 5,
		output:   true,
	},
	{
		trips:    [][]int{{2, 1, 5}, {3, 5, 7}},
		capacity: 3,
		output:   true,
	},
	{
		trips:    [][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}},
		capacity: 11,
		output:   true,
	},
}

func TestCarPooling(t *testing.T) {
	for idx, test := range tests {
		cpData := make([][]int, len(test.trips), len(test.trips))
		copy(cpData, test.trips)

		actual := carPooling(cpData, test.capacity)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: trips = %v capacity = %d, expected %t but get %t",
				idx,
				test.trips,
				test.capacity,
				test.output,
				actual,
			)
		}
	}
}
