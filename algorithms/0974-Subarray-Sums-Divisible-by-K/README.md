## [974. Subarray Sums Divisible by K](https://leetcode.com/problems/subarray-sums-divisible-by-k/)

### Description

<p>Given an array <code>A</code> of integers, return the number of (contiguous, non-empty) subarrays that have a sum divisible by <code>K</code>.</p>

<div>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input: </strong>A = <span id="example-input-1-1">[4,5,0,-2,-3,1]</span>, K = <span id="example-input-1-2">5</span>
<strong>Output: </strong><span id="example-output-1">7</span>
<strong>Explanation: </strong>There are 7 subarrays with a sum divisible by K = 5:
[4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]
</pre>

<p><strong>Note:</strong></p>

<ol>
  <li><code>1 &lt;= A.length &lt;= 30000</code></li>
  <li><code>-10000 &lt;= A[i] &lt;= 10000</code></li>
  <li><code>2 &lt;= K &lt;= 10000</code></li>
</ol>
</div>

**Difficulty:** `Medium`

**Tags:** `Array` `Hash Table`

### Solution One

首先有如下定义：

1. `prefixSum[i]`: `A[0] + ... + A[i]`

2. `remainder[i]`: `prefixSum[i] % K`

思路如下：

对于 **i**, 在 `[0, i)` 范围寻找 **j**, 使得 `(prefixSum[i] - prefixSum[j]) % K === 0`

那么问题则是如何高效地找到 **j**

我们对等式做下转换：

> `(prefixSum[i] - prefixSum[j]) % K = 0` 等式一
>
> 等式一两边乘以 K, 并且将 `prefixSum[j]` 移到等式右边，可得：
>
> `prefixSum[i] = prefixSum[j] + nK (n 为整数）` 等式二
>
> 由于 `prefixSum = remainder + nK (n 为整数）`, 代入等式二，可得：
>
> `remainder[i] + xK = remainder[j] + yK + nK (x, y, n 皆为整数）` 等式三
>
> 可得 `remainder[j] = remainder[i] + (x - y - n)K`
>
> 由于 `remainder[j]` 的取值范围为 (-K, K)
>
> 因此 (x - y -n) 的值只能是 -1, 0, 1

故 j 的条件如下：

1. `remainder[i] === remainder[j]`

2. `若 remainder[i] > 0, 则 remainder[i] === remainder[j] + K`

3. `若 remainder[i] < 0, 则 remainder[i] === remainder[j] - K`

注意：上述情况都是在描述 `prefixSum[i] - prefixSum[j] === A[j+1] + ... + A[i] (j >= 0)`

少了 `A[0] + ... + A[i]` 的情况，即 `remainder[i] === 0`, 此时 `prefixSum[i]` 也满足题目所说的 `subarrays`

```javascript
/**
 * @param {number[]} A
 * @param {number} K
 * @return {number}
 */
const subarraysDivByK = function (A, K) {
  const prefixRemainderDict = new Map();

  let result = 0;
  let remainder = 0;
  for (let i = 0, len = A.length; i < len; i++) {
    remainder = (remainder + A[i]) % K;

    if (remainder === 0) {
      result += 1;
    } else {
      const anotherRemainder =
        remainder - K * (remainder / Math.abs(remainder));
      result += prefixRemainderDict.get(anotherRemainder) || 0;
    }

    const count = prefixRemainderDict.get(remainder) || 0;
    result += count;
    prefixRemainderDict.set(remainder, count + 1);
  }

  return result;
};
```

### Solution Two - In Top Solutions

Link: [DETAILED WHITEBOARD! BEATS 100% (Do you really want to understand It?)](<https://leetcode.com/problems/subarray-sums-divisible-by-k/discuss/413234/DETAILED-WHITEBOARD!-BEATS-100-(Do-you-really-want-to-understand-It)>)

TODO: 补充该链接的思路；目前这份代码是 javascript 中最快的。

```javascript
var subarraysDivByK = function (A, K) {
  let freq = new Array(K).fill(0); // "moduloK : Times I've seen it so far"

  freq[0] = 1; //  Explained below

  // This is the accumulative sum of the elements of A
  let sum = 0;

  // The count of wanted subarrays, whose Sum%K= zero
  let count = 0;

  for (let i = 0; i < A.length; i++) {
    sum = sum + A[i];

    var remainder = sum % K;

    //ALWAYS CHOOSE THE POSITIVE REMAINDER
    if (remainder < 0) remainder += K; // Explained below

    count += freq[remainder];

    freq[remainder]++;
  }
  return count;
};
```
