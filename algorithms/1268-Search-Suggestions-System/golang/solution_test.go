package golang

import (
	"reflect"
	"testing"
)

var tests = []struct {
	products   []string
	searchWord string
	output     [][]string
}{
	{
		products:   []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
		searchWord: "mouse",
		output: [][]string{
			{"mobile", "moneypot", "monitor"}, // m
			{"mobile", "moneypot", "monitor"}, // mo
			{"mouse", "mousepad"},             // mou
			{"mouse", "mousepad"},             // mous
			{"mouse", "mousepad"},             // mouse
		},
	},
	{
		products:   []string{"havana"},
		searchWord: "havana",
		output: [][]string{
			{"havana"}, // h
			{"havana"}, // ha
			{"havana"}, // hav
			{"havana"}, // hava
			{"havana"}, // havan
			{"havana"}, // havana
		},
	},
}

func TestSuggestedProducts(t *testing.T) {
	for idx, test := range tests {
		copyData := make([]string, len(test.products), len(test.products))
		copy(copyData, test.products)

		actual := suggestedProducts(copyData, test.searchWord)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf(
				"TestCase[%d]: products = %v searchWord = %s, expected %v but get %v",
				idx,
				test.products,
				test.searchWord,
				test.output,
				actual,
			)
		}
	}
}

