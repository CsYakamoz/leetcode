## [1590. Make Sum Divisible by P](https://leetcode.com/problems/make-sum-divisible-by-p/)

### Description

Given an array of positive integers `nums`, remove the **smallest** subarray (possibly **empty**) such that the **sum** of the remaining elements is divisible by `p`. It is **not** allowed to remove the whole array.

Return _the length of the smallest subarray that you need to remove, or_ `-1` _if it's impossible_.

A **subarray** is defined as a contiguous block of elements in the array.

**Example 1:**

```
Input: nums = [3,1,4,2], p = 6
Output: 1
Explanation: The sum of the elements in nums is 10, which is not divisible by 6. We can remove the subarray [4], and the sum of the remaining elements is 6, which is divisible by 6.
```

**Example 2:**

```
Input: nums = [6,3,5,2], p = 9
Output: 2
Explanation: We cannot remove a single element to get a sum divisible by 9. The best way is to remove the subarray [5,2], leaving us with [6,3] with sum 9.
```

**Example 3:**

```
Input: nums = [1,2,3], p = 3
Output: 0
Explanation: Here the sum is 6. which is already divisible by 3. Thus we do not need to remove anything.
```

**Example 4:**

```
Input: nums = [1,2,3], p = 7
Output: -1
Explanation: There is no way to remove a subarray in order to get a sum divisible by 7.
```

**Example 5:**

```
Input: nums = [1000000000,1000000000,1000000000], p = 3
Output: 0
```

**Constraints:**

- `1 <= nums.length <= 10^5`
- `1 <= nums[i] <= 10^9`
- `1 <= p <= 10^9`

**Difficult:** `Medium`

**Tags:** `Array` `Binary Search`

### Solution One

首先有如下定义：

1. `prefixSum[i]`: `nums[0] + ... + nums[i]`

2. `remainder[i]`: `prefixSum[i] % p`

3. `expRemainder`: `sum(nums) % p`

思路如下：

设 **x** 为 `prefixSum[i] % p`, 则必然存在 **y** 满足 `(y + p - x ) % p === expRemainder`

寻找**最小**的 j, 使其满足 `j > i`, 且 `prefixSum[j] % p === y`

> 为何 `y = (x + expRemainder) % p` ? 证明如下：
>
> `(prefixSum[j] - prefixSum[i]) % p = expRemainder` 等式一
>
> 等式一两边乘以 p, 并且将 `prefixSum[i]` 移到右侧，可得到：
>
> `prefixSum[j] = prefixSum[i] + (expRemainder + np), 其中 n >= 0` 等式二
>
> 由于 `prefixSum[j]` 可写成 `remainder[j] + kp, 其中 k >= 0`
>
> 故等式二可转换成
>
> `(remainder[j] + ap) = (remainder[i] + bp) + (expRemainder + np), 其中 n >= 0, a >= 0, b >= 0` 等式三
>
> 将等式三左侧的 `ap` 移到右侧可得：
>
> `remainder[j] = remainder[i] + expRemainder + ap + bp + np` 等式四
>
> 等式四两边再次求余数可得：
>
> `remainder[j] = (remainder[i] + expRemainder) % p`
>
> 即 `y = (x + expRemainder) % p`

**注意:** 该代码可能有误, 因为忽略了此类情况:

在遍历 prefixRemainderDict 后, minLength 不为 Number.MAX_SAFE_INTEGER, 但存在 prefixSum[i] % p === expRemainder 且 i + 1 < minLength

```javascript
/**
 * @param {number[]} arr1
 * @param {number[]} arr2
 * @returns {number}
 */
const findMinDiff = (arr1, arr2) => {
  let minDiff = Number.MAX_SAFE_INTEGER;
  for (const i of arr1) {
    let begin = 0;
    let end = arr2.length;

    while (begin < end) {
      const mid = Math.floor((end - begin) / 2) + begin;
      if (arr2[mid] > i) {
        end = mid;
      } else {
        begin = mid + 1;
      }
    }
    if (end !== arr2.length) {
      minDiff = Math.min(minDiff, arr2[end] - i);
    }
  }

  return minDiff;
};

/**
 * @param {number[]} nums
 * @param {number} p
 * @return {number}
 */
const minSubarray = function (nums, p) {
  /** @type {Map<number, number[]>} */
  const prefixRemainderDict = new Map();

  const expRemainder = nums.reduce((acc, curr, idx) => {
    const remainder = (acc + curr) % p;
    if (!prefixRemainderDict.has(remainder)) {
      prefixRemainderDict.set(remainder, []);
    }
    prefixRemainderDict.get(remainder).push(idx);

    return remainder;
  }, 0);

  if (expRemainder === 0) {
    return 0;
  }

  let minLength = Number.MAX_SAFE_INTEGER;
  for (const [currRemainder, idxList] of prefixRemainderDict) {
    const anotherRemainder = (expRemainder + currRemainder) % p;

    if (!prefixRemainderDict.has(anotherRemainder)) {
      continue;
    }

    minLength = Math.min(
      minLength,
      findMinDiff(idxList, prefixRemainderDict.get(anotherRemainder))
    );
  }

  if (minLength !== Number.MAX_SAFE_INTEGER) {
    return minLength;
  } else if (prefixRemainderDict.has(expRemainder)) {
    const length = prefixRemainderDict.get(expRemainder)[0] + 1;
    return length !== nums.length ? length : -1;
  } else {
    return -1;
  }
};
```

### Solution Two - In Top Solutions

Link: [[Java/C++/Python] Prefix Sum](https://leetcode.com/problems/make-sum-divisible-by-p/discuss/854197/JavaC%2B%2BPython-Prefix-Sum)

思路本质上与 Solution One 一样，但 y 的定义和所寻找的 j 不一样

`y = (x - expRemainder + p ) % p`

寻找**最大**的 j, 使其满足 `j < i`, 且 `prefixSum[j] % p === y`

因此少了 Solution One 中所用到的 `findMinDiff` 函数

```javascript
/**
 * @param {number[]} nums
 * @param {number} p
 * @return {number}
 */
var minSubarray = function (nums, p) {
  let sum = 0;
  for (let num of nums) {
    sum += num;
  }
  const mod = sum % p;

  if (mod == 0) return 0;
  let map = new Map();
  map.set(0, -1);
  let curMod = 0;
  let min = Number.POSITIVE_INFINITY;
  for (let i = 0; i < nums.length; i++) {
    curMod = (curMod + nums[i]) % p;
    const req = (curMod - mod + p) % p;

    if (map.has(req)) {
      const n = map.get(req);

      min = Math.min(min, i - n);
    }
    map.set(curMod, i);
  }

  return min >= nums.length ? -1 : min;
};
```
