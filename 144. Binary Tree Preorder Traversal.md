## [144. Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/description/)

### Description

Given a binary tree, return the *preorder* traversal of its nodes' values.

For example:
Given binary tree `{1,#,2,3}`,

```
   1
    \
     2
    /
   3
```

return `[1,2,3]`.

**Note:** Recursive solution is trivial, could you do it iteratively?



**Difficult:** `Medium`

**Tags:** `Tree` `Stack`



### Solution One

```c++
class Solution {
public:
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> res;
        if (root == nullptr)
        {
            return res;
        }
        stack<TreeNode*> s;
        s.push(root);
        while (!s.empty())
        {
            auto node = s.top();
            s.pop();
            res.push_back(node->val);
            if (node->right)
            {
                s.push(node->right);
            }
            if (node->left)
            {
                s.push(node->left);
            }
        }
        return res;
    }
};
```



