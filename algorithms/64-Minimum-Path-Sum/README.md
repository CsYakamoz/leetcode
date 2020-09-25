## [64. Minimum Path Sum](https://leetcode.com/problems/minimum-path-sum/description/)

### Description

Given a _m_ x _n_ grid filled with non-negative numbers, find a path from top left to bottom right which _minimizes_ the sum of all numbers along its path.

**Note:** You can only move either down or right at any point in time.

**Difficult:** `Medium`

**Tags:** `Array` `Dynamic Programming`

### Solution One

```c++
class Solution {
public:
    int minPathSum(vector<vector<int>>& grid) {
        int m = grid.size();
        int n = grid[0].size();
        for (int i = 1; i < m; i++)
        {
            grid[i][0] += grid[i - 1][0];
        }
        for (int j = 1; j < n; j++)
        {
            grid[0][j] += grid[0][j - 1];
        }
        for (int i = 1; i < m; i++)
        {
            for (int j = 1; j < n; j++)
            {
                grid[i][j] += min(grid[i][j - 1], grid[i - 1][j]);
            }
        }
        return grid[m - 1][n - 1];
    }
};
```

### Solution Two

```c++
class Solution {
public:
    int minPathSum(vector<vector<int>>& grid) {
        int m = grid.size();
        int n = grid[0].size();
        deque<int> dp(grid[0].cbegin(), grid[0].cend());
        dp.push_front(0);
        for (int j = 1; j <= n; j++)
        {
            dp[j] += dp[j - 1];
        }
        dp[0] = dp[1];
        for (int i = 1; i < m; i++)
        {
            for (int j = 1; j <= n; j++)
            {
                dp[j] = min(dp[j], dp[j - 1]) + grid[i][j - 1];
            }
            dp[0] = dp[1];
        }
        return dp[n];
    }
};
```
