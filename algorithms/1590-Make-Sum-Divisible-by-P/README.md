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

1. 遍历 `nums`, 以 `remainder[i]` 为聚合标准，将对应索引存放在同一个数组中；

2. 遍历 `prefixRemainderDict`, 做如下事情：

   对于 `currRemainder` 所映射数组的每一个索引 `x`, 寻找**最小**的索引 `y`, 使得

   `y > x`, 且 `(prefixSum[y] - prefixSum[x]) % p = expRemainder` 且 `y` 是在 `(currRemainder + expRemainder) % p` 所映射数组中寻找

   为何是该数组中寻找呢？原因如下：

   > `(prefixSum[y] - prefixSum[x]) % p = expRemainder` 等式一
   >
   > 等式一两边乘以 p, 可得到：
   >
   > `prefixSum[y] = prefixSum[x] + (expRemainder + np), 其中 n >= 0` 等式二
   >
   > 由于 `prefixSum[i]` 可写成 `remainder[i] + kp, 其中 k >= 0`
   >
   > 故等式二可转换成
   >
   > `(remainder[y] + ap) = (remainder[x] + bp) + (expRemainder + np), 其中 n >= 0, a >= 0, b >= 0` 等式三
   >
   > 将等式三左侧的 `ap` 移到右侧可得：
   >
   > `remainder[y] = remainder[x] + expRemainder + ap + bp + np` 等式四
   >
   > 等式四两边再次求余数可得：
   >
   > `remainder[y] = (remainder[x] + expRemainder) % p`
   >
   > 故 `anotherRemainder = (currRemainder + expRemainder) % p`

   所以我们只需要在 `anotherRemainder` 所映射数组中寻找即可

3. 如果步骤 2 中未找到满足条件的 `y`, 则判断 `prefixRemainderDict` 有无 key 为 `expRemainder`

   若有，且数组第一个元素值不等于 `nums.length - 1`, 则返回该元素与 1 的和，否则返回 `-1`

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

TODO
