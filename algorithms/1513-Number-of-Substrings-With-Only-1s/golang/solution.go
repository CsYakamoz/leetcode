package golang

func numSub(s string) int {
	ans, modulo, length := 0, int(1e9+7), len(s)

	start := -1
	for idx, ch := range s {
		if ch == '0' {
			if start != -1 {
				ans += getSn(idx - start)
				ans %= modulo
				start = -1
			}
		} else {
			if start == -1 {
				start = idx
			}
		}

	}

	if start != -1 {
		ans += getSn(length - start)
		ans %= modulo
	}

	return ans
}

func getSn(n int) int {
	return (1 + n) * n / 2
}
