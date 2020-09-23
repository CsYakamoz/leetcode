## [661. Image Smoother](https://leetcode.com/problems/image-smoother/description/)

### Description

Given a 2D integer matrix M representing the gray scale of an image, you need to design a smoother to make the gray scale of each cell becomes the average gray scale (rounding down) of all the 8 surrounding cells and itself. If a cell has less than 8 surrounding cells, then use as many as you can.

**Example 1:**

```
Input:
[[1,1,1],
 [1,0,1],
 [1,1,1]]
Output:
[[0, 0, 0],
 [0, 0, 0],
 [0, 0, 0]]
Explanation:
For the point (0,0), (0,2), (2,0), (2,2): floor(3/4) = floor(0.75) = 0
For the point (0,1), (1,0), (1,2), (2,1): floor(5/6) = floor(0.83333333) = 0
For the point (1,1): floor(8/9) = floor(0.88888889) = 0

```

**Note:**

1. The value in the given matrix is in the range of [0, 255].
2. The length and width of the given matrix are in the range of [1, 150].

**Difficult:** `Easy`

**Tags:** `Array`

### Solution One

```c++
class Solution {
public:
    vector<vector<int>> imageSmoother(vector<vector<int>>& M) {
        int m = M.size();
        int n = M.front().size();
        int length = m * n;
        vector<vector<int>> res(m, vector<int>(n, 0));

        for (int i = 0; i < length; i++)
        {
            int row = (i / n) - 1;
            int col = (i % n) - 1;
            int val = 0;
            int count = 0;
            for (int r = 0; r < 3; r++)
            {
                if (row + r < 0 || row + r >= m)
                    continue;
                for (int c = 0; c < 3; c++)
                {
                    if (col + c < 0 || col + c >= n)
                        continue;
                    val += M[row + r][col + c];
                    count++;
                }
            }
            res[i / n][i % n] = val / count;
        }

        return res;
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    vector<vector<int>> imageSmoother(vector<vector<int>>& M) {
        if (!M.size()|| !M[0].size()) {
            return vector<vector<int>>();
        }
        vector<vector<int>> res(M.size(),vector<int>(M[0].size()));
        for (int i = 0; i < M.size(); i++) {
            for (int j = 0; j < M[i].size(); j++) {
                int sum = M[i][j];
                int count = 1;
                if ( i-1 >= 0) {
                    sum += M[i-1][j];
                    count ++;
                }
                if ( j-1 >= 0) {
                    sum += M[i][j-1];
                    count ++;
                }
                if ( i+1 < M.size() ) {
                    sum += M[i+1][j];
                    count ++;
                }
                if ( j+1 < M[i].size() ) {
                    sum += M[i][j+1];
                    count ++;
                }
                if ( i-1>=0 && j-1 >=0 ) {
                    sum += M[i-1][j-1];
                    count ++;
                }
                if ( i-1>=0 && j+1 < M[i].size() ) {
                    sum += M[i-1][j+1];
                    count ++;
                }
                if ( i+1 < M.size() && j-1 >=0 ) {
                    sum += M[i+1][j-1];
                    count ++;
                }
                if ( i+1 < M.size() && j+1 < M[i].size()) {
                    sum += M[i+1][j+1];
                    count ++;
                }
                res[i][j] = sum / count;
            }
        }
        return res;

    }
};
```
