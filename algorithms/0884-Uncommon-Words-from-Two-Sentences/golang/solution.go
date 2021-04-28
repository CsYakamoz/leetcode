package golang

import (
	"strings"
)

func uncommonFromSentences(A string, B string) []string {
	aMap, bMap := generateMap(&A), generateMap(&B)

	ans := []string{}
	findUncommonWords(&aMap, &bMap, &ans)
	findUncommonWords(&bMap, &aMap, &ans)

	return ans
}

func generateMap(sentences *string) map[string]int {
	dict := map[string]int{}
	for _, word := range strings.Split(*sentences, " ") {
		dict[word]++
	}

	return dict
}

func findUncommonWords(aMap, bMap *map[string]int, ans *[]string) {
	for word, count := range *aMap {
		if count != 1 {
			continue
		}

		_, exist := (*bMap)[word]
		if !exist {
			*ans = append(*ans, word)
		}
	}
}
