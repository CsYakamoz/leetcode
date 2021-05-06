## [1838. Frequency of the Most Frequent Element](https://leetcode.com/problems/frequency-of-the-most-frequent-element/)

### Description

<p>The <strong>frequency</strong> of an element is the number of times it occurs in an array.</p>

<p>You are given an integer array <code>nums</code> and an integer <code>k</code>. In one operation, you can choose an index of <code>nums</code> and increment the element at that index by <code>1</code>.</p>

<p>Return <em>the <strong>maximum possible frequency</strong> of an element after performing <strong>at most</strong> </em><code>k</code><em> operations</em>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,2,4], k = 5
<strong>Output:</strong> 3<strong>
Explanation:</strong> Increment the first element three times and the second element two times to make nums = [4,4,4].
4 has a frequency of 3.</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,4,8,13], k = 5
<strong>Output:</strong> 2
<strong>Explanation:</strong> There are multiple optimal solutions:
- Increment the first element three times to make nums = [4,4,8,13]. 4 has a frequency of 2.
- Increment the second element four times to make nums = [1,8,8,13]. 8 has a frequency of 2.
- Increment the third element five times to make nums = [1,4,13,13]. 13 has a frequency of 2.
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums = [3,9,6], k = 2
<strong>Output:</strong> 1
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= k &lt;= 10<sup>5</sup></code></li>
</ul>

**Difficulty:** `Medium`

**Tags:** `Greedy`

### Solution One

```go
type maxFreqUnit struct {
	num   int
	count int
}

func maxFrequency(nums []int, k int) int {
	dict := map[int]int{}
	for _, num := range nums {
		dict[num]++
	}

	list := make([]maxFreqUnit, 0, len(dict))
	for k, v := range dict {
		list = append(list, maxFreqUnit{
			num:   k,
			count: v,
		})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].num < list[j].num
	})

	ans := 0
	for i := 0; i < len(list); i++ {
		remainedK, maxFreq := k, list[i].count

		for j := i - 1; j >= 0; j-- {
			diff := list[i].num - list[j].num

			if remainedK < diff {
				break
			} else {
				if remainedK >= list[j].count*diff {
					maxFreq += list[j].count
					remainedK -= list[j].count * diff
				} else {
					maxFreq += remainedK / diff
					break
				}
			}
		}

		ans = max(ans, maxFreq)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
```

### Solution Two - In Top Solutions

[[Java/C++/Python] Sliding Window](https://leetcode.com/problems/frequency-of-the-most-frequent-element/discuss/1175090/JavaC%2B%2BPython-Sliding-Window)

```cpp
    int maxFrequency(vector<int>& A, long k) {
        int i = 0, j;
        sort(A.begin(), A.end());
        for (j = 0; j < A.size(); ++j) {
            k += A[j];
            if (k < (long)A[j] * (j - i + 1))
                k -= A[i++];
        }
        return j - i;
    }
```
