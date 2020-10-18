## [120. Triangle](https://leetcode.com/problems/triangle/description/)

### Description

Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

For example, given the following triangle

```
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]

```

The minimum path sum from top to bottom is `11` (i.e., 2 + 3 + 5 + 1 = 11).

**Note:**
Bonus point if you are able to do this using only _O_(_n_) extra space, where _n_ is the total number of rows in the triangle.

**Difficulty:** `Medium`

**Tags:** `Array` `Dynamic Programming`

### Solution One

```c++
class Solution {
public:
    int minimumTotal(vector<vector<int>>& triangle) {
        vector<int> dp = triangle.back();
        for (int i = triangle.size() - 2; i >= 0; i--)
        {
            for (int j = 0; j <= i; j++)
            {
                dp[j] = min(dp[j], dp[j + 1]) + triangle[i][j];
            }
        }
        return dp[0];
    }
};
```
