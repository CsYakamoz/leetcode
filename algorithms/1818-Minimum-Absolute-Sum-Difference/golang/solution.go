package golang

import (
	"sort"
)

type NumAndAbs struct {
	num int
	abs int
}

const mod = 1000000000 + 7

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	sli := make([]NumAndAbs, 0)
	absSumDiff, length := 0, len(nums1)

	for idx := 0; idx < length; idx++ {
		abs := getAbs(nums1[idx] - nums2[idx])

		absSumDiff = (absSumDiff + abs) % mod
		if abs != 0 {
			sli = append(sli, NumAndAbs{nums2[idx], abs})
		}
	}

	// length is always >= 1
	if absSumDiff == 0 || length == 1 {
		return absSumDiff
	}

	sort.Ints(nums1)
	sort.Slice(sli, func(i, j int) bool { return sli[i].abs > sli[j].abs })

	diff, sliLen := 0, len(sli)

	for i := 0; i < sliLen && diff < sli[i].abs; i++ {
		v := sli[i]
		idx := binarySearch(nums1, v.num)
		abs := getAbsByIdx(nums1, v.num, idx)

		if v.abs-abs > diff {
			diff = v.abs - abs
		}
	}

	return (absSumDiff - diff + mod) % mod
}

func getAbsByIdx(nums []int, num, idx int) int {
	if idx == len(nums) {
		return getAbs(num - nums[idx-1])
	} else if idx == 0 {
		return getAbs(num - nums[idx])
	} else {
		frontAbs, backAbs := getAbs(num-nums[idx-1]), getAbs(num-nums[idx])
		if frontAbs < backAbs {
			return frontAbs
		} else {
			return backAbs
		}
	}
}

func getAbs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func binarySearch(nums []int, target int) int {
	begin, end := 0, len(nums)

	for begin < end {
		mid := (end-begin)/2 + begin

		if nums[mid] >= target {
			end = mid
		} else {
			begin = mid + 1
		}
	}

	return end
}
