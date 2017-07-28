## [62. Unique Paths](https://leetcode.com/problems/unique-paths/#/description)

### Description

A robot is located at the top-left corner of a *m* x *n* grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?

![img](https://leetcode.com/static/images/problemset/robot_maze.png)

Above is a 3 x 7 grid. How many possible unique paths are there?

**Note:** *m* and *n* will be at most 100.



**Difficult:** `Medium`

**Tags:** `Array` `Dynamic Programming`



### Solution One

要知道到达`matrix[i][j]`有多少条不同的路径，那就得知道到达`matrix[i-1][j]`， ` matrix[i][j-1]`分别有多少条不同的路径。

因为到达`matrix[i][j]`的路只有`matrix[i-1][j]`或者`matrix[i][j-1]`

```c++
class Solution {
public:
    int uniquePaths(int m, int n) {
        vector<vector<int>> matrix(m, vector<int>(n, 1));
        for (size_t i = 1; i < m; i++)
        {
            for (size_t j = 1; j < n; j++)
            {
                matrix[i][j] = matrix[i - 1][j] + matrix[i][j - 1];
            }
        }
        return matrix[m - 1][n - 1];
    }
};
```



## Solution Two - In Top Solutions

[My AC solution using formula](https://discuss.leetcode.com/topic/2734/my-ac-solution-using-formula)

```c++
class Solution {
    public:
        int uniquePaths(int m, int n) {
            int N = n + m - 2;// how much steps we need to do
            int k = m - 1; // number of steps that need to go down
            double res = 1;
            // here we calculate the total possible path number 
            // Combination(N, k) = n! / (k!(n - k)!)
            // reduce the numerator and denominator and get
            // C = ( (n - k + 1) * (n - k + 2) * ... * n ) / k!
            for (int i = 1; i <= k; i++)
                res = res * (N - k + i) / i;
            return (int)res;
        }
    };
```



### Solution Three - In Top Solutions

[0ms, 5-lines DP Solution in C++ with Explanations](https://discuss.leetcode.com/topic/15265/0ms-5-lines-dp-solution-in-c-with-explanations)

