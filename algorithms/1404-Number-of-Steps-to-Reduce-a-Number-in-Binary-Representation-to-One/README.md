## [1404. Number of Steps to Reduce a Number in Binary Representation to One](https://leetcode.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/)

### Description

<p>Given a number&nbsp;<code>s</code> in their binary representation. Return the number of steps to reduce it to 1 under the following rules:</p>

<ul>
	<li>
	<p>If the current number is even, you have to divide it by 2.</p>
	</li>
	<li>
	<p>If the current number is odd, you have to add 1 to it.</p>
	</li>
</ul>

<p>It&#39;s guaranteed that you can always reach to one for all testcases.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;1101&quot;
<strong>Output:</strong> 6
<strong>Explanation:</strong> &quot;1101&quot; corressponds to number 13 in their decimal representation.
Step 1) 13 is odd, add 1 and obtain 14.&nbsp;
Step 2) 14 is even, divide by 2 and obtain 7.
Step 3) 7 is odd, add 1 and obtain 8.
Step 4) 8 is even, divide by 2 and obtain 4.&nbsp;
Step 5) 4 is even, divide by 2 and obtain 2.&nbsp;
Step 6) 2 is even, divide by 2 and obtain 1.&nbsp;
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;10&quot;
<strong>Output:</strong> 1
<strong>Explanation:</strong> &quot;10&quot; corressponds to number 2 in their decimal representation.
Step 1) 2 is even, divide by 2 and obtain 1.&nbsp;
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;1&quot;
<strong>Output:</strong> 0
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length&nbsp;&lt;= 500</code></li>
	<li><code>s</code> consists of characters &#39;0&#39; or &#39;1&#39;</li>
	<li><code>s[0] == &#39;1&#39;</code></li>
</ul>

**Difficulty:** `Medium`

**Tags:** `String` `Bit Manipulation`

### Solution One

```go
func numSteps(s string) int {
	ans, carry := 0, 0

	length := len(s)
	for i := length - 1; i > 0; i-- {
		ch := s[i]

		if carry == 0 {
			if ch == '0' {
				ans += 1
			} else {
				ans += 2
				carry = 1
			}
		} else {
			if ch == '0' {
				ans += 2
			} else {
				ans += 1
			}
		}
	}

	if carry == 1 {
		ans += 1
	}

	return ans
}
```