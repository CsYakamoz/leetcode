package golang

func areAlmostEqual(s1 string, s2 string) bool {
	diffIdxList, length := []int{}, len(s1)

	for i := 0; i < length; i++ {
		if s1[i] == s2[i] {
			continue
		}

		diffIdxList = append(diffIdxList, i)
	}

	diffLength := len(diffIdxList)
	if diffLength != 0 && diffLength != 2 {
		return false
	}

	if diffLength == 0 {
		return true
	}

	previous, latter := diffIdxList[0], diffIdxList[1]
	return s1[previous] == s2[latter] && s1[latter] == s2[previous]
}
