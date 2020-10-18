## [116. Populating Next Right Pointers in Each Node](https://leetcode.com/problems/populating-next-right-pointers-in-each-node/description/)

### Description

Given a binary tree

```
    struct TreeLinkNode {
      TreeLinkNode *left;
      TreeLinkNode *right;
      TreeLinkNode *next;
    }

```

Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to `NULL`.

Initially, all next pointers are set to `NULL`.

**Note:**

- You may only use constant extra space.
- You may assume that it is a perfect binary tree (ie, all leaves are at the same level, and every parent has two children).

For example,
Given the following perfect binary tree,

```
         1
       /  \
      2    3
     / \  / \
    4  5  6  7

```

After calling your function, the tree should look like:

```
         1 -> NULL
       /  \
      2 -> 3 -> NULL
     / \  / \
    4->5->6->7 -> NULL
```

**Difficulty:** `Medium`

**Tags:** `Tree` `Depth-first Search`

### Solution One - In Top Solutions

[A simple accepted solution](https://discuss.leetcode.com/topic/2202/a-simple-accepted-solution)

```c++
class Solution {
public:
    void connect(TreeLinkNode *root) {
        if (root == nullptr) return;

        TreeLinkNode *pre = root;
        while (pre->left)
        {
            TreeLinkNode *node = pre;
            while (node)
            {
                node->left->next = node->right;
                if (node->next) node->right->next = node->next->left;
                node = node->next;
            }
            pre = pre->left;
        }
    }
};
```

### Solution Two - In Top Solutions

[My recursive solution(Java)](https://discuss.leetcode.com/topic/12241/my-recursive-solution-java)

```c++
class Solution {
public:
    void connect(TreeLinkNode *root) {
        if (root == nullptr || root->left == nullptr)
            return;

        root->left->next = root->right;
        if (root->next != nullptr)
            root->right->next = root->next->left;

        connect(root->left);
        connect(root->right);
    }
};
```

### Solution Three

在刚看到题目时，我就有了下面的思路，但后面重新看题目时，发现不满足条件 - `You may only use constant extra space.`，同时也想不出 Top 的思路~~~

空间复杂度为 O(n)，n 为树的深度

类似方法还有`BFS`，此思路的空间复杂度为 O(n)，n 为树的宽度（最宽的一层）

**注意：** 此思路可被 AC

```c++
class Solution {
public:
    void connect(TreeLinkNode *root) {
        vector<TreeLinkNode*> level;
        helper(root, 0, level);
    }

private:
    void helper(TreeLinkNode *root, int row, vector<TreeLinkNode*> &level)
    {
        if (root)
        {
            if (row == level.size())
            {
                level.push_back(root);
            }
            else
            {
                level[row]->next = root;
                level[row] = root;
            }
            helper(root->left, row + 1, level);
            helper(root->right, row + 1, level);
        }
    }
};
```
