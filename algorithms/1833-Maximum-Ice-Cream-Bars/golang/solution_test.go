package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	costs  []int
	coins  int
	output int
}{
	{costs: []int{1, 3, 2, 4, 1}, coins: 7, output: 4},
	{costs: []int{10, 6, 8, 7, 7, 8}, coins: 5, output: 0},
	{costs: []int{1, 6, 3, 1, 2, 5}, coins: 20, output: 6},
}

func TestMaxIceCream(t *testing.T) {
	for idx, test := range tests {
		copyData := make([]int, len(test.costs), len(test.costs))
		copy(copyData, test.costs)

		actual := maxIceCream(copyData, test.coins)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: costs = %v coins = %d, expected %d but get %d",
				idx,
				test.costs,
				test.coins,
				test.output,
				actual,
			)
		}
	}
}
