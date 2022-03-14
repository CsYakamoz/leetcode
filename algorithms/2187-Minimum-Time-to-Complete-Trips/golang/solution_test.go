package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	time       []int
	totalTrips int
	output     int64
}{
	/*
	 * {time: []int{1, 2, 3}, totalTrips: 5, output: 3},
	 * {time: []int{2}, totalTrips: 1, output: 2},
	 * {
	 *     time:       []int{3, 2, 3, 2, 1, 3, 5, 6, 2, 1, 3, 7, 100},
	 *     totalTrips: 10000000,
	 *     output:     1868162,
	 * },
	 */
	{time: []int{5, 10, 10}, totalTrips: 9, output: 25},
}

func TestMinimumTime(t *testing.T) {
	for idx, test := range tests {
		copyData := make([]int, len(test.time), len(test.time))
		copy(copyData, test.time)

		actual := minimumTime(copyData, test.totalTrips)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: times = %v totalTrips = %d, expected %d but get %d",
				idx,
				test.time,
				test.totalTrips,
				test.output,
				actual,
			)
		}
	}
}
