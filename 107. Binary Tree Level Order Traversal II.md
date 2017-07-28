## [107. Binary Tree Level Order Traversal II](https://leetcode.com/problems/binary-tree-level-order-traversal-ii/#/description)

### Description

Given a binary tree, return the *bottom-up level order* traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree `[3,9,20,null,null,15,7]`,

```
    3
   / \
  9  20
    /  \
   15   7

```

return its bottom-up level order traversal as:

```
[
  [15,7],
  [9,20],
  [3]
]
```



**Difficult:** `Easy`

**Tags:** `Tree` `Breadth-first Search`



### Solution One

`DFS` `Recursion`

```c++
class Solution {
public:
    vector<vector<int>> levelOrderBottom(TreeNode* root) {
        vector<vector<int>> res;
        helper(root, 0, res);
        reverse(res.begin(), res.end());
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



### Solution Two:

`BFS` `Queue`

```c++
class Solution {
public:
    vector<vector<int>> levelOrderBottom(TreeNode* root) {
        vector<vector<int>> res;
        queue<TreeNode*> q;
        if (root) {
            q.push(root);
        } else {
            return res;
        }
        while (!q.empty()) {
            vector<int> level;
            size_t length = q.size();
            for (size_t i = 0; i < length; i++) {
                TreeNode *current = q.front();
                level.push_back(current->val);
                q.pop();
                if (current->left) {
                    q.push(current->left);
                }
                if (current->right) {
                    q.push(current->right);
                }
            }
            res.push_back(level);
        }
        reverse(res.begin(), res.end());
        return res;
    }
};
```



