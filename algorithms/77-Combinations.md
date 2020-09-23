## [77. Combinations](https://leetcode.com/problems/combinations/description/)

### Description

Given two integers _n_ and _k_, return all possible combinations of _k_ numbers out of 1 ... _n_.

For example,
If _n_ = 4 and _k_ = 2, a solution is:

```
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
```

**Difficult:** `Medium`

**Tags:** `Backtracking`

### Solution One

```c++
class Solution
{
  public:
    vector<vector<int>> combine(int n, int k)
    {
        vector<int> vec;
        vector<vector<int>> res;
        helper(vec, 1, n, k, res);
        return res;
    }

  private:
    void helper(vector<int> &vec, int i, const int &n,
                const int &k, vector<vector<int>> &res)
    {
        // Stop Recursion early
        if (k - vec.size() > n - i + 1)
            return;

        if (vec.size() == k)
        {
            res.push_back(vec);
            return;
        }

        for (int j = i; j <= n; j++)
        {
            vec.push_back(j);
            helper(vec, j + 1, n, k, res);
            vec.pop_back();
        }
    }
};
```
