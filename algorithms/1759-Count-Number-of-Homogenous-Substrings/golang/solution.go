package golang

func countHomogenous(s string) int {
	ans, modulo, length := 0, int(1e9+7), len(s)

	start := 0
	for i := 1; i < length; i++ {
		if s[i] == s[i-1] {
			continue
		}

		ans += getSn(i - start)
		ans %= modulo
		start = i
	}

	ans += getSn(length - start)
	ans %= modulo

	return ans
}

func getSn(n int) int {
	return (1 + n) * n / 2
}
