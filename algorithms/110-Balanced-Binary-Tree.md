## [110. Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/description/)

### Description

Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of *every* node never differ by more than 1.



**Difficult:** `Easy`

**Tags:** `Tree` `Depth-first Search`



### Solution One

```c++
class Solution {
public:
    bool isBalanced(TreeNode* root) {
        int height;
        return dfs(root, height);
    }

private:
    bool dfs(TreeNode *root, int &height)
    {
        if (root == nullptr)
            return true;

        int l = 0, r = 0;
        if (dfs(root->left, l) == false || dfs(root->right, r) == false)
            return false;

        if (abs(l - r) > 1)
            return false;

        height = max(l, r) + 1;
        return true;
    }
};
```



### Solution Two - In Top Solutions

```c++
class Solution {
public:
    bool isBalanced(TreeNode* root) {
        return helper(root).second;
    }
    pair<int, bool> helper(TreeNode* root){
        if (root == NULL){
            return {0, true};
        }
        auto left = helper(root -> left);
        auto right = helper(root -> right);
        return {max(left.first, right.first) + 1 ,left.second && right.second && abs(left.first - right.first) <= 1};
    }
};
```



