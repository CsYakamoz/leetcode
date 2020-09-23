## [598. Range Addition II](https://leetcode.com/problems/range-addition-ii/description/)

### Description

Given an m * n matrix **M** initialized with all **0**'s and several update operations.

Operations are represented by a 2D array, and each operation is represented by an array with two **positive** integers **a** and **b**, which means **M\[i][j]** should be **added by one** for all **0 <= i < a** and **0 <= j < b**.

You need to count and return the number of maximum integers in the matrix after performing all the operations.

**Example 1:**

```
Input: 
m = 3, n = 3
operations = [[2,2],[3,3]]
Output: 4
Explanation: 
Initially, M = 
[[0, 0, 0],
 [0, 0, 0],
 [0, 0, 0]]

After performing [2,2], M = 
[[1, 1, 0],
 [1, 1, 0],
 [0, 0, 0]]

After performing [3,3], M = 
[[2, 2, 1],
 [2, 2, 1],
 [1, 1, 1]]

So the maximum integer in M is 2, and there are four of it in M. So return 4.
```

**Note:**

1. The range of m and n is [1,40000].
2. The range of a is [1,m], and the range of b is [1,n].
3. The range of operations size won't exceed 10,000.



**Difficult:** `Easy`

**Tags:** `Math`



### Solution One

由于每次 +1 都是 `M[i][j]` （0 <= i < a）（0 <= j < b），即所加范围是左上顶点一定为 (0,0) 的矩形

所以我们只需要求出所有矩形的交集即可

```c++
class Solution {
public:
    int maxCount(int m, int n, vector<vector<int>>& ops) {
        int a = m;
        int b = n;
        for (auto &op : ops)
        {
            a = min(a, op[0]);
            b = min(b, op[1]);
        }
        return a * b;
    }
};
```



### Other Solutions

[598. Range Addition II - Solution](https://leetcode.com/problems/range-addition-ii/solution/)