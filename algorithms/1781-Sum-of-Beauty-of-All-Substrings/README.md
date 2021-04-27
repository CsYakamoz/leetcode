## [1781. Sum of Beauty of All Substrings](https://leetcode.com/problems/sum-of-beauty-of-all-substrings/)

### Description

<p>The <strong>beauty</strong> of a string is the difference in frequencies between the most frequent and least frequent characters.</p>

<ul>
	<li>For example, the beauty of <code>&quot;abaacc&quot;</code> is <code>3 - 1 = 2</code>.</li>
</ul>

<p>Given a string <code>s</code>, return <em>the sum of <strong>beauty</strong> of all of its substrings.</em></p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;aabcb&quot;
<strong>Output:</strong> 5
<strong>Explanation: </strong>The substrings with non-zero beauty are [&quot;aab&quot;,&quot;aabc&quot;,&quot;aabcb&quot;,&quot;abcb&quot;,&quot;bcb&quot;], each with beauty equal to 1.</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;aabcbaa&quot;
<strong>Output:</strong> 17
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;=<sup> </sup>500</code></li>
	<li><code>s</code> consists of only lowercase English letters.</li>
</ul>

**Difficulty:** `Medium`

**Tags:** `Hash Table` `String`

### Solution One

```go
func beautySum(s string) int {
	var dict [26]int

	ans := 0
	prefix := [][26]int{}

	base := int('a')
	for _, c := range s {
		idx := int(c) - base
		dict[idx]++

		ans += calcBeauty(&dict)

		prefix = append(prefix, dict)
	}

	length := len(s)
	for i := 2; i < length; i++ {
		for j := 0; j < length-i; j++ {
			min, max := math.MaxInt32, math.MinInt32

			for k := 0; k < 26; k++ {
				count := prefix[j+i][k] - prefix[j][k]
				if count == 0 {
					continue
				}

				if count < min {
					min = count
				}

				if count > max {
					max = count
				}
			}

			ans += max - min
		}
	}

	return ans
}

func calcBeauty(dict *[26]int) int {
	min, max := math.MaxInt32, math.MinInt32

	for _, val := range *dict {
		if val == 0 {
			continue
		}

		if val < min {
			min = val
		}

		if val > max {
			max = val
		}
	}

	return max - min
}
```
