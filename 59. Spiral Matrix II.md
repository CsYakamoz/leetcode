## [59. Spiral Matrix II](https://leetcode.com/problems/spiral-matrix-ii/description/)

### Description

Given an integer *n*, generate a square matrix filled with elements from 1 to *n*2 in spiral order.

For example,
Given *n* = `3`,

You should return the following matrix:

```
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
```



**Difficult:** `Medium`

**Tags:** `Array`



### Solution One

来源于题目 [54. Spiral Matrix](https://leetcode.com/problems/spiral-matrix/description/) 中的 [Super Simple and Easy to Understand Solution](https://discuss.leetcode.com/topic/3713/super-simple-and-easy-to-understand-solution)

```c++
class Solution {
public:
    vector<vector<int>> generateMatrix(int n) {
        if (n <= 0)
        {
            return {};
        }

        vector<vector<int>> matrix(n, vector<int>(n, 0));
        int val = 1;
        int rowBegin = 0;
        int rowEnd = n - 1;
        int colBegin = 0;
        int colEnd = n - 1;

        while (rowBegin <= rowEnd && colBegin <= colEnd)
        {
            // Right
            for (int i = colBegin; i <= colEnd; i++)
            {
                matrix[rowBegin][i] = val++;
            }
            rowBegin++;

            // Down
            for (int i = rowBegin; i <= rowEnd; i++)
            {
                matrix[i][colEnd] = val++;
            }
            colEnd--;

            // Left
            for (int i = colEnd; i >= colBegin; i--)
            {
                matrix[rowEnd][i] = val++;
            }
            rowEnd--;

            // Up
            for (int i = rowEnd; i >= rowBegin; i--)
            {
                matrix[i][colBegin] = val++;
            }
            colBegin++;
        }
        return matrix;
    }
};
```



