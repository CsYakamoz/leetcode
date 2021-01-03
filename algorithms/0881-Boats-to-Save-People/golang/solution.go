package golang

import (
	"sort"
)

func findFirstBiggerIdx(people []int, limit int) int {
	begin, end := 0, len(people)

	for begin < end {
		mid := (end-begin)/2 + begin
		if people[mid] > limit {
			end = mid
		} else {
			begin = mid + 1
		}
	}

	return end
}

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	firstBiggerIdx := findFirstBiggerIdx(people, limit)
	boat := len(people) - firstBiggerIdx

	left, right := 0, firstBiggerIdx-1
	for left <= right {
		sum := people[left] + people[right]
		if sum > limit {
			right--
		} else {
			left++
			right--
		}
		boat++
	}

	return boat
}
