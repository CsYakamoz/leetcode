## [98. Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/description/)

### Description

Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

- The left subtree of a node contains only nodes with keys **less than** the node's key.
- The right subtree of a node contains only nodes with keys **greater than** the node's key.
- Both the left and right subtrees must also be binary search trees.

**Example 1:**

```
    2
   / \
  1   3

```

Binary tree [2,1,3], return true.

**Example 2:**

```
    1
   / \
  2   3

```

Binary tree [1,2,3], return false.

**Difficult:** `Medium`

**Tags:** `Tree` `Depth-first Search`

### Solution One

来自于书 ——《算法》

```c++
class Solution {
public:
    bool isValidBST(TreeNode* root) {
        return dfs(root, nullptr, nullptr);
    }

    bool dfs(TreeNode *root, int *min, int *max)
    {
        if (root == nullptr)
            return true;

        if (min != nullptr && root->val <= *min)
            return false;
        if (max != nullptr && root->val >= *max)
            return false;

        return dfs(root->left, min, &root->val) && dfs(root->right, &root->val, max);
    }
};
```
