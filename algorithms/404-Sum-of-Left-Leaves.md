## [404. Sum of Left Leaves](https://leetcode.com/problems/sum-of-left-leaves/#/description)

### Description

Find the sum of all left leaves in a given binary tree.

**Example:**

```
    3
   / \
  9  20
    /  \
   15   7

There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.
```

**Difficult:** `Easy`

**Tags:** `Tree`

### Solution One

`DFS` `Recursion`

在函数中传递一个字符变量`c`，代表方向是左还是右，然后检测该结点是否为叶结点和方向是否为左

经过测试，假如只有根结点，它不算是`left levae`，所以第 6 行传递字符 'R'

```c++
class Solution {
public:
    int sumOfLeftLeaves(TreeNode* root) {
        if (root)
        {
            helper(root, 'R');
        }
        return sum;
    }

private:
    int sum = 0;

    void helper(TreeNode *root, char c)
    {
        if (root->left || root->right)
        {	// If root is not leave node
            if (root->left)
            {
                helper(root->left, 'L');
            }
            if (root->right)
            {
                helper(root->right, 'R');
            }
        }
        else if (c == 'L')
        {	// root is left leave node
            sum += root->val;
        }
    }
};
```

### Solution Two

`BFS` `Queue`

```c++
class Solution {
public:
    int sumOfLeftLeaves(TreeNode* root) {
        int sum = 0;
        queue<pair<TreeNode*, char>> q;
        if (root)
        {
            q.push(pair<TreeNode*, char>{root, 'R'});
        }
        while (!q.empty())
        {
            auto p = q.front();
            auto i = p.first;
            if (i->left || i->right)
            {	// It is not leave node
                if (i->left) q.push(pair<TreeNode*, char>{i->left, 'L'});
                if (i->right) q.push(pair<TreeNode*, char>{i->right, 'R'});
            }
            else if (p.second == 'L')
            {	// It is left leave node
                sum += i->val;
            }
            q.pop();
        }
        return sum;
    }
};
```

### Solution Three - In Top Solutions

[Java iterative and recursive solutions](https://discuss.leetcode.com/topic/60403/java-iterative-and-recursive-solutions)
