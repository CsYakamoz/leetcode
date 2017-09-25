## [532. K-diff Pairs in an Array](https://leetcode.com/problems/k-diff-pairs-in-an-array/description/)

### Description

Given an array of integers and an integer **k**, you need to find the number of **unique** k-diff pairs in the array. Here a **k-diff** pair is defined as an integer pair (i, j), where **i** and **j** are both numbers in the array and their [absolute difference](https://en.wikipedia.org/wiki/Absolute_difference) is **k**.

**Example 1:**

```
Input: [3, 1, 4, 1, 5], k = 2
Output: 2
Explanation: There are two 2-diff pairs in the array, (1, 3) and (3, 5).
Although we have two 1s in the input, we should only return the number of unique pairs.

```

**Example 2:**

```
Input:[1, 2, 3, 4, 5], k = 1
Output: 4
Explanation: There are four 1-diff pairs in the array, (1, 2), (2, 3), (3, 4) and (4, 5).
```

**Example 3:**

```
Input: [1, 3, 1, 5, 4], k = 0
Output: 1
Explanation: There is one 0-diff pair in the array, (1, 1).

```

**Note:**

1. The pairs (i, j) and (j, i) count as the same pair.
2. The length of the array won't exceed 10,000.
3. All the integers in the given input belong to the range: [-1e7, 1e7].



**Difficult:** `Easy`

**Tags:** `Array` `Two Pointers`



### Solution One

```c++
class Solution {
public:
    int findPairs(vector<int>& nums, int k) {
        sort(nums.begin(), nums.end());

        int res = 0;
        size_t i = 0;
        size_t j = 1;
        while (j < nums.size())
        {
            if (i >= j)
            {
                j = i + 1;
            }
            else if (nums[j] - nums[i] == k)
            {
                res++;
                i++;
                while (i < nums.size() && nums[i] == nums[i - 1])
                    i++;
            }
            else if (nums[j] - nums[i] < k)
            {
                j++;
            }
            else
            {
                i++;
            }
        }
        
        return res;
    }
};
```



### Solution Two - In Top Solutions

```c++
class Solution {
public:
    int findPairs(vector<int>& nums, int k) {
        if (k < 0)
            return 0;

        unordered_map<int, int> hash;
        int res = 0;

        for (int i : nums)
            ++hash[i];

        if (k != 0)
        {
            for (auto p: hash)
            {
                if (hash.find(p.first - k) != hash.end())
                    res++;
            }
        }
        else
        {
            for (auto p: hash)
            {
                if (p.second > 1)
                    res++;
            }
        }

        return res;
    }
};
```



