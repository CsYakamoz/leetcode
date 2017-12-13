## [200. Number of Islands](https://leetcode.com/problems/number-of-islands/description/)

### Description

Given a 2d grid map of `'1'`s (land) and `'0'`s (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

***Example 1:***

```
11110
11010
11000
00000
```

Answer: 1

***Example 2:***

```
11000
11000
00100
00011
```

Answer: 3



**Difficult:** `Medium`

**Tags:** `Depth-first Search` `Breadth-first Search` `Union Find`



### Solution One

```c++
class Solution {
public:
    int numIslands(vector<vector<char>>& grid) {
        if (grid.empty() || grid[0].empty()) {
            return 0;
        }

        int count = 0;
        m = grid.size();
        n = grid[0].size();

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == '1') {
                    dfsMarking(i, j, grid);
                    count++;
                }
            }
        }

        return count;
    }

private:
    int m, n;

    void dfsMarking(int i, int j, vector<vector<char>> &grid) {
        if (i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == '0') {
            return;
        }

        grid[i][j] = '0';

        dfsMarking(i - 1, j, grid);
        dfsMarking(i + 1, j, grid);
        dfsMarking(i, j - 1, grid);
        dfsMarking(i, j + 1, grid);
    }
};
```



