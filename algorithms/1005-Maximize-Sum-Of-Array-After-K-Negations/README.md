## [1005. Maximize Sum Of Array After K Negations](https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/)

### Description

Given an array `A` of integers, we **must** modify the array in the following way: we choose an `i` and replace `A[i]` with `-A[i]`, and we repeat this process `K` times in total. (We may choose the same index `i` multiple times.)

Return the largest possible sum of the array after modifying it in this way.

**Difficult:** `Easy`

**Tags:** `Greedy`

### Solution One

1. 首先遍历一次数组，并做如下事：
   - 分别记录正数，负数的和
   - 分别记录最小正数，最大负数
   - 将负数和零记录到数组 `negativeAndZeroList`
2. 如果 `negativeAndZeroList` 的长度为 0, 则只能选择最小的正数进行变换
3. 如果 `negativeAndZeroList` 的长度小于等于 K, 则做如下事：
   - 将 `negativeAndZeroList` 都变为正数
   - 此时所有数字都是正数，且还剩 `K - negativeAndZeroList.length` 次变换
   - 情况如同 2
4. 如果 `negativeAndZeroList` 的长度大于 K, 则做如下事：
   - 将 `negativeAndZeroList` 按从小到大排序
   - 将前 K 个元素变成正数

```javascript
/**
 * @param {number[]} A
 * @param {number} K
 * @return {number}
 */
const largestSumAfterKNegations = function (A, K) {
  let positiveSum = 0;
  let negativeSum = 0;
  let maxNegative = Number.MIN_SAFE_INTEGER;
  let minPositive = Number.MAX_SAFE_INTEGER;

  const negativeAndZeroList = [];

  for (const val of A) {
    if (val > 0) {
      positiveSum += val;
      minPositive = Math.min(minPositive, val);
    } else {
      negativeAndZeroList.push(val);
      maxNegative = Math.max(maxNegative, val);
      negativeSum += val;
    }
  }

  const nzLen = negativeAndZeroList.length;

  if (nzLen === 0) {
    return positiveSum - (K % 2 === 1 ? minPositive * 2 : 0);
  }

  if (nzLen <= K) {
    positiveSum -= negativeSum;
    negativeSum = 0;

    const remain = K - nzLen;
    if (remain % 2 !== 0) {
      const min = Math.min(minPositive, -maxNegative);
      negativeSum -= min;
      positiveSum -= min;
    }
  } else {
    negativeAndZeroList.sort((a, b) => a - b);
    for (let i = 0; i < K; i++) {
      const val = negativeAndZeroList[i];
      negativeSum -= val;
      positiveSum -= val;
    }
  }

  return positiveSum + negativeSum;
};
```
