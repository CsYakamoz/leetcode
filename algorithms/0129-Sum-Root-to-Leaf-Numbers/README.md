## [129. Sum Root to Leaf Numbers](https://leetcode.com/problems/sum-root-to-leaf-numbers/#/description)

Given a binary tree containing digits from `0-9` only, each root-to-leaf path could represent a number.

An example is the root-to-leaf path `1->2->3` which represents the number `123`.

Find the total sum of all root-to-leaf numbers.

For example,

```
    1
   / \
  2   3

```

The root-to-leaf path `1->2` represents the number `12`.
The root-to-leaf path `1->3` represents the number `13`.

Return the sum = 12 + 13 = `25`.

**Difficulty:** `Medium`

**Tags:** `Tree` `Depth-first Search`

### Solution One

```c++
class Solution {
public:
    int sumNumbers(TreeNode* root) {
        int num = 0;
        helper(root, num);
        return res;
    }

private:
    int res = 0;

    void helper(TreeNode *root, int num)
    {
        if (root == nullptr) return;
        num = num * 10 + root->val;
        if (root->left == nullptr && root->right == nullptr)
        {
            res += num;
            return;
        }
        helper(root->left, num);
        helper(root->right, num);
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
    int sum = 0;

    void dfs(int val, TreeNode* p) {
        val = val * 10 + p -> val;
        if (!p -> left && !p -> right) {
            sum += val;
            return;
        }
        if (p -> left) dfs(val, p -> left);
        if (p -> right) dfs(val, p -> right);
    }

public:
    int sumNumbers(TreeNode* root) {
        if (!root) return 0;
        dfs(0, root);
        return sum;
    }
};
```
