package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	stockPrices [][]int
	output      int
}{
	{
		stockPrices: [][]int{{1, 7}, {2, 6}, {3, 5}, {4, 4}, {5, 4}, {6, 3}, {7, 2}, {8, 1}},
		output:      3,
	},
	{
		stockPrices: [][]int{{3, 4}, {1, 2}, {7, 8}, {2, 3}},
		output:      1,
	},
	{
		stockPrices: [][]int{{1, 1}, {4, 2}, {7, 3}},
		output:      1,
	},
	{
		stockPrices: [][]int{{1, 1}, {500000000, 499999999}, {1000000000, 999999998}},
		output:      2,
	},
}

func TestMinimumLines(t *testing.T) {
	for idx, test := range tests {
		copyData := make([][]int, len(test.stockPrices), len(test.stockPrices))
		copy(copyData, test.stockPrices)

		actual := minimumLines(copyData)

		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: stockPrices is %v, expected %d but get %d",
				idx,
				test.stockPrices,
				test.output,
				actual,
			)
		}
	}
}
