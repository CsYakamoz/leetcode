package golang

func checkIfPangram(sentence string) bool {
	dict := [26]int{}

	for _, c := range sentence {
		idx := int(c) - 97

		dict[idx]++
	}

	for _, count := range dict {
		if count == 0 {
			return false
		}
	}

	return true
}
