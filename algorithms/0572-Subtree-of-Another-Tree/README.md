## [572. Subtree of Another Tree](https://leetcode.com/problems/subtree-of-another-tree/description/)

### Description

Given two non-empty binary trees **s** and **t**, check whether tree **t** has exactly the same structure and node values with a subtree of **s**. A subtree of **s** is a tree consists of a node in **s** and all of this node's descendants. The tree **s** could also be considered as a subtree of itself.

**Example 1:**
Given tree s:

```
     3
    / \
   4   5
  / \
 1   2

```

Given tree t:

```
   4
  / \
 1   2

```

Return **true**, because t has the same structure and node values with a subtree of s.

**Example 2:**
Given tree s:

```
     3
    / \
   4   5
  / \
 1   2
    /
   0
```

Given tree t:

```
   4
  / \
 1   2

```

Return **false**.

**Difficulty:** `Easy`

**Tags:** `Tree`

### Solution One

```c++
class Solution {
public:
    bool isSubtree(TreeNode* s, TreeNode* t) {
        if (isSameTree(s, t))
            return true;

        if (s->left && isSubtree(s->left, t))
            return true;

        if (s->right && isSubtree(s->right, t))
            return true;

        return false;
    }

private:
    bool isSameTree(TreeNode *s, TreeNode *t)
    {
        if (s == nullptr && t == nullptr)
            return true;

        if (s == nullptr || t == nullptr)
            return false;

        if (s->val != t->val)
            return false;

        return isSameTree(s->left, t->left) && isSameTree(s->right, t->right);
    }
};
```
