## [96. Unique Binary Search Tree](https://leetcode.com/problems/unique-binary-search-trees/description/)

### Description

Given _n_, how many structurally unique **BST's** (binary search trees) that store values 1..._n_?

For example,
Given _n_ = 3, there are a total of 5 unique BST's.

```
   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
```

**Difficult:** `Medium`

**Tags:** `Dynamic Programming` `Tree`

### Solution One

```c++
class Solution {
public:
    int numTrees(int n) {
        vector<int> dp(n + 1, 1);
        for (int i = 2; i <= n; i++)
        {
            int count = 0;
            int mid = i / 2 + (i % 2);

            for (int j = 1; j <= mid; j++)
                count += dp[j - 1] * dp[i - j];

            dp[i] = count * 2;
            if (i % 2 == 1)
                dp[i] -= dp[mid - 1] * dp[i - mid];
        }

        return dp.back();
    }
};
```
