## [216. Combination Sum III](https://leetcode.com/problems/combination-sum-iii/description/)

### Description

Find all possible combinations of ***k*** numbers that add up to a number ***n***, given that only numbers from 1 to 9 can be used and each combination should be a unique set of numbers.

***Example 1:***

Input: **\*k*** = 3, **\*n*** = 7

Output:

```
[[1,2,4]]

```

***Example 2:***

Input: **\*k*** = 3, **\*n*** = 9

Output:

```
[[1,2,6], [1,3,5], [2,3,4]]
```



**Difficult:** `Medium`

**Tags:** `Array` `Backtracking`



### Solution One

```c++
class Solution {
public:
    vector<vector<int>> combinationSum3(int k, int n) {
        int digits[9] = { 1,2,3,4,5,6,7,8,9 };
        vector<int> nums;
        vector<vector<int>> res;

        if (n < 1 || n > 45) return res;

        helper(0, 0, k, n, digits, nums, res);
        return res;
    }

    void helper(int idx, int sum, int k, const int &target,
                int *digits, vector<int> &nums, vector<vector<int>> &res) {
        if (k == 0) {
            if (sum == target) res.push_back(nums);
            return;
        }

        for (int i = idx; i < 9; i++) {
            if (sum + digits[i] > target || 9 - i < k) break;
            nums.push_back(digits[i]);
            helper(i + 1, sum + digits[i], k - 1, target, digits, nums, res);
            nums.pop_back();
        }
    }
};
```



### Solution Two - In Top Solutions

```c++
class Solution {
public:
    void combination(vector<vector<int>>& result, vector<int> sol, int k, int n){
        if(sol.size() == k && n == 0){
            result.push_back(sol);
            return;
        }
        if(sol.size()<k){
            for(int i = sol.empty() ? 1 : sol.back() + 1; i<=9; i++){
                if(n-i < 0)
                    break;
                sol.push_back(i);
                combination(result,sol,k,n-i);
                sol.pop_back();
            }
        }
    }
    
    vector<vector<int>> combinationSum3(int k, int n) {
        vector<vector<int>> result;
        vector<int> sol;
        combination(result,sol,k,n);
        return result;
    }
};
```



