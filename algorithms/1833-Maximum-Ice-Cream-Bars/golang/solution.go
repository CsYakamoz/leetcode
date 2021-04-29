package golang

import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)

	ans := 0

	for _, cost := range costs {
		if coins >= cost {
			ans++
			coins -= cost
		} else {
			break
		}
	}

	return ans
}
