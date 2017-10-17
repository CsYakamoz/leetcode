## [145. Binary Tree Postorder Traversal](https://leetcode.com/problems/binary-tree-postorder-traversal/description/)

### Description

Given a binary tree, return the *postorder* traversal of its nodes' values.

For example:
Given binary tree `{1,#,2,3}`,

```
   1
    \
     2
    /
   3

```

return `[3,2,1]`.

**Note:** Recursive solution is trivial, could you do it iteratively?



**Difficult:** `Medium`

**Tags:** `Tree` `Stack`



### Solution One

```c++
class Solution {
public:
    vector<int> postorderTraversal(TreeNode* root) {
        helper(root);
        return res;
    }

private:
    vector<int> res;

    void helper(TreeNode *root)
    {
        if (root)
        {
            helper(root->left);
            helper(root->right);
            res.push_back(root->val);
        }
    }
};
```



### Solution Two - In Top Solutions

[Tree traversal](https://en.wikipedia.org/wiki/Tree_traversal#Post-order)

```c++
class Solution {
public:
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> res;
        if (root == nullptr)
            return res;
        stack<TreeNode*> s;
        TreeNode *lastNodeVisited = nullptr;
        while (!s.empty() || root != nullptr) {
            if (root != nullptr) {
                s.push(root);
                root = root->left;
            } else {
                TreeNode *node = s.top();
                if (node->right != nullptr && lastNodeVisited != node->right)
                    root = node->right;
                else {
                    res.push_back(node->val);
                    lastNodeVisited = node;
                    s.pop();
                }
            }
        }

        return res;
    }
};
```



