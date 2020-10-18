## [645. Set Mismatch](https://leetcode.com/problems/set-mismatch/tabs/description)

### Description

The set `S` originally contains numbers from 1 to `n`. But unfortunately, due to the data error, one of the numbers in the set got duplicated to **another** number in the set, which results in repetition of one number and loss of another number.

Given an array `nums` representing the data status of this set after the error. Your task is to firstly find the number occurs twice and then find the number that is missing. Return them in the form of an array.

**Example 1:**

```
Input: nums = [1,2,2,4]
Output: [2,3]

```

**Note:**

1. The given array size will in the range [2, 10000].
2. The given array's numbers won't have any order.

**Difficulty:** `Easy`

**Tags:** `Hash Table` `Math`

### Solution One

```c++
class Solution {
public:
    vector<int> findErrorNums(vector<int>& nums) {
        vector<int> index(nums.size(), 0);
        // sum = (a1 + an) * n / 2;
        int sum = (1 + nums.size()) * nums.size() / 2;
        vector<int> res;
        for (size_t i = 0; i < nums.size(); i++)
        {
            if (index[nums[i]-1] == 0)
            {
                index[nums[i] - 1] = 1;
                sum -= nums[i];
            }
            else
            {
                res.push_back(nums[i]);
            }
        }
        res.push_back(sum);
        return res;
    }
};
```

### Other Solutions

[645. Set Mismatch - Solutions](https://leetcode.com/problems/set-mismatch/solution/)
