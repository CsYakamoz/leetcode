## [1818. Minimum Absolute Sum Difference](https://leetcode.com/problems/minimum-absolute-sum-difference/)

### Description

<p>You are given two positive integer arrays <code>nums1</code> and <code>nums2</code>, both of length <code>n</code>.</p>

<p>The <strong>absolute sum difference</strong> of arrays <code>nums1</code> and <code>nums2</code> is defined as the <strong>sum</strong> of <code>|nums1[i] - nums2[i]|</code> for each <code>0 &lt;= i &lt; n</code> (<strong>0-indexed</strong>).</p>

<p>You can replace <strong>at most one</strong> element of <code>nums1</code> with <strong>any</strong> other element in <code>nums1</code> to <strong>minimize</strong> the absolute sum difference.</p>

<p>Return the <em>minimum absolute sum difference <strong>after</strong> replacing at most one<strong> </strong>element in the array <code>nums1</code>.</em> Since the answer may be large, return it <strong>modulo</strong> <code>10<sup>9</sup> + 7</code>.</p>

<p><code>|x|</code> is defined as:</p>

<ul>
	<li><code>x</code> if <code>x &gt;= 0</code>, or</li>
	<li><code>-x</code> if <code>x &lt; 0</code>.</li>
</ul>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums1 = [1,7,5], nums2 = [2,3,5]
<strong>Output:</strong> 3
<strong>Explanation: </strong>There are two possible optimal solutions:
- Replace the second element with the first: [1,<u><strong>7</strong></u>,5] =&gt; [1,<u><strong>1</strong></u>,5], or
- Replace the second element with the third: [1,<u><strong>7</strong></u>,5] =&gt; [1,<u><strong>5</strong></u>,5].
Both will yield an absolute sum difference of <code>|1-2| + (|1-3| or |5-3|) + |5-5| = </code>3.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums1 = [2,4,6,8,10], nums2 = [2,4,6,8,10]
<strong>Output:</strong> 0
<strong>Explanation: </strong>nums1 is equal to nums2 so no replacement is needed. This will result in an
absolute sum difference of 0.
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums1 = [1,10,4,4,2,7], nums2 = [9,3,5,1,7,4]
<strong>Output:</strong> 20
<strong>Explanation: </strong>Replace the first element with the second: [<u><strong>1</strong></u>,10,4,4,2,7] =&gt; [<u><strong>10</strong></u>,10,4,4,2,7].
This yields an absolute sum difference of <code>|10-9| + |10-3| + |4-5| + |4-1| + |2-7| + |7-4| = 20</code>
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>n == nums1.length</code></li>
	<li><code>n == nums2.length</code></li>
	<li><code>1 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums1[i], nums2[i] &lt;= 10<sup>5</sup></code></li>
</ul>

**Difficulty:** `Medium`

**Tags:** `Array` `Binary Search` `Sorting` `Ordered Set`

### Solution One

```go
type NumAndAbs struct {
	num int
	abs int
}

const mod = 1000000000 + 7

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	sli := make([]NumAndAbs, 0)
	absSumDiff, length := 0, len(nums1)

	for idx := 0; idx < length; idx++ {
		abs := getAbs(nums1[idx] - nums2[idx])

		absSumDiff = (absSumDiff + abs) % mod
		if abs != 0 {
			sli = append(sli, NumAndAbs{nums2[idx], abs})
		}
	}

	// length is always >= 1
	if absSumDiff == 0 || length == 1 {
		return absSumDiff
	}

	sort.Ints(nums1)
	sort.Slice(sli, func(i, j int) bool { return sli[i].abs > sli[j].abs })

	diff, sliLen := 0, len(sli)

	for i := 0; i < sliLen && diff < sli[i].abs; i++ {
		v := sli[i]
		idx := binarySearch(nums1, v.num)
		abs := getAbsByIdx(nums1, v.num, idx)

		if v.abs-abs > diff {
			diff = v.abs - abs
		}
	}

	return (absSumDiff - diff + mod) % mod
}

func getAbsByIdx(nums []int, num, idx int) int {
	if idx == len(nums) {
		return getAbs(num - nums[idx-1])
	} else if idx == 0 {
		return getAbs(num - nums[idx])
	} else {
		frontAbs, backAbs := getAbs(num-nums[idx-1]), getAbs(num-nums[idx])
		if frontAbs < backAbs {
			return frontAbs
		} else {
			return backAbs
		}
	}
}

func getAbs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func binarySearch(nums []int, target int) int {
	begin, end := 0, len(nums)

	for begin < end {
		mid := (end-begin)/2 + begin

		if nums[mid] >= target {
			end = mid
		} else {
			begin = mid + 1
		}
	}

	return end
}
```

### Solution Two - In Top Solutions

```go
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
    diff := 0
	maxDiff := 0
	for i, n2 := range nums2 {
		d := abs(nums1[i] - n2)
		diff += d
		if maxDiff < d {
			for _, n1 := range nums1 {
				maxDiff = max(maxDiff, d-abs(n1-n2))
			}
		}
	}
	return (diff - maxDiff) % (1e9 + 7)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```

### Solution Three - In Top Solutions

[DO NOT USE O(n) find the max Diff pair, USE binary search](<https://leetcode.com/problems/minimum-absolute-sum-difference/discuss/1141490/DO-NOT-USE-O(n)-find-the-max-Diff-pair-USE-binary-search>)
