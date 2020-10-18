## [114. Flatten Binary Tree to Linked List](https://leetcode.com/problems/flatten-binary-tree-to-linked-list/description/)

### Description

Given a binary tree, flatten it to a linked list in-place.

For example,
Given

```
         1
        / \
       2   5
      / \   \
     3   4   6
```

The flattened tree should look like:

```
   1
    \
     2
      \
       3
        \
         4
          \
           5
            \
             6
```

**Hints:**

If you notice carefully in the flattened tree, each node's right child points to the next node of a pre-order traversal.

**Difficulty:** `Medium`

**Tags:** `Tree` `Depth-first Search`

### Solution One

```c++
class Solution {
public:
    void flatten(TreeNode* root) {
        if (root == nullptr)
            return;

        stack<TreeNode*> s;
        s.push(root);
        while (!s.empty())
        {
            TreeNode *node = s.top();
            s.pop();
            if (node->left || node->right)
            {
                if (node->right)
                    s.push(node->right);
                if (node->left)
                {
                    s.push(node->left);
                    node->right = node->left;
                    node->left = nullptr;
                }
            }
            else
            {
                if (!s.empty())
                    node->right = s.top();
            }
        }
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    void flatten(TreeNode* root) {
        if (!root) {
            return;
        }
        TreeNode* t = root->right;
        flatten(root->left);
        root->right = root->left;
        root->left = NULL;
        while (root->right)
            root = root->right;
        root->right = t;
        flatten(t);
    }
};
```
