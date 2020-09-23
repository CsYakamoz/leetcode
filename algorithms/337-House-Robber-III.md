## [337. House Robber III](https://leetcode.com/problems/house-robber-iii/description/)

### Description

The thief has found himself a new place for his thievery again. There is only one entrance to this area, called the "root." Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that "all houses in this place forms a binary tree". It will automatically contact the police if two directly-linked houses were broken into on the same night.

Determine the maximum amount of money the thief can rob tonight without alerting the police.

**Example 1:**

```
     3
    / \
   2   3
    \   \ 
     3   1
```

Maximum amount of money the thief can rob = 3 + 3 + 1 = **7**.

**Example 2:**

```
     3
    / \
   4   5
  / \   \ 
 1   3   1
```

Maximum amount of money the thief can rob = 4 + 5 = **9**.



**Difficult:** `Medium`

**Tags:** `Tree` `Depth-first Search`



### Solution One

```c++
class Solution {
public:
    int rob(TreeNode* root)
    {
        auto res = dfs(root);
        return max(res.first, res.second);
    }
    
private:
    /*
     * first means yes
     * second means no
     **/
    pair<int, int> dfs(TreeNode *root)
    {
        if (root == nullptr)
            return make_pair(0, 0);

        auto left = dfs(root->left);
        auto right = dfs(root->right);
        int y = left.second + right.second + root->val;
        int n = max(left.first, left.second) + max(right.first, right.second);
        return make_pair(y, n);
    }
};
```



### Solution Two - In Top Solutions

[Step by step tackling of the problem](https://discuss.leetcode.com/topic/39834/step-by-step-tackling-of-the-problem)

