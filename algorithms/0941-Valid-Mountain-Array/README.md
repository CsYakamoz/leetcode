## [941. Valid Mountain Array](https://leetcode.com/problems/valid-mountain-array/)

### Description

<p>Given an array of integers <code>arr</code>, return <em><code>true</code> if and only if it is a valid mountain array</em>.</p>

<p>Recall that arr is a mountain array if and only if:</p>

<ul>
	<li><code>arr.length &gt;= 3</code></li>
	<li>There exists some <code>i</code> with <code>0 &lt; i &lt; arr.length - 1</code> such that:
	<ul>
		<li><code>arr[0] &lt; arr[1] &lt; ... &lt; arr[i - 1] &lt; arr[i] </code></li>
		<li><code>arr[i] &gt; arr[i + 1] &gt; ... &gt; arr[arr.length - 1]</code></li>
	</ul>
	</li>
</ul>
<img src="https://assets.leetcode.com/uploads/2019/10/20/hint_valid_mountain_array.png" width="500" />
<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<pre><strong>Input:</strong> arr = [2,1]
<strong>Output:</strong> false
</pre><p><strong>Example 2:</strong></p>
<pre><strong>Input:</strong> arr = [3,5,5]
<strong>Output:</strong> false
</pre><p><strong>Example 3:</strong></p>
<pre><strong>Input:</strong> arr = [0,3,2,1]
<strong>Output:</strong> true
</pre>
<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= arr.length &lt;= 10<sup>4</sup></code></li>
	<li><code>0 &lt;= arr[i] &lt;= 10<sup>4</sup></code></li>
</ul>

**Difficulty:** `Easy`

**Tags:** `Array`

### Solution One

```go
func validMountainArray(arr []int) bool {
	length := len(arr)

	// required: arr.length >= 3, and strictly increasing at first
	if length < 3 || arr[0] >= arr[1] {
		return false
	}

	dir := true // true -> inc, false -> dec
	for i := 2; i < length; i++ {
		// required: strictly increasing or decreasing
		if arr[i] == arr[i-1] {
			return false
		}

		if arr[i] > arr[i-1] {
			if !dir {
				return false
			}
		} else {
			if dir {
				dir = false
			}
		}
	}

	return !dir
}
```

### Solution Two - In Top Solutions

```go
func validMountainArray(arr []int) bool {
 if arr == nil || len(arr) < 3 {
		return false
	}
	N := len(arr)
	i := 0

	//walking up
	for i+1 < N && arr[i] < arr[i+1] {
		i++
	}

	//peak can't be first or last
	if i == 0 || i == N-1 {
		return false
	}

	//walk down
	for i+1 < N && arr[i] > arr[i+1] {
		i++
	}

	return i == N-1
}
```
