package golang

func sortString(s string) string {
	length := len(s)
	dict := [26]int{}

	base := int('a')
	for _, c := range s {
		dict[int(c)-base]++
	}

	ans := ""
	for length > 0 {
		for i := 0; length > 0 && i < 26; i++ {
			if dict[i] > 0 {
				ans += string(rune('a' + i))
				dict[i]--
				length--
			}
		}

		for i := 25; length > 0 && i >= 0; i-- {
			if dict[i] > 0 {
				ans += string(rune('a' + i))
				dict[i]--
				length--
			}
		}
	}

	return ans
}
