package golang

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/CsYakamoz/leetcode/lib/golang/utils"
)

var tests = []struct {
	arr    string
	output int
}{
	{
		arr:    "[2,3,1,3,1,null,1]",
		output: 2,
	},
	{
		arr:    "[2,1,1,1,3,null,null,null,null,null,1]",
		output: 1,
	},
	{
		arr:    "[9]",
		output: 1,
	},
}

func TestPseudoPalindromicPaths(t *testing.T) {
	for idx, test := range tests {
		var pointArr []*int
		err := json.Unmarshal([]byte(test.arr), &pointArr)
		if err != nil {
			t.Errorf("failed to parse json(%v)", test.arr)
			return
		}

		actual := pseudoPalindromicPaths(
			utils.PointArrayToBinaryTree(pointArr),
		)

		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: arr = %v, expected %d but get %d",
				idx,
				test.arr,
				test.output,
				actual,
			)
		}
	}
}
