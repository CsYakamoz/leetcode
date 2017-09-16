## [623. Add One Row to Tree](https://leetcode.com/problems/add-one-row-to-tree/description/)

### Description

Given the root of a binary tree, then value `v` and depth `d`, you need to add a row of nodes with value `v` at the given depth `d`. The root node is at depth 1.

The adding rule is: given a positive integer depth `d`, for each NOT null tree nodes `N` in depth `d-1`, create two tree nodes with value `v` as `N's` left subtree root and right subtree root. And `N's`**original left subtree** should be the left subtree of the new left subtree root, its **original right subtree** should be the right subtree of the new right subtree root. If depth `d` is 1 that means there is no depth d-1 at all, then create a tree node with value **v** as the new root of the whole original tree, and the original tree is the new root's left subtree.

**Example 1:**

```
Input: 
A binary tree as following:
       4
     /   \
    2     6
   / \   / 
  3   1 5   

v = 1

d = 2

Output: 
       4
      / \
     1   1
    /     \
   2       6
  / \     / 
 3   1   5   


```

**Example 2:**

```
Input: 
A binary tree as following:
      4
     /   
    2    
   / \   
  3   1    

v = 1

d = 3

Output: 
      4
     /   
    2
   / \    
  1   1
 /     \  
3       1

```

**Note:**

1. The given d is in range [1, maximum depth of the given tree + 1].
2. The given binary tree has at least one tree node.



**Difficult:** `Medium`

**Tags:** `Tree`



### Solution One

`BFS`

```c++
class Solution {
public:
    TreeNode* addOneRow(TreeNode* root, int v, int d) {
        if (d == 1)
        {
            TreeNode *newRoot = new TreeNode(v);
            newRoot->left = root;
            return newRoot;
        }

        int level = 2;
        queue<TreeNode*> q;
        q.push(root);
        while (level != d)
        {
            size_t length = q.size();
            for (size_t i = 0; i < length; i++)
            {
                TreeNode *node = q.front();
                q.pop();
                if (node->left)
                    q.push(node->left);
                if (node->right)
                    q.push(node->right);
            }
            level++;
        }
        
        size_t length = q.size();
        for (size_t i = 0; i < length; i++)
        {
            TreeNode *node = q.front();
            q.pop();
            TreeNode *lSubtree = new TreeNode(v);
            TreeNode *rSubtree = new TreeNode(v);
            lSubtree->left = node->left;
            rSubtree->right = node->right;
            node->left = lSubtree;
            node->right = rSubtree;
        }

        return root;
    }
};
```



### Solution Two

`DFS`

```c++
class Solution {
public:
    TreeNode* addOneRow(TreeNode* root, int v, int d) {
        if (d == 1)
        {
            TreeNode *newRoot = new TreeNode(v);
            newRoot->left = root;
            return newRoot;
        }
        addOneRow(root, v, d, 1);
        return root;
    }

private:
    void addOneRow(TreeNode *root, int v, int d, int level)
    {
        if (root == nullptr)
            return;

        if (level == d - 1)
        {
            TreeNode *lSubtree = new TreeNode(v);
            TreeNode *rSubtree = new TreeNode(v);
            lSubtree->left = root->left;
            rSubtree->right = root->right;
            root->left = lSubtree;
            root->right = rSubtree;
            return;
        }
        addOneRow(root->left, v, d, level + 1);
        addOneRow(root->right, v, d, level + 1);
    }
};
```



