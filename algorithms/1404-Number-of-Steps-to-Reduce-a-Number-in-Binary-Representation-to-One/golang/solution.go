package golang

func numSteps(s string) int {
	ans, carry := 0, 0

	length := len(s)
	for i := length - 1; i > 0; i-- {
		ch := s[i]

		if carry == 0 {
			if ch == '0' {
				ans += 1
			} else {
				ans += 2
				carry = 1
			}
		} else {
			if ch == '0' {
				ans += 2
			} else {
				ans += 1
			}
		}
	}

	if carry == 1 {
		ans += 1
	}

	return ans
}
