## [199. Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view/description/)

### Description

Given a binary tree, imagine yourself standing on the _right_ side of it, return the values of the nodes you can see ordered from top to bottom.

For example:
Given the following binary tree,

```
   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
```

You should return `[1, 3, 4]`.

**Difficulty:** `Medium`

**Tags:** `Tree` `Depth-first Search` `Breadth-first Search`

### Solution One

```c++
class Solution {
public:
    vector<int> rightSideView(TreeNode* root) {
        vector<int> res;
        queue<TreeNode*> q;
        if (root)
        {
            q.push(root);
        }
        while (!q.empty())
        {
            int length = q.size();
            int val;
            for (size_t i = 0; i < length; i++)
            {
                TreeNode *node = q.front();
                q.pop();
                val = node->val;
                if (node->left)
                {
                    q.push(node->left);
                }
                if (node->right)
                {
                    q.push(node->right);
                }
            }
            res.push_back(val);
        }
        return res;
    }
};
```

### Solution Two

```c++
class Solution {
public:
    vector<int> rightSideView(TreeNode* root) {
        vector<int> res;
        helper(root, 0, res);
        return res;
    }

    void helper(TreeNode *root, int row, vector<int> &res)
    {
        if (root)
        {
            if (row == res.size())
            {
                res.push_back(root->val);
            }
            helper(root->right, row + 1, res);
            helper(root->left, row + 1, res);
        }
    }
};
```
