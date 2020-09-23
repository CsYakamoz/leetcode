## [543. Diameter of Binary Tree](https://leetcode.com/problems/diameter-of-binary-tree/description/)

### Description

Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the **longest** path between any two nodes in a tree. This path may or may not pass through the root.

**Example:**
Given a binary tree

```
          1
         / \
        2   3
       / \
      4   5

```

Return **3**, which is the length of the path [4,2,1,3] or [5,2,1,3].

**Note:** The length of path between two nodes is represented by the number of edges between them.

**Difficult:** `Easy`

**Tags:** `Tree`

### Solution One

看了 **Two** 后，发现 4 ~ 9 行的确可以直接改为 `helper(root)`

```c++
class Solution {
public:
    int diameterOfBinaryTree(TreeNode* root) {
        if (root)
        {
            int leftDepth = helper(root->left, 0);
            int rightDepth = helper(root->right, 0);
            res = max(res, leftDepth + rightDepth);
        }
        return res;
    }

private:
    int res = 0;

    int helper(TreeNode *root, int len)
    {
        if (root)
        {
            int leftDepth = helper(root->left, 0);
            int rightDepth = helper(root->right, 0);
            res = max(res, leftDepth + rightDepth);
            return max(leftDepth, rightDepth) + 1;
        }
        else
        {
            return len;
        }
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    int ret = 0;
    int diameterOfBinaryTree(TreeNode* root) {
        depth(root);
        return ret;
    }

    int depth(TreeNode* root)
    {
        if (!root) return 0;
        int left = depth(root->left);
        int right = depth(root->right);
        ret = max(ret, left+right);
        return 1+max(left,right);
    }
};
```
