## [525. Contiguous Array](https://leetcode.com/problems/contiguous-array/description/)

### Description

Given a binary array, find the maximum length of a contiguous subarray with equal number of 0 and 1.

**Example 1:**

```
Input: [0,1]
Output: 2
Explanation: [0, 1] is the longest contiguous subarray with equal number of 0 and 1.

```

**Example 2:**

```
Input: [0,1,0]
Output: 2
Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.

```

**Note:** The length of the given binary array will not exceed 50,000.

**Difficulty:** `Medium`

**Tags:** `Hash Table`

### Solution One - In Top Solutions

```c++
class Solution {
public:
    int findMaxLength(vector<int>& nums) {
        map<int, int> hash;
        hash[0] = -1;
        int maxLength = 0;
        int count = 0;
        for (int i = 0; i < nums.size(); i++) {
            count += nums[i] == 1 ? 1 : -1;
            if (hash.find(count) != hash.end())
                maxLength = max(maxLength, i - hash[count]);
            else
                hash[count] = i;
        }
        return maxLength;
    }
};
```

### Other Solutions

[525. Contiguous Array - Solution](https://leetcode.com/problems/contiguous-array/solution/)
