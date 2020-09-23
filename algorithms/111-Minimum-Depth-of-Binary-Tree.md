## [111. Minimum Depth of Binary Tree](https://leetcode.com/problems/minimum-depth-of-binary-tree/#/description)

### Description

Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.



**Difficult:** `Easy`

**Tags:** `Tree` `Depth-first Search` `Breadth-first Search`



### Solution One:

`DFS` `Recursion`

```c++
class Solution {
public:
    int minDepth(TreeNode* root) {
        if (root) {
            minLevel = INT_MAX;
            helper(root, 1);
        }
        return minLevel;
    }

private:
    int minLevel;

    void helper(TreeNode *root, int row) {
        if (root->left || root->right) {
            if (root->left) {
                helper(root->left, row + 1);
            }
            if (root->right) {
                helper(root->right, row + 1);
            }
        } else {
            minLevel = min(row, minLevel);
        }
    }
};
```



### Solution Two

`BFS` `Queue`

```c++
class Solution {
public:
    int minDepth(TreeNode* root) {
        int minLevel = 0;
        int currentLevel = 0;
        queue<TreeNode*> q;
        if (root) {
            q.push(root);
            minLevel = INT_MAX;
        } else {
            return minLevel;
        }
        while (!q.empty()) {
            currentLevel++;
            size_t length = q.size();
            for (size_t i = 0; i < length; i++) {
                int isLeaf = 0;
                TreeNode *current = q.front();
                q.pop();
                if (current->left) {
                    q.push(current->left);
                    isLeaf++;
                }
                if (current->right) {
                    q.push(current->right);
                    isLeaf++;
                }
                if (!isLeaf) {
                    minLevel = min(currentLevel, minLevel);
                }
            }
        }
        return minLevel;
    }
};
```



### Solution Three - In Top Solutions

```c++
class Solution {
    /*thought: by lvl break when find the first leaf*/
    /*recursive: null is lvl 0; leaf is lvl 1; otherwise is min of subtree+1*/
public:
    int minDepth(TreeNode* root) {
        if (!root) return 0;
        if (!root->left && !root->right) return 1;
        if (!root->left)
            return minDepth(root->right)+1;
        if (!root->right)
            return minDepth(root->left)+1;
        return min(minDepth(root->left), minDepth(root->right))+1;
    }
#if 0    
    int minDepth(TreeNode* root) {
        if (!root) return 0;
        TreeNode *node = root;
        int lvl=0;
        queue<TreeNode *> q;
        q.push(node);
        while(q.size()) {
            int len=q.size();
            lvl++;
            for (int i=0;i<len;i++) {
                node=q.front();
                q.pop();
                if (!node->left && !node->right)
                    return lvl;
                if (node->left) q.push(node->left);
                if (node->right) q.push(node->right);
            }
            
        }
        return lvl;
    }
#endif
};
```



