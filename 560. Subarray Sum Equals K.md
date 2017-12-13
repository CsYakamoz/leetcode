## [560. Subarray Sum Equals K](https://leetcode.com/problems/subarray-sum-equals-k/description/)

### Description

Given an array of integers and an integer **k**, you need to find the total number of continuous subarrays whose sum equals to **k**.

**Example 1:**

```
Input:nums = [1,1,1], k = 2
Output: 2
```

**Note:**

1. The length of the array is in range [1, 20,000].
2. The range of numbers in the array is [-1000, 1000] and the range of the integer **k** is [-1e7, 1e7].



**Difficult:** `Medium`

**Tags:** `Array` `Map`



### Solution One

```c++
class Solution
{
  public:
    int subarraySum(vector<int> &nums, int k)
    {
        unordered_map<int, int> hash;
        long long sum = 0;
        int res = 0;
        hash[0]++;

        for (auto i : nums)
        {
            sum += i;
            int diff = sum - k;
            if (hash.find(diff) != hash.end())
                res += hash[diff];
            hash[sum]++;
        }

        return res;
    }
};
```



