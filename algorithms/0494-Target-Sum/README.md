## [494. Target Sum](https://leetcode.com/problems/target-sum/#/description)

### Description

You are given a list of non-negative integers, a1, a2, ..., an, and a target, S. Now you have 2 symbols `+` and `-`. For each integer, you should choose one from `+` and `-` as its new symbol.

Find out how many ways to assign symbols to make sum of integers equal to target S.

**Example 1:**

```
Input: nums is [1, 1, 1, 1, 1], S is 3.
Output: 5
Explanation:

-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

There are 5 ways to assign symbols to make the sum of nums be target 3.
```

**Note:**

1. The length of the given array is positive and will not exceed 20.
2. The sum of elements in the given array will not exceed 1000.
3. Your output answer is guaranteed to be fitted in a 32-bit integer.

**Difficulty:** `Medium`

**Tags:** `Depth-first Search` `Dynamic Programming`

### Solution One

`DFS`

```c++
class Solution {
public:
    int findTargetSumWays(vector<int>& nums, int S) {
        if (nums.empty())
        {
            return 0;
        }
        target = S;
        dfs(nums.size() - 1, nums, 0);
        return res;
    }

private:
    int res = 0;
    int target;

    void dfs(size_t i, vector<int> &nums, int sum)
    {
        if (i == 0)
        {
            if (sum + nums[0] == target)
            {
                res++;
            }
            if (sum - nums[0] == target)
            {
                res++;
            }
        }
        else
        {
            dfs(i - 1, nums, sum + nums[i]);
            dfs(i - 1, nums, sum - nums[i]);
        }
    }
};
```

### Solution Two

`DFS with Memorization`

```c++
class Solution {
public:
    int findTargetSumWays(vector<int>& nums, int S) {
        if (nums.empty())
        {
            return 0;
        }
        target = S;
        return dfs_with_memorization(0, 0, nums);

    }
private:
    int target;
    map<string, int> hash;

    int dfs_with_memorization(size_t i, int sum, vector<int> &nums)
    {
        string s = to_string(i) + "->" + to_string(sum);
        if (hash.find(s) != hash.end())
        {
            return hash[s];
        }
        if (i == nums.size())
        {
            if (sum == target )
            {
                return 1;
            }
            else
            {
                return 0;
            }
        }
        int add = dfs_with_memorization(i + 1, sum + nums[i], nums);
        int minus = dfs_with_memorization(i + 1, sum - nums[i], nums);
        hash[s] = add + minus;
        return add + minus;
    }
};
```

### Other Solutions - In Top Solutions

[494. Target Sum - Solution](https://leetcode.com/problems/target-sum/solution/)
