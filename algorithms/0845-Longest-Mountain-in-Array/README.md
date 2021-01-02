## [845. Longest Mountain in Array](https://leetcode.com/problems/longest-mountain-in-array/)

### Description

<p>You may recall that an array <code>arr</code> is a <strong>mountain array</strong> if and only if:</p>

<ul>
	<li><code>arr.length &gt;= 3</code></li>
	<li>There exists some index <code>i</code> (<strong>0-indexed</strong>) with <code>0 &lt; i &lt; arr.length - 1</code> such that:
	<ul>
		<li><code>arr[0] &lt; arr[1] &lt; ... &lt; arr[i - 1] &lt; arr[i]</code></li>
		<li><code>arr[i] &gt; arr[i + 1] &gt; ... &gt; arr[arr.length - 1]</code></li>
	</ul>
	</li>
</ul>

<p>Given an integer array <code>arr</code>,&nbsp;return <em>the length of the longest subarray, which is a mountain</em>.&nbsp;Return <code>0</code> if there is no mountain subarray.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> arr = [2,1,4,7,3,2,5]
<strong>Output:</strong> 5
<strong>Explanation:</strong> The largest mountain is [1,4,7,3,2] which has length 5.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> arr = [2,2,2]
<strong>Output:</strong> 0
<strong>Explanation:</strong> There is no mountain.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1&nbsp;&lt;= arr.length &lt;= 10<sup>4</sup></code></li>
	<li><code>0 &lt;= arr[i] &lt;= 10<sup>4</sup></code></li>
</ul>

<p>&nbsp;</p>
<p><strong>Follow up:</strong></p>

<ul>
	<li>Can you solve it using only one pass?</li>
	<li>Can you solve it in <code>O(1)</code> space?</li>
</ul>

**Difficulty:** `Medium`

**Tags:** `Two Pointers`

### Solution One

```golang
func findLeftBottom(arr []int, length, begin int) int {
	for i := begin + 1; i < length; i++ {
		if arr[i] > arr[i-1] {
			return i - 1
		}
	}

	return length
}

func findTop(arr []int, length, begin int) int {
	for i := begin + 1; i < length; i++ {
		if arr[i] <= arr[i-1] {
			return i - 1
		}
	}

	return length
}

func findRightBottom(arr []int, length, begin int) int {
	for i := begin + 1; i < length; i++ {
		if arr[i] >= arr[i-1] {
			return i - 1
		}
	}

	return length - 1
}

func longestMountain(arr []int) int {
	length := len(arr)

	// subarray.length should be >= 3
	if length < 3 {
		return 0
	}

	largestMountain, point := 0, 0
	for point < length {
		left := findLeftBottom(arr, length, point)
		if left >= length-2 {
			return largestMountain
		}

		top := findTop(arr, length, left+1)
		if top >= length-1 {
			return largestMountain
		}

		right := findRightBottom(arr, length, top)
		if right > top {
			subarrayLength := right - left + 1
			if subarrayLength > largestMountain {
				largestMountain = subarrayLength
			}
		}

		point = right
	}

	return largestMountain
}
```

### Other Solutions

[845. Longest Mountain in Array - Solution](https://leetcode.com/problems/longest-mountain-in-array/solution/)
