## [318. Maximum Product of Word Lengths](https://leetcode.com/problems/maximum-product-of-word-lengths/)

### Description

<p>Given a string array <code>words</code>, find the maximum value of <code>length(word[i]) * length(word[j])</code> where the two words do not share common letters. You may assume that each word will contain only lower case letters. If no such two words exist, return 0.</p>

<p><b>Example 1:</b></p>

<pre>
<b>Input:</b> <code>[&quot;abcw&quot;,&quot;baz&quot;,&quot;foo&quot;,&quot;bar&quot;,&quot;xtfn&quot;,&quot;abcdef&quot;]</code>
<b>Output: </b><code>16
<strong>Explanation: </strong></code>The two words can be <code>&quot;abcw&quot;, &quot;xtfn&quot;</code><span style="font-family: sans-serif, Arial, Verdana, &quot;Trebuchet MS&quot;;">.</span></pre>

<p><b>Example 2:</b></p>

<pre>
<b>Input:</b> <code>[&quot;a&quot;,&quot;ab&quot;,&quot;abc&quot;,&quot;d&quot;,&quot;cd&quot;,&quot;bcd&quot;,&quot;abcd&quot;]</code>
<b>Output: </b><code>4
<strong>Explanation: </strong></code>The two words can be <code>&quot;ab&quot;, &quot;cd&quot;</code><span style="font-family: sans-serif, Arial, Verdana, &quot;Trebuchet MS&quot;;">.</span></pre>

<p><b>Example 3:</b></p>

<pre>
<b>Input:</b> <code>[&quot;a&quot;,&quot;aa&quot;,&quot;aaa&quot;,&quot;aaaa&quot;]</code>
<b>Output: </b><code>0
<strong>Explanation: </strong></code>No such pair of words.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>0 &lt;= words.length &lt;= 10^3</code></li>
	<li><code>0 &lt;= words[i].length &lt;= 10^3</code></li>
	<li><code>words[i]</code> consists only of lowercase English letters.</li>
</ul>

**Difficulty:** `Medium`

**Tags:** `Bit Manipulation`

### Solution One

```go
type unit struct {
	bit        int
	wordLength int
}

func maxProduct(words []string) int {
	ans, length := 0, len(words)
	list := make([]unit, length, length)

	for i, word := range words {
		bit, wordLength := wordToBits(word), len(word)

		for j := 0; j < i; j++ {
			if bit&list[j].bit == 0 {
				ans = max(ans, wordLength*list[j].wordLength)
			}
		}

		list[i].bit = bit
		list[i].wordLength = wordLength
	}

	return ans
}

func wordToBits(word string) int {
	ans := 0

	for _, ch := range word {
		ans |= 1 << (int(ch) - 97)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
```
