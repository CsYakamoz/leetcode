package golang

import (
	"reflect"
	"strconv"
	"testing"
)

type testDataType struct {
	val       int
	randomIdx string
}

var tests = [][]testDataType{
	{
		{val: 7, randomIdx: "null"},
		{val: 13, randomIdx: "0"},
		{val: 11, randomIdx: "4"},
		{val: 10, randomIdx: "2"},
		{val: 1, randomIdx: "0"},
	},
	{
		{val: 1, randomIdx: "1"},
		{val: 2, randomIdx: "1"},
	},
	{
		{val: 3, randomIdx: "null"},
		{val: 3, randomIdx: "0"},
		{val: 3, randomIdx: "null"},
	},
	{},
}

func ctorListWithRandomPointer(arr []testDataType) *Node {
	length := len(arr)
	if length == 0 {
		return nil
	}

	result := make([]*Node, length)
	for idx, item := range arr {
		result[idx] = &Node{Val: item.val}
	}

	for i := 0; i < length; i++ {
		item := arr[i]
		if i != length-1 {
			result[i].Next = result[i+1]
		}

		if item.randomIdx != "null" {
			idx, _ := strconv.Atoi(item.randomIdx)
			result[i].Random = result[idx]
		}
	}

	return result[0]
}

func TestCopyRandomList(t *testing.T) {
	for idx, test := range tests {
		head := ctorListWithRandomPointer(test)
		actual := copyRandomList(head)
		if !reflect.DeepEqual(actual, head) {
			t.Errorf("TestCase[%d]", idx)
		}
	}
}
