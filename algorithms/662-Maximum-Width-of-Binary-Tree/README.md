## [662. Maximum Width of Binary Tree](https://leetcode.com/problems/maximum-width-of-binary-tree/description/)

### Description

Given a binary tree, write a function to get the maximum width of the given tree. The width of a tree is the maximum width among all levels. The binary tree has the same structure as a **full binary tree**, but some nodes are null.

The width of one level is defined as the length between the end-nodes (the leftmost and right most non-null nodes in the level, where the `null` nodes between the end-nodes are also counted into the length calculation.

**Example 1:**

```
Input:

           1
         /   \
        3     2
       / \     \
      5   3     9

Output: 4
Explanation: The maximum width existing in the third level with the length 4 (5,3,null,9).

```

**Example 2:**

```
Input:

          1
         /
        3
       / \
      5   3

Output: 2
Explanation: The maximum width existing in the third level with the length 2 (5,3).

```

**Example 3:**

```
Input:

          1
         / \
        3   2
       /
      5

Output: 2
Explanation: The maximum width existing in the second level with the length 2 (3,2).
```

**Example 4:**

```
Input:

          1
         / \
        3   2
       /     \
      5       9
     /         \
    6           7
Output: 8
Explanation:The maximum width existing in the fourth level with the length 8 (6,null,null,null,null,null,null,7).


```

**Note:** Answer will in the range of 32-bit signed integer.

**Difficult:** `Medium`

**Tags:** `Tree`

### Solution One

```c++
class Solution {
public:
    int widthOfBinaryTree(TreeNode* root) {
        if (root == nullptr)
            return 0;

        int maxWidth = 0;
        queue<pair<TreeNode*, int>> q;
        q.push(make_pair(root, 0));

        while (!q.empty())
        {
            size_t length = q.size();
            vector<int> index;
            for (size_t i = 0; i < length; i++)
            {
                auto node = q.front();
                q.pop();
                index.push_back(node.second);
                if (node.first->left)
                    q.push(make_pair(node.first->left, node.second * 2));
                if (node.first->right)
                    q.push(make_pair(node.first->right, node.second * 2 + 1));
            }
            maxWidth = max(maxWidth, index.back() - index.front() + 1);
        }

        return maxWidth;
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    int widthOfBinaryTree(TreeNode* root) {
        vector<int> lefts;
        return dfs(root, 1, 0, lefts);
    }
    int dfs(TreeNode* root, int idx, int d, vector<int>& lefts) {
        if(root == NULL)
            return 0;
        if(d >= lefts.size())
            lefts.push_back(idx);
        return max({idx - lefts[d] + 1, dfs(root->left, idx*2, d+1, lefts), dfs(root->right, idx*2+1, d+1, lefts)});
    }
};
```
