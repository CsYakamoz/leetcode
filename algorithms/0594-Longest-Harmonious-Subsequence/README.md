## [594. Longest Harmonious Subsequence](https://leetcode.com/problems/longest-harmonious-subsequence/description/)

### Description

We define a harmonious array is an array where the difference between its maximum value and its minimum value is **exactly** 1.

Now, given an integer array, you need to find the length of its longest harmonious subsequence among all its possible [subsequences](https://en.wikipedia.org/wiki/Subsequence).

**Example 1:**

```
Input: [1,3,2,2,5,2,3,7]
Output: 5
Explanation: The longest harmonious subsequence is [3,2,2,2,3].

```

**Note:** The length of the input array will not exceed 20,000.

**Difficulty:** `Easy`

**Tags:** `Hash Table`

### Solution One

```c++
class Solution {
public:
    int findLHS(vector<int>& nums) {
        if (nums.empty())
        {
            return 0;
        }
        map<int, int> m;
        for (auto i : nums)
        {
            m[i]++;
        }
        int res = 0;
        auto iter = m.begin();
        int pre_val = iter->first;
        int pre_count = iter->second;
        while (++iter != m.end())
        {
            if (iter->first - pre_val == 1)
            {
                res = max(res, iter->second + pre_count);
            }
            pre_val = iter->first;
            pre_count = iter->second;
        }
        return res;
    }
};
```

### Other Solutions

[594. Longest Harmonious Subsequence - Solution](https://leetcode.com/problems/longest-harmonious-subsequence/solution/)
