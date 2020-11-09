package golang

var mark int = 1000000

// -1000 ≤ nums[i] ≤ 1000 && nums[i] ≠ 0
func isInCorrectRange(num int) bool {
	return -1000 <= num && num <= 1000
}

func getValAfterMark(num int) int {
	return num + mark
}

func getOriginalVal(num int) int {
	if isInCorrectRange(num) {
		return num
	}

	return num - mark
}

func getNextIdx(movement int, idx int, length int) int {
	trulyMov := movement % length
	return (trulyMov + idx + length) % length
}

func circularArrayLoop(nums []int) bool {
	length := len(nums)

	for i := 0; i < length; i++ {
		if !isInCorrectRange(nums[i]) {
			continue
		}

		direction, idx := nums[i] > 0, i

		for isInCorrectRange(nums[idx]) {
			movement := nums[idx]
			if (movement > 0) != direction {
				break
			}

			nums[idx] = getValAfterMark(movement)
			idx = getNextIdx(movement, idx, length)
		}

		curr, loopLength := idx, 0
		for {
			movement := getOriginalVal(nums[curr])
			if movement > 0 != direction || movement%length == 0 {
				break
			}
			curr = getNextIdx(movement, curr, length)
			loopLength++

			if curr == idx {
				break
			}
		}

		if curr == idx && loopLength > 1 {
			return true
		}
	}

	return false
}
