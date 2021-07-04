package golang

import "sort"

func maxProductDifference(nums []int) int {
	sort.Ints(nums)

	length := len(nums)
	return nums[length-1]*nums[length-2] - nums[0]*nums[1]
}
