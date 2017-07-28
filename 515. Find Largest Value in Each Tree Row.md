## [515. Find Largest Value in Each Tree Row](https://leetcode.com/problems/find-largest-value-in-each-tree-row/#/description)

### Description

You need to find the largest value in each row of a binary tree.

**Example:**

```
Input: 

          1
         / \
        3   2
       / \   \  
      5   3   9 

Output: [1, 3, 9]
```



**Difficult:** `Medium`

**Tags:** `Tree` `Depth-first Search` `Breadth-first Search`



### Solution One

`DFS` `Recursion`

从左到右进行 DFS

若递归到新的一行，则添加到`result`末尾，否则比较当前结点值与`result[row]`

```c++
struct TreeNode
{
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution
{
  public:
    vector<int> largestValues(TreeNode *root)
    {
        vector<int> result;
        getLargestValues(root, 0, result);
        return result;
    }

  private:
    void getLargestValues(TreeNode *root, size_t row, vector<int> &result)
    {
        if (root == nullptr)
        {
            return;
        }
        if (row != result.size())
        {
            result[row] = root->val > result[row] ? root->val : result[row];
        }
        else
        {
            result.push_back(root->val);
        }
        getLargestValues(root->left, row + 1, result);
        getLargestValues(root->right, row + 1, result);
    }
};
```



### Solution Two

`BFS` `Queue`

```c++
struct TreeNode
{
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

struct node_row
{
    TreeNode *node;
    size_t row;
    node_row(TreeNode *node, size_t x) :node(node), row(x) {}
};

class Solution
{
  public:
    vector<int> largestValues(TreeNode *root)
    {
        vector<int> result;
        if (root == nullptr)
        {
            return result;
        }
        queue<node_row> unvisited;
        node_row current(root, 0);
        unvisited.push(current);
        while (!unvisited.empty())
        {
            current = unvisited.front();
            if (current.row != result.size())
            {
                result[current.row] = current.node->val > result[current.row] ? 
                    current.node->val : result[current.row];
            }
            else
            {
                result.push_back(current.node->val);
            }
            if (current.node->left != nullptr)
            {
                node_row tmp(current.node->left, current.row + 1);
                unvisited.push(tmp);
            }
            if (current.node->right != nullptr)
            {
                node_row tmp(current.node->right, current.row + 1);
                unvisited.push(tmp);
            }
            unvisited.pop();
        }
        return result;
    }
};
```



