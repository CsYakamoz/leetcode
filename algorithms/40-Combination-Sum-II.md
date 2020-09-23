## [40. Combination Sum II](https://leetcode.com/problems/combination-sum-ii/description/)

### Description

Given a collection of candidate numbers (**C**) and a target number (**T**), find all unique combinations in **C** where the candidate numbers sums to **T**.

Each number in **C** may only be used **once** in the combination.

**Note:**

- All numbers (including target) will be positive integers.
- The solution set must not contain duplicate combinations.

For example, given candidate set `[10, 1, 2, 7, 6, 1, 5]` and target `8`,
A solution set is:

```
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
```

**Difficult:** `Medium`

**Tags:** `Array` `Backtracking`

### Solution One

```c++
class Solution {
public:
    vector<vector<int>> combinationSum2(vector<int>& candidates, int target) {
        res.clear();
        if (candidates.empty()) return res;

        vector<int> vec;

        sort(candidates.begin(), candidates.end());
        helper(0, 0, target, vec, candidates);
        return res;
    }

private:
    vector<vector<int>> res;

    void helper(int idx, int sum, const int &target, vector<int> &vec, const vector<int> &candidates) {
        if (sum == target) {
            res.push_back(vec);
            return;
        }

        for (int i = idx; i < candidates.size(); i++) {
            if (sum + candidates[i] > target)
                break;
            vec.push_back(candidates[i]);
            helper(i + 1, sum + candidates[i], target, vec, candidates);
            vec.pop_back();
            while (i < candidates.size() - 1 && candidates[i + 1] == candidates[i])
                i++;
        }
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
    void helper(vector<vector<int>> &res, vector<int> & nums,
               vector<int> &cur, int beg, int target)
    {
        if(target == 0)
        {
            res.push_back(cur);
            return;
        }
        int prev = 0;
        for(int i = beg; i < nums.size(); i++)
        {
            if(i != beg && nums[i] == prev ) continue;
            if(target - nums[i] >= 0)
            {
                cur.push_back(nums[i]);
                helper(res, nums, cur, i+1, target - nums[i]);
                cur.pop_back();
            }
            prev = nums[i];
        }
        return;
    }
public:
    vector<vector<int>> combinationSum2(vector<int>& candidates, int target)
    {
        vector<vector<int>> res;
        vector<int> cur;
        sort(candidates.begin(), candidates.end());
        helper(res, candidates, cur, 0, target);
        return res;
    }
};
```
