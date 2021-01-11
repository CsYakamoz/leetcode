package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  []int
	output int
}{
	/*
	 * {input: []int{1, 3, 5, 7, 9}, output: 4},
	 * {input: []int{1, 1, 1, 3, 3, 3, 7}, output: 15},
	 * {input: []int{149, 107, 1, 63, 0, 1, 6867, 1325, 5611, 2581, 39, 89, 46, 18, 12, 20, 22, 234}, output: 12},
	 */
	{input: []int{2160, 1936, 3, 29, 27, 5, 2503, 1593, 2, 0, 16, 0, 3860, 28908, 6, 2, 15, 49, 6246, 1946, 23, 105, 7996, 196, 0, 2, 55, 457, 5, 3, 924, 7268, 16, 48, 4, 0, 12, 116, 2628, 1468}, output: 53},
}

func TestCountPairs(t *testing.T) {
	for idx, test := range tests {
		copyData := make([]int, len(test.input), len(test.input))
		copy(copyData, test.input)

		actual := countPairs(copyData)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: deliciousness = %v, expected %d but get %d",
				idx,
				test.input,
				test.output,
				actual,
			)
		}
	}
}
