package golang

import (
	"github.com/CsYakamoz/leetcode/lib/golang/structure"
)

type TreeNode = structure.TreeNode

func isPseudoPalindromic(iter *[]int) int {
	singleCount, dict := 0, [10]int{}

	for _, val := range *iter {
		dict[val]++
		if dict[val]%2 == 1 {
			singleCount++
		} else {
			singleCount--
		}
	}

	if singleCount <= 1 {
		return 1
	} else {
		return 0
	}
}

func recursion(root *TreeNode, iter *[]int, count *int) {
	if root == nil {
		return
	}

	*iter = append(*iter, root.Val)

	if root.Left == nil && root.Right == nil {
		*count += isPseudoPalindromic(iter)
	} else {
		recursion(root.Left, iter, count)
		recursion(root.Right, iter, count)
	}

	*iter = (*iter)[:len(*iter)-1]
}

func pseudoPalindromicPaths(root *TreeNode) int {
	if root == nil {
		return 0
	}

	count, iter := 0, make([]int, 0)
	recursion(root, &iter, &count)

	return count
}
