## [113. Path Sum II](https://leetcode.com/problems/path-sum-ii/#/description)

### Description

Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.

For example:
Given the below binary tree and `sum = 22`,

```
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
```

return

```
[
   [5,4,11,2],
   [5,8,4,5]
]
```

**Difficulty:** `Medium`

**Tags:** `Depth-first Search`

### Solution One

```c++
class Solution {
public:
    vector<vector<int>> pathSum(TreeNode* root, int sum) {
        vector<int> record;
        if (root)
        {
            helper(root, 0, record, sum);
        }
        return res;
    }

private:
    vector<vector<int>> res;

    void helper(TreeNode *root, int sum, vector<int> &record, const int &target)
    {
        sum += root->val;
        record.push_back(root->val);
        if (root->left || root->right)
        {
            if (root->left)
            {
                helper(root->left, sum, record, target);
            }
            if (root->right)
            {
                helper(root->right, sum, record, target);
            }
        }
        else
        {
            if (sum == target)
            {
                res.push_back(record);
            }
        }
        record.pop_back();
    }
};
```
