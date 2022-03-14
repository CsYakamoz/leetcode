package golang

func wonderfulSubstrings(word string) int64 {
	ans, length := int64(0), len(word)
	xor, prefixXor := 0, make([]int, 0, length)

	for _, char := range word {
		xor ^= getNumAfterConvert(char)
		if isWonderful(xor) {
			ans++
		}

		prefixXor = append(prefixXor, xor)
	}

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if isWonderful(prefixXor[j] ^ prefixXor[i]) {
				ans++
			}
		}
	}

	return ans
}

const base = int('a')

func getNumAfterConvert(char rune) int {
	n := int(char) - base

	return 1 << n
}

func isWonderful(n int) bool {
	return n&(n-1) == 0
}
