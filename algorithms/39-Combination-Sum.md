## [39. Combination Sum](https://leetcode.com/problems/combination-sum/description/)

### Description

Given a **set** of candidate numbers (**C**) **(without duplicates)** and a target number (**T**), find all unique combinations in **C** where the candidate numbers sums to **T**.

The **same** repeated number may be chosen from **C** unlimited number of times.

**Note:**

- All numbers (including target) will be positive integers.
- The solution set must not contain duplicate combinations.

For example, given candidate set `[2, 3, 6, 7]` and target `7`,
A solution set is:

```
[
  [7],
  [2, 2, 3]
]
```

**Difficult:** `Medium`

**Tags:** `Array` `Backtracking`

### Solution One

```c++
class Solution {
public:
    vector<vector<int>> combinationSum(vector<int>& candidates, int target) {
        vector<int> vec;

        res.clear();
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

        for (size_t i = idx; i < candidates.size(); i++) {
            sum += candidates[i];
            // Stop Recursion Early
            if (sum > target) break;
            vec.push_back(candidates[i]);
            helper(i, sum, target, vec, candidates);
            vec.pop_back();
            sum -= candidates[i];
        }
    }
};
```
