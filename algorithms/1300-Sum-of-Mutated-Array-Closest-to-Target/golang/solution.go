package golang

import (
	"sort"
)

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	length := len(arr)

	idx, prevSum := getIdxPrevSum(target, length, arr)

	if idx == length {
		return arr[length-1]
	} else if idx == 0 {
		mod := target % length
		quotient := target / length

		if mod <= (length / 2) {
			return quotient
		} else {
			return quotient + 1
		}
	} else {
		begin, end := arr[idx-1], arr[idx]
		for begin < end {
			mid := (end-begin)/2 + begin
			midSum := getChangedSum(idx-1, length, mid, prevSum)

			if midSum >= target {
				end = mid
			} else {
				begin = mid + 1
			}
		}

		x, y :=
			getChangedSum(idx-1, length, end-1, prevSum),
			getChangedSum(idx-1, length, end, prevSum)

		return helper(x, y, target, end)
	}
}

func helper(x int, y int, target int, end int) int {
	if (target - x) <= (y - target) {
		return end - 1
	} else {
		return end
	}
}

func getChangedSum(idx int, length int, changedVal int, idxSum int) int {
	return idxSum + changedVal*(length-idx-1)
}

func getIdxPrevSum(target int, length int, arr []int) (int, int) {
	prevSum, idxSum := 0, 0
	for idx := 0; idx < length; idx++ {
		val := arr[idx]
		idxSum += val

		if idx != 0 {
			prevSum += arr[idx-1]
		}

		if getChangedSum(idx, length, val, idxSum) >= target {
			return idx, prevSum
		}
	}

	return length, 0
}
