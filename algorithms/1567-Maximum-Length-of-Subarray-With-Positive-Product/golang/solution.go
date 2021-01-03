package golang

func getMaxLen(nums []int) int {
	prevZeroIdx := -1
	longest, length := 0, len(nums)
	negNumIdxRecord := make([]int, 0)

	for i := 0; i < length; i++ {
		if nums[i] == 0 {
			longest = max(
				longest,
				calcMaxLength(negNumIdxRecord, prevZeroIdx, i),
			)

			prevZeroIdx = i
			negNumIdxRecord = negNumIdxRecord[0:0]
		} else if nums[i] < 0 {
			negNumIdxRecord = append(negNumIdxRecord, i)
		}
	}

	longest = max(
		longest,
		calcMaxLength(negNumIdxRecord, prevZeroIdx, length),
	)

	return longest

}

func calcMaxLength(negNumIdxRecord []int, prevZeroIdx, i int) int {
	recordLength := len(negNumIdxRecord)

	if recordLength%2 == 0 {
		return i - prevZeroIdx - 1
	} else {
		first, last :=
			negNumIdxRecord[0], negNumIdxRecord[recordLength-1]
		return max(last-prevZeroIdx-1, i-first-1)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
