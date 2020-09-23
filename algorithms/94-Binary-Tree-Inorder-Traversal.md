## [94. Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/#/description)

### Description

Given a binary tree, return the *inorder* traversal of its nodes' values.

For example:
Given binary tree `[1,null,2,3]`,

```
   1
    \
     2
    /
   3
```

return `[1,3,2]`.

**Note:** Recursive solution is trivial, could you do it iteratively?



**Difficult:** `Medium`

**Tags:** `Tree` `Hash Table` `Stack`



### Solution One

`DFS` `Recursion`

```c++
class Solution {
public:
    vector<int> inorderTraversal(TreeNode* root) {
        helper(root);
        return res;
    }

private:
    vector<int> res;
    void helper(TreeNode *root) {
        if (root) {
            helper(root->left);
            res.push_back(root->val);
            helper(root->right);
        }
    }
};
```



### Solution Two

`DFS` `Stack`

[Tree traversal](https://en.wikipedia.org/wiki/Tree_traversal#In-order_2)

```c++
class Solution {
public:
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> res;
        stack<TreeNode*> s;
        while (!s.empty() || root) {
            if (root != nullptr) {
                s.push(root);
                root = root->left;
            } else {
                root = s.top();
                s.pop();
                res.push_back(root->val);
                root = root->right;
            }
        }
        return res;
    }
};
```


