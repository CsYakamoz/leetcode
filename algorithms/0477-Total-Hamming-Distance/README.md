## [477. Total Hamming Distance](https://leetcode.com/problems/total-hamming-distance/)

### Description

<p>The <a href="https://en.wikipedia.org/wiki/Hamming_distance" target="_blank">Hamming distance</a> between two integers is the number of positions at which the corresponding bits are different.</p>

<p>Now your job is to find the total Hamming distance between all pairs of the given numbers.</p>

<p><b>Example:</b><br />
<pre>
<b>Input:</b> 4, 14, 2

<b>Output:</b> 6

<b>Explanation:</b> In binary representation, the 4 is 0100, 14 is 1110, and 2 is 0010 (just
showing the four bits relevant in this case). So the answer will be:
HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2 + 2 + 2 = 6.

</pre>
</p>

<p><b>Note:</b><br>
<ol>
<li>Elements of the given array are in the range of <code>0 </code> to <code>10^9</code>
<li>Length of the array will not exceed <code>10^4</code>. </li>
</ol>
</p>

**Difficulty:** `Medium`

**Tags:** `Bit Manipulation`

### Solution One

```go
func totalHammingDistance(nums []int) int {
	ans, length, dict := 0, len(nums), map[int]int{}

	for _, val := range nums {
		str := strconv.FormatInt(int64(val), 2)
		length := len(str)

		for i := length - 1; i >= 0; i-- {
			if str[i] == '0' {
				continue
			}
			dict[length-i]++
		}
	}

	for _, count := range dict {
		ans += count * (length - count)
	}

	return ans
}
```

### Solution Two - In Top Solutions

```go
func totalHammingDistance(nums []int) int {
	r := 0
	n := len(nums)
	for i := 0; i < 32; i++ {
		on := 0
		for _, v := range nums {
			if v&(1<<i) != 0 {
				on++
			}
		}
		r += on * (n - on)
	}
	return r
}
```
