## [226. Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/#/description)

### Description

Invert a binary tree.

```
     4
   /   \
  2     7
 / \   / \
1   3 6   9
```

to

```
     4
   /   \
  7     2
 / \   / \
9   6 3   1
```

**Difficulty:** `Easy`

**Tags:** `Tree`

### Solution One

从树的第一层开始交换左右结点

```c++
class Solution {
public:
    TreeNode* invertTree(TreeNode* root) {
        helper(root);
        return root;
    }

private:
    void helper(TreeNode *root)
    {
        if (root)
        {
            if (root->left || root->right)
            {
                swap(root->left, root->right);
                helper(root->left);
                helper(root->right);
            }
        }
    }
};
```

### Solution Two - In Top Solutions

从树的最低层开始交换左右结点

```c++
class Solution {
public:
    TreeNode* invertTree(TreeNode* root) {
        return invertTreeHelper(root);
    }
    TreeNode* invertTreeHelper(TreeNode* node) {
        if(node == NULL)return NULL;
        invertTreeHelper(node->left);
        invertTreeHelper(node->right);
        swap(node->left, node->right);
        return node;
    }
};
```
