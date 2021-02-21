package golang

func integerReplacement(n int) int {
	operation, intervalList := helper(n)

	length := len(intervalList)
	for i := 0; i < length; i++ {
		if i == length-1 {
			count := intervalList[i][1] - intervalList[i][0]

			if count == 1 {
				break
			} else if count == 2 {
				operation += 2
			} else {
				operation += count + 1
			}
		} else {
			if intervalList[i+1][0]-intervalList[i][1] == 1 {
				count := intervalList[i][1] - intervalList[i][0]
				if count == 1 {
					operation += 2
				} else {
					operation += intervalList[i][1] - intervalList[i][0] + 1
					intervalList[i+1][0]--
					operation--
				}
			} else {
				count := intervalList[i][1] - intervalList[i][0]
				if count > 2 {
					operation += count + 2
				} else {
					operation += count * 2
				}
			}
		}
	}

	return operation
}

func helper(n int) (int, [][2]int) {
	operation, intervalList := 0, make([][2]int, 0)

	i, start := 0, -1
	for curr := n; curr != 0 && i < 31; i++ {
		if n&(1<<i) != 0 {
			if start == -1 {
				start = i
			}
		} else {
			if start != -1 {
				intervalList = append(intervalList, [2]int{start, i})
				start = -1
			}
			operation++
		}

		curr >>= 1
	}

	if start != -1 {
		intervalList = append(intervalList, [2]int{start, i})
	}

	return operation, intervalList
}
