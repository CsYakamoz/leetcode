## [257. Binary Tree Paths](https://leetcode.com/problems/binary-tree-paths/#/description)

### Description

Given a binary tree, return all root-to-leaf paths.

For example, given the following binary tree:

```
   1
 /   \
2     3
 \
  5
```

All root-to-leaf paths are:

```
["1->2->5", "1->3"]
```

**Difficult:** `Easy`

**Tags:** `Tree` `Depth-first Search`

### Solution One

`DFS` `Recursion`

```c++
class Solution {
public:
    vector<string> binaryTreePaths(TreeNode* root) {
        string s;
        if (root)
        {
            helper(root, s);
        }
        return res;
    }

private:
    vector<string> res;

    void helper(TreeNode *root, string s)
    {
        s += to_string(root->val);
        if (root->left || root->right)
        {
            s += "->";
            if (root->left) helper(root->left, s);
            if (root->right) helper(root->right, s);
        }
        else
        {	// The root is leaf node
            res.push_back(s);
        }
    }
};
```

### Solution Two - In Top Solutions

`BFS` `Queue`

```c++
class Solution {
public:
    vector<string> binaryTreePaths(TreeNode* root) {
        vector<string> res;
        queue<pair<TreeNode*, string>> q;
        if(!root) return res;
        q.push({root, ""});
        while(!q.empty()) {
            auto cur = q.front(); q.pop();
            if(!cur.first->left && !cur.first->right) {
                string tmp = (cur.second == "" ? to_string(cur.first->val) : cur.second + "->" + to_string(cur.first->val));
                res.push_back(tmp);
            }
            else {
                string tmp = (cur.second == "" ? to_string(cur.first->val) : cur.second + "->" + to_string(cur.first->val));
                if(cur.first->left) {
                    q.push({cur.first->left, tmp});
                }
                if(cur.first->right) {
                    q.push({cur.first->right, tmp});
                }
            }
        }
        return res;
    }
};
```
