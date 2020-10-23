package utils

import (
	"github.com/CsYakamoz/leetcode/lib/golang/structure"
)

type TreeNode = structure.TreeNode

func PointArrayToBinaryTree(arr []*int) *TreeNode {
	length := len(arr)
	if length == 0 {
		return nil
	}

	root := &TreeNode{Val: *arr[0]}
	queue := make([]*TreeNode, 0, length)
	queue = append(queue, root)

	isLeftChild := true
	for _, p := range arr[1:] {
		var curr *TreeNode = nil
		if p != nil {
			curr = &TreeNode{Val: *p}
		}

		if isLeftChild {
			queue[0].Left = curr
		} else {
			queue[0].Right = curr
			queue = queue[1:]
		}

		if curr != nil {
			queue = append(queue, curr)
		}

		isLeftChild = !isLeftChild
	}

	return root
}

func PreOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}

	var result []int
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, curr.Val)

		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}

		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}

	return result
}

func InOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}

	var result []int
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			curr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result = append(result, curr.Val)
			root = curr.Right
		}
	}

	return result
}

// TODO: loop method instead of recursion method
func PostOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}

	result := make([]int, 0)

	result = append(result, PostOrderTraversal(root.Left)...)
	result = append(result, PostOrderTraversal(root.Right)...)
	result = append(result, root.Val)

	return result
}
