## [450. Delete Node in a BST](https://leetcode.com/problems/delete-node-in-a-bst/description/)

### Description

Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:

1. Search for a node to remove.
2. If the node is found, delete the node.

**Note:** Time complexity should be O(height of tree).

**Example:**

```
root = [5,3,6,2,4,null,7]
key = 3

    5
   / \
  3   6
 / \   \
2   4   7

Given key to delete is 3. So we find the node with value 3 and delete it.

One valid answer is [5,4,6,2,null,null,7], shown in the following BST.

    5
   / \
  4   6
 /     \
2       7

Another valid answer is [5,2,6,null,4,null,7].

    5
   / \
  2   6
   \   \
    4   7
```

**Difficulty:** `Medium`

**Tags:** `Tree`

### Solution One

```c++
class Solution {
public:
    TreeNode* deleteNode(TreeNode* root, int key) {
        if (root == nullptr)
            return nullptr;

        if (key < root->val)
        {
            root->left = deleteNode(root->left, key);
        }
        else if (key > root->val)
        {
            root->right = deleteNode(root->right, key);
        }
        else
        {
            if (root->right == nullptr)
                return root->left;
            if (root->left == nullptr)
                return root->right;

            TreeNode *node = root;
            root = min(node->right);
            root->right = deleteMin(node->right);
            root->left = node->left;
        }
        return root;
    }

private:
    TreeNode* min(TreeNode *root)
    {
        while (root->left)
            root = root->left;
        return root;
    }

    TreeNode* deleteMin(TreeNode *root)
    {
        if (root->left == nullptr)
            return root->right;

        root->left = deleteMin(root->left);
        return root;
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    TreeNode* deleteNode(TreeNode* root, int key) {
        if(!root) return root;
        if(root->val >key)
        {
            root->left = deleteNode(root->left,key);
        }
        else if(root->val<key)
        {
            root->right = deleteNode(root->right,key);
        }
        else
        {
            if(!root->left || !root->right)
                return root->left?root->left:root->right;
            TreeNode *pre = root->right;
            while(pre->left)
            {
                pre = pre->left;
            }
            root->val = pre->val;
            root->right = deleteNode(root->right,root->val);
        }
        return root;
    }
};
```
