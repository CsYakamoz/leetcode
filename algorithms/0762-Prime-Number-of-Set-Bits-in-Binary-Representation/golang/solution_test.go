package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	L      int
	R      int
	output int
}{
	{L: 6, R: 10, output: 4},
	{L: 10, R: 15, output: 5},
	{L: 842, R: 888, output: 23},
	{L: 100, R: 999, output: 466},
	{L: 4, R: 85, output: 55},
}

func TestCountPrimeSetBits(t *testing.T) {
	for idx, test := range tests {
		actual := countPrimeSetBits(test.L, test.R)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: L = %d R = %d, expected %d but get %d",
				idx,
				test.L,
				test.R,
				test.output,
				actual,
			)
		}
	}
}
