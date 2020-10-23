package utils

import (
	"encoding/json"
	"reflect"
	"testing"
)

var treeTests = []struct {
	jsonStr   []byte
	treeRoot  *TreeNode
	preOrder  []int
	inOrder   []int
	postOrder []int
}{
	{
		// empty tree
		jsonStr:   []byte(`[]`),
		treeRoot:  nil,
		preOrder:  []int{},
		inOrder:   []int{},
		postOrder: []int{},
	},
	{
		/*
		 *    1
		 *  /   \
		 * 2     3
		 *  \     \
		 *   5     4
		 */
		jsonStr: []byte(`[1, 2, 3, null, 5, null, 4]`),
		treeRoot: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 5}},
			Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
		},
		preOrder:  []int{1, 2, 5, 3, 4},
		inOrder:   []int{2, 5, 1, 3, 4},
		postOrder: []int{5, 2, 4, 3, 1},
	},
	{
		/*
		 *    6
		 *  /   \
		 * 3     5
		 *  \    /
		 *   2  0
		 *    \
		 *     1
		 */
		jsonStr: []byte(`[6, 3, 5, null, 2, 0, null, null, 1]`),
		treeRoot: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:   3,
				Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}},
			Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 0}},
		},
		preOrder:  []int{6, 3, 2, 1, 5, 0},
		inOrder:   []int{3, 2, 1, 6, 0, 5},
		postOrder: []int{1, 2, 3, 0, 5, 6},
	},
}

func isSameTree(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 != nil && root2 != nil {
		return root1.Val == root2.Val &&
			isSameTree(root1.Left, root2.Left) &&
			isSameTree(root1.Right, root2.Right)
	}

	if root1 == nil && root2 == nil {
		return true
	}

	return false
}

func TestPointArrayToBinaryTree(t *testing.T) {
	var pointArr []*int

	for idx, item := range treeTests {
		err := json.Unmarshal(item.jsonStr, &pointArr)
		if err != nil {
			t.Fatalf(err.Error())
		}

		if !isSameTree(item.treeRoot, PointArrayToBinaryTree(pointArr)) {
			t.Fatalf("TestCase[%d]: trees are not same", idx)
		}
	}
}

func TestPreOrderTraversal(t *testing.T) {
	var pointArr []*int

	for idx, item := range treeTests {
		json.Unmarshal(item.jsonStr, &pointArr)

		if !reflect.DeepEqual(
			item.preOrder,
			PreOrderTraversal(PointArrayToBinaryTree(pointArr)),
		) {
			t.Errorf("TestCase[%d]: preOrders is not same", idx)
		}
	}
}

func TestInOrderTraversal(t *testing.T) {
	var pointArr []*int

	for idx, item := range treeTests {
		json.Unmarshal(item.jsonStr, &pointArr)

		if !reflect.DeepEqual(
			item.inOrder,
			InOrderTraversal(PointArrayToBinaryTree(pointArr)),
		) {
			t.Errorf("TestCase[%d]: inOrders is not same", idx)
		}
	}
}

func TestPostOrderTraversal(t *testing.T) {
	var pointArr []*int

	for idx, item := range treeTests {
		json.Unmarshal(item.jsonStr, &pointArr)

		if !reflect.DeepEqual(
			item.postOrder,
			PostOrderTraversal(PointArrayToBinaryTree(pointArr)),
		) {
			t.Errorf("TestCase[%d]: postOrders is not same", idx)
		}
	}
}
