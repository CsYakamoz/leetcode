## [565. Array Nesting](https://leetcode.com/problems/array-nesting/description/)

### Description

A zero-indexed array A consisting of N different integers is given. The array contains all integers in the range [0, N - 1].

Sets S[K] for 0 <= K < N are defined as follows:

S[K] = { A[K], A[A[K]], A[A[A[K]]], ... }.

Sets S[K] are finite for each K and should NOT contain duplicates.

Write a function that given an array A consisting of N integers, return the size of the largest set S[K] for this array.

**Example 1:**

```
Input: A = [5,4,0,3,1,6,2]
Output: 4
Explanation:
A[0] = 5, A[1] = 4, A[2] = 0, A[3] = 3, A[4] = 1, A[5] = 6, A[6] = 2.

One of the longest S[K]:
S[0] = {A[0], A[5], A[6], A[2]} = {5, 6, 2, 0}
```

**Note:**

1. N is an integer within the range [1, 20,000].
2. The elements of A are all distinct.
3. Each element of array A is an integer within the range [0, N-1].

**Difficult:** `Medium`

**Tags:** `Array`

### Solution One

基本思路（暴力解法）：对每一个 K 进行寻找

> Status: Time Limit Exceeded

```c++
class Solution {
public:
    int arrayNesting(vector<int>& nums) {
        int res = INT_MIN;
        for (int k = 0; k < nums.size(); k++)
        {
            int length = 0;
            int next = k;
            do
            {
                next = nums[next];
                length++;
            } while (next != k);
            res = max(res, length);
        }
        return res;
    }
};
```

### Solution Two

`DFS`

以例子 `A = [5,4,0,3,1,6,2]` 中的 `S[0] = {A[0], A[5], A[6], A[2]} = {5, 6, 2, 0}` 为例：

当 K 值为 S[0] 的任意一个元素时， S[K] 的长度等于 S[0]，即 S[5] = S[6] = S[2] = S[0]，

```c++
class Solution {
public:
    int arrayNesting(vector<int>& nums) {
        record = vector<bool>(nums.size(), false);
        int res = 0;
        for (size_t k = 0; k < nums.size(); k++)
        {
            if (record[k] == false)
            {
                res = max(res, helper(k, k, nums, 0));
            }
        }
        return res;
    }

private:
    vector<bool> record;

    int helper(const int k, int next, vector<int> &nums, int length)
    {
        if (nums[next] == k)
        {
            record[next] = true;
            return length + 1;
        }
        record[next] = true;
        return helper(k, nums[next], nums, length + 1);
    }
};
```

### Other Solutions

[565 Array Nesting - Solution](https://leetcode.com/problems/array-nesting/solution/#approach-3-without-using-extra-space-accepted)
