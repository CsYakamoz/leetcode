## [884. Uncommon Words from Two Sentences](https://leetcode.com/problems/uncommon-words-from-two-sentences/)

### Description

<p>We are given two sentences <code>A</code> and <code>B</code>.&nbsp; (A <em>sentence</em>&nbsp;is a string of space separated words.&nbsp; Each <em>word</em> consists only of lowercase letters.)</p>

<p>A word is <em>uncommon</em>&nbsp;if it appears exactly once in one of the sentences, and does not appear in the other sentence.</p>

<p>Return a list of all uncommon words.&nbsp;</p>

<p>You may return the list in any order.</p>

<p>&nbsp;</p>

<ol>
</ol>

<div>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input: </strong>A = <span id="example-input-1-1">&quot;this apple is sweet&quot;</span>, B = <span id="example-input-1-2">&quot;this apple is sour&quot;</span>
<strong>Output: </strong><span id="example-output-1">[&quot;sweet&quot;,&quot;sour&quot;]</span>
</pre>

<div>
<p><strong>Example 2:</strong></p>

<pre>
<strong>Input: </strong>A = <span id="example-input-2-1">&quot;apple apple&quot;</span>, B = <span id="example-input-2-2">&quot;banana&quot;</span>
<strong>Output: </strong><span id="example-output-2">[&quot;banana&quot;]</span>
</pre>

<p>&nbsp;</p>

<p><strong>Note:</strong></p>

<ol>
	<li><code>0 &lt;= A.length &lt;= 200</code></li>
	<li><code>0 &lt;= B.length &lt;= 200</code></li>
	<li><code>A</code> and <code>B</code> both contain only spaces and lowercase letters.</li>
</ol>
</div>
</div>

**Difficulty:** `Easy`

**Tags:** `Hash Table`

### Solution One

```go
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
```

### Solution Two - In Top Solutions

[[C++/Java/Python] Easy Solution with Explanation](https://leetcode.com/problems/uncommon-words-from-two-sentences/discuss/158967/C%2B%2BJavaPython-Easy-Solution-with-Explanation)

```cpp
    vector<string> uncommonFromSentences(string A, string B) {
        unordered_map<string, int> count;
        istringstream iss(A + " " + B);
        while (iss >> A) count[A]++;
        vector<string> res;
        for (auto w: count)
            if (w.second == 1)
                res.push_back(w.first);
        return res;
    }
```

### Other Solutions

[884. Uncommon Words from Two Sentences](https://leetcode.com/problems/uncommon-words-from-two-sentences/solution/)
