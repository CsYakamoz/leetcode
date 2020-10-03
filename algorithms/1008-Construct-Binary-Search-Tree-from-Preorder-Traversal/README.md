## [1008. Construct Binary Search Tree from Preorder Traversal](https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/)

### Description

Return the root node of a binary **search** tree that matches the given `preorder` traversal.

_(Recall that a binary search tree is a binary tree where for every node, any descendant of `node.left` has a value `<` `node.val`, and any descendant of `node.right` has a value `>` `node.val`. Also recall that a preorder traversal displays the value of the `node` first, then traverses `node.left`, then traverses `node.right`.)_

It's guaranteed that for the given test cases there is always possible to find a binary search tree with the given requirements.

**Example 1:**

```
Input: [8,5,1,7,10,12]
Output: [8,5,10,1,7,null,12]
```

![](https://assets.leetcode.com/uploads/2019/03/06/1266.png)

**Constraints:**

- `1 <= preorder.length <= 100`
- `1 <= preorder[i] <= 10^8`
- The values of `preorder` are distinct.

**Difficult:** `Medium`

**Tags:** `Tree`

### Solution One

1. 若当前数组长度为 0, 则返回 `null`, 否则进行下一步；

2. 由于是前序遍历，且树是二叉搜索树，故首元素必定是根结点，且右子树的值必定大于根结点的值；

   因此在 `[1, preorder.length)` 中寻找第一个大于根结点的值的下标；

   此时，`[1, idx)` 为左子树，`[idx, preorder.length)` 为右子树；

   注意：1 <= idx <= preorder.length:

   - 当 preorder.length 等于 1 时，左右子树皆为空；
   - 当 idx 等于 preorder.length 时，右子树为空；
   - 当 idx 等于 1 时，左子树为空；

3) 对子树再次进行 1, 2 步骤；

```javascript
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */

/**
 * @param {number[]} preorder
 * @return {TreeNode}
 */
var bstFromPreorder = function (preorder) {
  if (preorder.length === 0) {
    return null;
  }

  const rootVal = preorder[0];
  const root = new TreeNode(rootVal);

  /*
   * find the first element that is greater than rootVal
   * after the loop
   * end will point it if exist, else point the preorder.length
   */
  let begin = 1;
  let end = preorder.length;
  while (begin < end) {
    const mid = begin + Math.floor((end - begin) / 2);
    const val = preorder[mid];

    if (rootVal < val) {
      end = mid;
    } else {
      begin = mid + 1;
    }
  }

  root.left = bstFromPreorder(preorder.slice(1, end));
  root.right = bstFromPreorder(preorder.slice(end, preorder.length));

  return root;
};
```

### Solution Two - In Top Solutions

> Give the function a bound the maximum number it will handle.
>
> The left recursion will take the elements smaller than node.val
>
> The right recursion will take the remaining elements smaller than bound

[原代码链接](<https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/discuss/252232/JavaC%2B%2BPython-O(N)-Solution>)

此处记录个人理解：

`build` 函数作用：返回以 `A[i]` 为根结点，且其子树（包括左右）节点的值从下标范围 `[i + 1, j)`, 其中 `j` 指向第一个大于 `bound` 的元素；

```c++
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    TreeNode* bstFromPreorder(vector<int>& A) {
        int i = 0;
        return build(A, i, INT_MAX);
    }

    TreeNode* build(vector<int>& A, int& i, int bound) {
        if (i == A.size() || A[i] > bound) return NULL;
        TreeNode* root = new TreeNode(A[i++]);
        root->left = build(A, i, root->val);
        root->right = build(A, i, bound);
        return root;
    }
};
```

### Solutions

[1008. Construct Binary Search Tree from Preorder Traversal - Solution](https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/solution/)
