## [104. Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/#/description)

### Description

Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

**Difficult:** `Easy`

**Tags:** `Tree` `Depth-first Search`

### Solution One

`DFS` `Recursion`

```c++
class Solution {
public:
    int maxDepth(TreeNode* root) {
        return helper(root, 0);
    }

private:
    int helper(TreeNode *root, int row) {
        if (root != nullptr) {
            int l = helper(root->left, row + 1);
            int r = helper(root->right, row + 1);
            return max(l, r);
        } else {
            return row;
        }
    }
};
```

### Solution Two:

`BFS` `Queue`

```c++
class Solution {
public:
    int maxDepth(TreeNode* root) {
        int row = 0;
        queue<TreeNode*> q;
        TreeNode *current;
        if (!root) {
            return row;
        }
        q.push(root);
        while (!q.empty()) {
            row++;
            size_t i = q.size();
            while (i) {
                current = q.front();
                if (current->left) {
                    q.push(current->left);
                }
                if (current->right) {
                    q.push(current->right);
                }
                q.pop();
                i--;
            }
        }
        return row;
    }
};
```
