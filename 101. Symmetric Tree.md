## [101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/#/description)

### Description

Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree `[1,2,2,3,4,4,3]` is symmetric:

```
    1
   / \
  2   2
 / \ / \
3  4 4  3

```

But the following `[1,2,2,null,3,null,3]` is not:

```
    1
   / \
  2   2
   \   \
   3    3

```

**Note:**
Bonus points if you could solve it both recursively and iteratively.



**Difficult:** `Easy`

**Tags:** `Tree` `Depth-first Search` `Breadth-first Search`



### Solution One - In Top Solutions

`DFS` `Recursion`

```c++
class Solution {
public:
    bool isSymmetric(TreeNode* root) {
        if (!root)
        {
            return true;
        }
        return helper(root->left, root->right);
    }

private:
    bool helper(TreeNode *l, TreeNode *r) {
        if (l == nullptr && r == nullptr)
        {	// Both of them are nullptr
            return true;
        }
        else if (l == nullptr || r == nullptr)
        {	// One of them is nullptr
            return false;
        }
        else if (l->val != r->val)
        {
            return false;
        }
        return helper(l->left, r->right) && helper(l->right, r->left);
    }
};
```



### Solution Two - In Top Solutions

`BFS` `Queue`

[My C++ Accepted code in 16ms with iteration solution](https://discuss.leetcode.com/topic/4332/my-c-accepted-code-in-16ms-with-iteration-solution)

```c++
class Solution {
public:
    bool isSymmetric(TreeNode* root) {
        queue<TreeNode*> q1;
        queue<TreeNode*> q2;
        if (root)
        {
            q1.push(root->left);
            q2.push(root->right);
        }
        // The Top Solutions use "!q1.empty() && !q2.empty()" as condition
        // But I think using "!q1.empty()" or "!q2.empty()" is ok
        while (!q1.empty())
        {
            TreeNode *node1 = q1.front();
            q1.pop();
            TreeNode *node2 = q2.front();
            q2.pop();
            if (node1 == nullptr && node2 == nullptr)
            {	// Both of them are nullptr
                continue;
            }
            else if (node1 == nullptr || node2 == nullptr)
            {	// One of them is nullptr
                return false;
            }
            else if (node1->val != node2->val)
            {
                return false;
            }
            q1.push(node1->left);
            q1.push(node1->right);
            q2.push(node2->right);
            q2.push(node2->left);
        }
        return true;
    }
};
```



