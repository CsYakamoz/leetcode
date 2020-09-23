## [102. Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)

### Description

Given a binary tree, return the *level order* traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree `[3,9,20,null,null,15,7]`,

```
    3
   / \
  9  20
    /  \
   15   7

```

return its level order traversal as:

```
[
  [3],
  [9,20],
  [15,7]
]
```



**Difficult:** `Medium`

**Tags:** `Tree` `Breadth-first Search`



### Solution One

`BFS` `Queue`

```c++
class Solution {
public:
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector<vector<int>> res;
        if (root == nullptr) {
            return res;
        }
        queue<TreeNode*> q;
        q.push(root);
        TreeNode *current = nullptr;
        while (!q.empty()) {
            vector<int> level;
            size_t i = q.size();
            while (i) {
                current = q.front();
                if (current->left) {
                    q.push(current->left);
                }
                if (current->right) {
                    q.push(current->right);
                }
                level.push_back(current->val);
                q.pop();
                i--;
            }
            res.push_back(level);
        }
        return res;
    }
};
```



### Solution Two:

`DFS` `Recursion`

```c++
class Solution {
public:
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector<vector<int>> res;
        helper(root, 0, res);
        return res;
    }

    void helper(TreeNode *root, size_t level, vector<vector<int>> &res) {
        if (root) {
            if (level == res.size()) {
                res.push_back(vector<int>{root->val});
            } else {
                res[level].push_back(root->val);
            }
            helper(root->left, level + 1, res);
            helper(root->right, level + 1, res);
        }
    }
};
```



