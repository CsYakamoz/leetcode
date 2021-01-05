package golang

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-5

func floatEqual(x, y float64) bool {
	return math.Abs(x-y) <= float64EqualityThreshold
}

var tests = []struct {
	hour    int
	minutes int
	output  float64
}{
	{hour: 12, minutes: 30, output: 165},
	{hour: 3, minutes: 30, output: 75},
	{hour: 3, minutes: 15, output: 7.5},
	{hour: 4, minutes: 50, output: 155},
	{hour: 12, minutes: 0, output: 0},
	{hour: 1, minutes: 57, output: 76.50000},
}

func TestAngleClock(t *testing.T) {
	for idx, test := range tests {
		actual := angleClock(test.hour, test.minutes)
		if !floatEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: hour = %d, minutes = %d, expected %f but get %f",
				idx,
				test.hour,
				test.minutes,
				test.output,
				actual,
			)
		}
	}
}
