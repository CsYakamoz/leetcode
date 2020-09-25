## [105. Construct Binary Tree from Preorder and Inorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/)

### Description

Given preorder and inorder traversal of a tree, construct the binary tree.

**Note:**
You may assume that duplicates do not exist in the tree.

**Difficult:** `Medium`

**Tags:** `Tree` `Array` `Depth-first Search`

### Solution One

```c++
class Solution {
public:
    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        if (preorder.empty())
            return nullptr;

        for (size_t i = 0; i < inorder.size(); i++)
            hash[inorder[i]] = i;

        int i = 0;
        return helper(i, 0, inorder.size() - 1, preorder, inorder);
    }

private:
    map<int, int> hash;

    TreeNode* helper(int &i, int lo, int hi, vector<int> &preorder, vector<int> &inorder)
    {
        TreeNode *root = new TreeNode(preorder[i]);

        int index = hash[preorder[i]];

        if (lo < index)
            root->left = helper(++i, lo, index - 1, preorder, inorder);

        if (index < hi)
            root->right = helper(++i, index + 1, hi, preorder, inorder);

        return root;
    }
};
```
