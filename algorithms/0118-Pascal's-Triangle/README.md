## [118. Pascal's Triangle](https://leetcode.com/problems/pascals-triangle/description/)

### Description

Given _numRows_, generate the first _numRows_ of Pascal's triangle.

For example, given _numRows_ = 5,
Return

```
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
[1,4,6,4,1]
]
```

**Difficulty:** `Easy`

**Tags:** `Array`

### Solution One

```c++
class Solution {
public:
    vector<vector<int>> generate(int numRows) {
        vector<vector<int>> res;
        if (numRows == 0)
            return res;

        res.push_back(vector<int>{1});
        if (numRows == 1)
            return res;

        for (int i = 1; i < numRows; i++)
        {
            vector<int> &pre = res.back();
            vector<int> level{ 1 };

            for (int j = 1; j < i; j++)
                level.push_back(pre[j] + pre[j - 1]);

            level.push_back(1);
            res.push_back(level);
        }

        return res;
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    vector<vector<int>> generate(int numRows) {
        vector<vector<int>> re;
        for(int i = 0; i < numRows; i++){
            vector<int> row;

            row.push_back(1);

            for(int j = 0; j < i; j++){
                int elem = (j == re[i - 1].size() - 1 ?  0 : re[i - 1][j + 1]) + re[i - 1][j];
                row.push_back(elem);
            }
            re.push_back(row);
        }
       return re;
    }
};
```
