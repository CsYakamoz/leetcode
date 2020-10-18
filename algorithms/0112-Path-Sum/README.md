## [112. Path Sum](https://leetcode.com/problems/path-sum/#/description)

### Description

Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.

For example:

Given the below binary tree and `sum = 22`,

```
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
```

return true, as there exist a root-to-leaf path `5->4->11->2` which sum is 22.

**Difficulty:** `Easy`

**Tags:**`Tree` `Depth-first Search`

### Solution One

`DFS` `Recursion`

```c++
class Solution {
public:
    bool hasPathSum(TreeNode* root, int sum) {
        if (root)
        {
            sum -= root->val;
            if (root->left == nullptr && root->right == nullptr) return sum == 0;
            return hasPathSum(root->left, sum) || hasPathSum(root->right, sum);
        }
        else
        {
            return false;
        }
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    bool DFS(TreeNode* root, int sum, int value) {
        if (root==nullptr) return false;
        value += root->val;
        if (root->left == nullptr && root->right == nullptr) return (sum==value);
        return DFS(root->left, sum, value) || DFS(root->right, sum, value);
    }

    bool hasPathSum(TreeNode* root, int sum) {
        return DFS(root, sum, 0);
    }
};
```
