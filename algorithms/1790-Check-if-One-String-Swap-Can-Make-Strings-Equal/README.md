## [1790. Check if One String Swap Can Make Strings Equal](https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/)

### Description

<p>You are given two strings <code>s1</code> and <code>s2</code> of equal length. A <strong>string swap</strong> is an operation where you choose two indices in a string (not necessarily different) and swap the characters at these indices.</p>

<p>Return <code>true</code> <em>if it is possible to make both strings equal by performing <strong>at most one string swap </strong>on <strong>exactly one</strong> of the strings. </em>Otherwise, return <code>false</code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> s1 = &quot;bank&quot;, s2 = &quot;kanb&quot;
<strong>Output:</strong> true
<strong>Explanation:</strong> For example, swap the first character with the last character of s2 to make &quot;bank&quot;.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> s1 = &quot;attack&quot;, s2 = &quot;defend&quot;
<strong>Output:</strong> false
<strong>Explanation:</strong> It is impossible to make them equal with one string swap.
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> s1 = &quot;kelb&quot;, s2 = &quot;kelb&quot;
<strong>Output:</strong> true
<strong>Explanation:</strong> The two strings are already equal, so no string swap operation is required.
</pre>

<p><strong>Example 4:</strong></p>

<pre>
<strong>Input:</strong> s1 = &quot;abcd&quot;, s2 = &quot;dcba&quot;
<strong>Output:</strong> false
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s1.length, s2.length &lt;= 100</code></li>
	<li><code>s1.length == s2.length</code></li>
	<li><code>s1</code> and <code>s2</code> consist of only lowercase English letters.</li>
</ul>

**Difficulty:** `Easy`

**Tags:** `String`

### Solution One

看了 `Hint` 才想出来的 😑

看了 `Discuss` 后, 发现 `for` 循环中的 `i < length` 可以改成 `i < length && len(diffIdxList) < 3`

```go
func areAlmostEqual(s1 string, s2 string) bool {
	diffIdxList, length := []int{}, len(s1)

	for i := 0; i < length; i++ {
		if s1[i] == s2[i] {
			continue
		}

		diffIdxList = append(diffIdxList, i)
	}

	diffLength := len(diffIdxList)
	if diffLength != 0 && diffLength != 2 {
		return false
	}

	if diffLength == 0 {
		return true
	}

	previous, latter := diffIdxList[0], diffIdxList[1]
	return s1[previous] == s2[latter] && s1[latter] == s2[previous]
}
```
