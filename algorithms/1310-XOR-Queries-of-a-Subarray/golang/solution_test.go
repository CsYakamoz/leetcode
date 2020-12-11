package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arr     []int
	queries [][]int
	output  []int
}{
	{
		arr:     []int{1, 3, 4, 8},
		queries: [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}},
		output:  []int{2, 7, 14, 8},
	},
	{
		arr:     []int{4, 8, 2, 10},
		queries: [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}},
		output:  []int{8, 0, 4, 4},
	},
}

func TestXorQueries(t *testing.T) {
	for idx, test := range tests {
		actual := xorQueries(test.arr, test.queries)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: arr = %v queries = %v, expected %v but get %v",
				idx,
				test.arr,
				test.queries,
				test.output,
				actual,
			)
		}
	}
}
