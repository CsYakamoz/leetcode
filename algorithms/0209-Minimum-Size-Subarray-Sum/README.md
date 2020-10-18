## [209. Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/description/)

### Description

Given an array of **n** positive integers and a positive integer **s**, find the minimal length of a **contiguous** subarray of which the sum ≥ **s**. If there isn't one, return 0 instead.

For example, given the array `[2,3,1,2,4,3]` and `s = 7`,
the subarray `[4,3]` has the minimal length under the problem constraint.

**More practice:**

If you have figured out the _O_(_n_) solution, try coding another solution of which the time complexity is _O_(_n_ log _n_).

**Difficulty:** `Medium`

**Tags:** `Array` `Two Pointers` `Binary Search`

### Solution One

```c++
class Solution {
public:
    int minSubArrayLen(int s, vector<int>& nums) {

        if (nums.empty()) return 0;

        int i = 0;
        int j = 0;
        int length = INT_MAX;
        int sum = 0;

        while (j < nums.size()) {
            sum += nums[j++];

            while (sum >= s) {
                length = min(length, j - i);
                sum -= nums[i++];
            }
        }

        return length == INT_MAX ? 0 : length;
    }
};
```

### Solutions

[209. Minimum Size Subarray Sum - Solution](https://leetcode.com/problems/minimum-size-subarray-sum/solution/#approach-3-using-binary-search-accepted)
