package golang

func mergeAlternately(word1 string, word2 string) string {
	ans, i, j := "", 0, 0
	word1Length, word2Length := len(word1), len(word2)

	for i < word1Length && j < word2Length {
		ans += word1[i:i+1] + word2[j:j+1]

		i++
		j++
	}

	if i < word1Length {
		ans += word1[i:]
	}

	if j < word2Length {
		ans += word2[j:]
	}

	return ans
}
