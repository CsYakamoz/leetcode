## [637. Average of Levels in Binary Tree](https://leetcode.com/problems/average-of-levels-in-binary-tree/#/description)

### Description

Given a non-empty binary tree, return the average value of the nodes on each level in the form of an array.

**Example 1:**

```
Input:
    3
   / \
  9  20
    /  \
   15   7
Output: [3, 14.5, 11]
Explanation:
The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11. Hence return [3, 14.5, 11].

```

**Note:**

1. The range of node's value is in the range of 32-bit signed integer.



**Difficult:** `Easy`

**Tags:** `Tree`



### Solution One

`BFS` `Queue`

```c++
class Solution {
public:
    vector<double> averageOfLevels(TreeNode* root) {
        vector<double> res;
        if (!root) {
            return res;
        }
        queue<TreeNode*> q;
        TreeNode *current;
        q.push(root);
        while (!q.empty()) {
            double length = q.size();
            double sum = 0;
            for (double i = 0; i < length; i++) {
                current = q.front();
                sum += current->val;
                if (current->left) {
                    q.push(current->left);
                }
                if (current->right) {
                    q.push(current->right);
                }
                q.pop();
            }
            res.push_back(sum / length);
        }
        return res;
    }
};
```



