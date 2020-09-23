## [655. Print Binary Tree](https://leetcode.com/problems/print-binary-tree/description/)

### Description

Print a binary tree in an m\*n 2D string array following these rules:

1. The row number `m` should be equal to the height of the given binary tree.
2. The column number `n` should always be an odd number.
3. The root node's value (in string format) should be put in the exactly middle of the first row it can be put. The column and the row where the root node belongs will separate the rest space into two parts (**left-bottom part and right-bottom part**). You should print the left subtree in the left-bottom part and print the right subtree in the right-bottom part. The left-bottom part and the right-bottom part should have the same size. Even if one subtree is none while the other is not, you don't need to print anything for the none subtree but still need to leave the space as large as that for the other subtree. However, if two subtrees are none, then you don't need to leave space for both of them.
4. Each unused space should contain an empty string `""`.
5. Print the subtrees following the same rules.

**Example 1:**

```
Input:
     1
    /
   2
Output:
[["", "1", ""],
 ["2", "", ""]]

```

**Example 2:**

```
Input:
     1
    / \
   2   3
    \
     4
Output:
[["", "", "", "1", "", "", ""],
 ["", "2", "", "", "", "3", ""],
 ["", "", "4", "", "", "", ""]]
```

**Example 3:**

```
Input:
      1
     / \
    2   5
   /
  3
 /
4
Output:

[["",  "",  "", "",  "", "", "", "1", "",  "",  "",  "",  "", "", ""]
 ["",  "",  "", "2", "", "", "", "",  "",  "",  "",  "5", "", "", ""]
 ["",  "3", "", "",  "", "", "", "",  "",  "",  "",  "",  "", "", ""]
 ["4", "",  "", "",  "", "", "", "",  "",  "",  "",  "",  "", "", ""]]
```

**Note:** The height of binary tree is in the range of [1, 10].

**Difficult:** `Medium`

**Tags:** `Tree`

### Solution One

```c++
class Solution {
public:
    vector<vector<string>> printTree(TreeNode* root) {
        int height = getHeight(root);
        int width = pow(2, height) - 1;
        vector<vector<string>> res(height, vector<string>(width));
        dfs(root, 0, 0, width - 1, res);
        return res;
    }

private:
    int getHeight(TreeNode *root)
    {
        if (root == nullptr)
            return 0;
        queue<TreeNode*> q;
        q.push(root);
        int height = 0;
        while (!q.empty())
        {
            height++;
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
        }
        return height;
    }

    void dfs(TreeNode *root, int height, int left, int right, vector<vector<string>> &res)
    {
        if (root == nullptr)
            return;

        int mid = left + (right - left) / 2;
        res[height][mid] = to_string(root->val);
        dfs(root->left, height + 1, left, mid - 1, res);
        dfs(root->right, height + 1, mid + 1, right, res);
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    int getHeight(TreeNode *root)
    {
	    if (root == nullptr) return 0;
	    return max(getHeight(root->left), getHeight(root->right)) + 1;
    }

void setValue(TreeNode *root, vector<vector<string>> &rst, int s, int e,int h)
{
	if (root == nullptr) return;
	int mid = (s + e) / 2;
	rst[h][mid] = to_string(root->val);
	setValue(root->left, rst, s, mid-1, h + 1);
	setValue(root->right, rst, mid+1, e, h + 1);
}

vector<vector<string>> printTree(TreeNode* root) {
	int h = getHeight(root);
	int col = pow(2, h) - 1;
	vector<vector<string>> rst(h, vector<string>(col, ""));
    setValue(root,rst,0,col-1,0);
	return rst;
    }
};
```
