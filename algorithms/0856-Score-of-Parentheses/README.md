## [856. Score of Parentheses](https://leetcode.com/problems/score-of-parentheses/)

### Description

<p>Given a balanced parentheses string <code>S</code>, compute the score of the string based on the following rule:</p>

<ul>
	<li><code>()</code> has score 1</li>
	<li><code>AB</code> has score <code>A + B</code>, where A and B are balanced parentheses strings.</li>
	<li><code>(A)</code> has score <code>2 * A</code>, where A is a balanced parentheses string.</li>
</ul>

<p>&nbsp;</p>

<div>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input: </strong><span id="example-input-1-1">&quot;()&quot;</span>
<strong>Output: </strong><span id="example-output-1">1</span>
</pre>

<div>
<p><strong>Example 2:</strong></p>

<pre>
<strong>Input: </strong><span id="example-input-2-1">&quot;(())&quot;</span>
<strong>Output: </strong><span id="example-output-2">2</span>
</pre>

<div>
<p><strong>Example 3:</strong></p>

<pre>
<strong>Input: </strong><span id="example-input-3-1">&quot;()()&quot;</span>
<strong>Output: </strong><span id="example-output-3">2</span>
</pre>

<div>
<p><strong>Example 4:</strong></p>

<pre>
<strong>Input: </strong><span id="example-input-4-1">&quot;(()(()))&quot;</span>
<strong>Output: </strong><span id="example-output-4">6</span>
</pre>

<p>&nbsp;</p>

<p><strong>Note:</strong></p>

<ol>
	<li><code>S</code> is a balanced parentheses string, containing only <code>(</code> and <code>)</code>.</li>
	<li><code>2 &lt;= S.length &lt;= 50</code></li>
</ol>
</div>
</div>
</div>
</div>

**Difficulty:** `Medium`

**Tags:** `String` `Stack`

### Solution One

```go
func scoreOfParentheses(S string) int {
	scoreStack := []int{}
	unclosedStack := []int{}

	for _, ch := range S {
		if ch == '(' {
			scoreStack = append(scoreStack, 1)
			unclosedStack = append(unclosedStack, len(scoreStack)-1)
			continue
		}

		scoreStackLength, unclosedStackLength :=
			len(scoreStack), len(unclosedStack)

		idx := unclosedStack[unclosedStackLength-1]
		if idx != scoreStackLength-1 {
			sum := 0
			for i := idx + 1; i < scoreStackLength; i++ {
				sum += scoreStack[i]
			}
			scoreStack[idx] = 2 * sum
		}

		unclosedStack = unclosedStack[:unclosedStackLength-1]
		scoreStack = scoreStack[:idx+1]
	}

	ans := 0
	for _, score := range scoreStack {
		ans += score
	}

	return ans
}
```

### Other Solutions

[856. Score of Parentheses](https://leetcode.com/problems/score-of-parentheses/solution/)
