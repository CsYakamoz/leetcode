## [303. Range Sum Query - Immutable](https://leetcode.com/problems/range-sum-query-immutable/description/)

### Description

Given an integer array _nums_, find the sum of the elements between indices _i_ and _j_ (_i_ ≤ _j_), inclusive.

**Example:**

```
Given nums = [-2, 0, 3, -5, 2, -1]

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3
```

**Note:**

1. You may assume that the array does not change.
2. There are many calls to _sumRange_ function.

**Difficult:** `Easy`

**Tags:** `Dynamic Programming`

### Solution One

> Status: **Memory Limit Exceeded**
>
> **15 / 16** test cases passed

最后一个 test cases 中`nums`很长很长，长到申请 **n \* n** 矩阵会导致内存大小超过限制

```c++
class NumArray {
public:
    NumArray(vector<int> nums) {
        size_t length = nums.size();
        record = vector<vector<int>>(length, vector<int>(length, 0));
        for (size_t i = 0; i < length; i++)
        {
            record[i][i] = nums[i];
        }
        for (size_t i = 0; i < length; i++)
        {
            for (size_t j = i + 1; j < length; j++)
            {   // sumRange(i, j) = sumRange(i, j - 1) + nums[j]
                record[i][j] = record[i][j - 1] + nums[j];
            }
        }
    }

    int sumRange(int i, int j) {
        return record[i][j];
    }

private:
    vector<vector<int>> record;
};
```

### Solution Two

在 **One**，由于申请 **n \* n** 会溢出，故这里申请 **n \* n / 2 **的内存，然而这次轮到时间超过限制

个人认为原因是第一个 for 循环导致的，这里每次循环都申请内存，而 **One** 中是一次性申请内存

> Status: **Time Limit Exceeded**
>
> **15 / 16** test cases passed.

```c++
class NumArray {
public:
    NumArray(vector<int> nums) {
        size_t length = nums.size();
        record.resize(length);
        for (size_t i = 0; i < length; i++)
        {
            record[i] = vector<int>(length - i, 0);
        }
        for (size_t i = 0; i < length; i++)
        {
            record[i][0] = nums[i];
        }
        for (size_t i = 0; i < length; i++)
        {
            for (size_t j = 1; j < record[i].size(); j++)
            {
                record[i][j] = record[i][j - 1] + nums[j + i];
            }
        }
    }

    int sumRange(int i, int j) {
        return record[i][j - i];
    }

private:
    vector<vector<int>> record;
};
```

### Solution Three - In Top Solutions

[303. Range Sum Query - Solutions](https://leetcode.com/problems/range-sum-query-immutable/solution/)

基本思路为：$a_i + \dots + a_j$ = $(a_1 + \dots + a_j) - (a_1 + \dots + a_{i-1})$

```c++
class NumArray {
public:
    NumArray(vector<int> nums) {
        sum.resize(nums.size() + 1);
        for (size_t i = 0; i < nums.size(); i++)
        {
            sum[i + 1] = sum[i] + nums[i];
        }
    }

    int sumRange(int i, int j) {
        return sum[j + 1] - sum[i];
    }

private:
    vector<int> sum;
};
```
