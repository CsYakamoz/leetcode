## [106. Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/)

### Description

Given inorder and postorder traversal of a tree, construct the binary tree.

**Note:**
You may assume that duplicates do not exist in the tree.



**Difficult:** `Medium`

**Tags:** `Tree` `Array` `Depth-first Search`



### Solution One

```c++
class Solution {
public:
    TreeNode* buildTree(vector<int>& inorder, vector<int>& postorder) {
        if (postorder.empty())
            return nullptr;

        for (size_t i = 0; i < postorder.size(); i++)
            hash[inorder[i]] = i;

        int i = postorder.size() - 1;
        return helper(i, 0, inorder.size() - 1, inorder, postorder);
    }

private:
    map<int, int> hash;

    TreeNode* helper(int &i, int lo, int hi, vector<int> &inorder, vector<int> &postorder)
    {
        TreeNode *root = new TreeNode(postorder[i]);

        int index = hash[postorder[i]];

        if (index < hi)
            root->right = helper(--i, index + 1, hi, inorder, postorder);

        if (lo < index)
            root->left = helper(--i, lo, index - 1, inorder, postorder);

        return root;
    }
};
```

