## [513. Find Bottom Left Tree Value](https://leetcode.com/problems/find-bottom-left-tree-value/#/description)

### Description

Given a binary tree, find the leftmost value in the last row of the tree.

**Example 1:**

```
Input:

    2
   / \
  1   3

Output:
1

```

**Example 2: **

```
Input:

        1
       / \
      2   3
     /   / \
    4   5   6
       /
      7

Output:
7

```

**Note:** You may assume the tree (i.e., the given root node) is not **NULL**.

**Difficult:** `Medium`

**Tags:** `Tree` `Breadth-first Search` `Depth-first Search`

### Solution One

`DFS` `Recursion`

这里使用 DFS 向左子数递归，一旦递归到新的一行，会将该行最左侧的结点的值添加到`result`中

所以最后只需要返回`result.back()`

```c++
class Solution {
public:
    int findBottomLeftValue(TreeNode* root) {
        getValue(root, 0);
        return result.back();
    }
private:
    vector<int> result;

    void getValue(TreeNode *&root, int row)
    {
        if (root != nullptr)
        {
            if (row == result.size())
            {
                result.push_back(root->val);
            }
            getValue(root->left, row + 1);
            getValue(root->right, row + 1);
        }
    }
};
```

### Solution Two - In Top Solutions

`BFS` `Queue`

[Right-to-Left BFS (Python + Java)](https://discuss.leetcode.com/topic/78981/right-to-left-bfs-python-java)

从右向左进行 BFS，

当进行到最后一步时，root 指向最底层，最左边的结点，故此时直接返回`root.val`

```java
public int findLeftMostNode(TreeNode root) {
    Queue<TreeNode> queue = new LinkedList<>();
    queue.add(root);
    while (!queue.isEmpty()) {
        root = queue.poll();
        if (root.right != null)
            queue.add(root.right);
        if (root.left != null)
            queue.add(root.left);
    }
    return root.val;
}
```
