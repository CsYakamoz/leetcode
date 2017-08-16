## [498. Diagonal Traverse](https://leetcode.com/problems/diagonal-traverse/description/)

### Description

Given a matrix of M x N elements (M rows, N columns), return all elements of the matrix in diagonal order as shown in the below image.

**Example:**

```
Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output:  [1,2,4,7,5,3,6,8,9]
Explanation:
```

<img src="https://leetcode.com/static/images/problemset/diagonal_traverse.png" width="20%">

**Note:**

1. The total number of elements of the given matrix will not exceed 10,000.



**Difficult:** `Medium`

**Tags:** 



### Solution One

首先将每条对角线添加到`diagonal`

然后，根据其是第几条对角线便可知道方向

```c++
class Solution {
public:
    vector<int> findDiagonalOrder(vector<vector<int>>& matrix) {
        vector<int> res;
        if (matrix.empty() || matrix[0].empty())
        {
            return res;
        }
        int m = matrix.size();
        int n = matrix[0].size();
        int i = 0;
        int j = 0;
        bool dir = true;
        while (j < n)
        {
            vector<int> diagonal;
            int _i = i;
            int _j = j;
            while (_i >= 0 && _j < n)
            {
                diagonal.push_back(matrix[_i][_j]);
                _i--;
                _j++;
            }
            if (dir)
            {   // left bottom to right top
                res.insert(res.end(), diagonal.begin(), diagonal.end());
            }
            else
            {   // right top toleft bottom
                res.insert(res.end(), diagonal.rbegin(), diagonal.rend());
            }
            dir = !dir;
            if (i == m - 1)
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

[Concise Java Solution](https://discuss.leetcode.com/topic/77865/concise-java-solution)

> I don't think this is a hard problem. It is easy to figure out the walk pattern. Anyway...
> Walk patterns:
>
> - If out of `bottom border` (row>= m) then row = m - 1; col += 2; change walk direction.
> - if out of `right border` (col>= n) then col = n - 1; row += 2; change walk direction.
> - if out of `top border` (row < 0) then row = 0; change walk direction.
> - if out of `left border` (col < 0) then col = 0; change walk direction.
> - Otherwise, just go along with the current direction.
>
> Time complexity: O(m * n), m = number of rows, n = number of columns.
> Space complexity: O(1).

```java
public class Solution {
    public int[] findDiagonalOrder(int[][] matrix) {
        if (matrix == null || matrix.length == 0) return new int[0];
        int m = matrix.length, n = matrix[0].length;
        
        int[] result = new int[m * n];
        int row = 0, col = 0, d = 0;
        int[][] dirs = {{-1, 1}, {1, -1}};
        
        for (int i = 0; i < m * n; i++) {
            result[i] = matrix[row][col];
            row += dirs[d][0];
            col += dirs[d][1];
            
            if (row >= m) { row = m - 1; col += 2; d = 1 - d;}
            if (col >= n) { col = n - 1; row += 2; d = 1 - d;}
            if (row < 0)  { row = 0; d = 1 - d;}
            if (col < 0)  { col = 0; d = 1 - d;}
        }
        
        return result;
    }
}
```



