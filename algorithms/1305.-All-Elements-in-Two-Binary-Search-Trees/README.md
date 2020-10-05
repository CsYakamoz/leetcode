## [1305. All Elements in Two Binary Search Trees](https://leetcode.com/problems/all-elements-in-two-binary-search-trees/)

### Description

Given two binary search trees `root1` and `root2`.

Return a list containing _all the integers_ from _both trees_ sorted in **ascending** order.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2019/12/18/q2-e1.png)

```
Input: root1 = [2,1,4], root2 = [1,0,3]
Output: [0,1,1,2,3,4]
```

**Example 2:**

```
Input: root1 = [0,-10,10], root2 = [5,1,7,0,2]
Output: [-10,0,0,1,2,5,7,10]
```

**Example 3:**

```
Input: root1 = [], root2 = [5,1,7,0,2]
Output: [0,1,2,5,7]
```

**Example 4:**

```
Input: root1 = [0,-10,10], root2 = []
Output: [-10,0,10]
```

**Example 5:**

![img](https://assets.leetcode.com/uploads/2019/12/18/q2-e5-.png)

```
Input: root1 = [1,null,8], root2 = [8,1]
Output: [1,1,8,8]
```

**Constraints:**

- Each tree has at most `5000` nodes.
- Each node's value is between `[-10^5, 10^5]`.

**Difficult:** `Medium`

**Tags:** `Sort` `Tree`

### Solution One

由于二叉搜索树的中序遍历，就是从小到大。

所以思路是生成树的中序遍历的迭代器。

```javascript
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */

/**
 * @param {TreeNode} root
 * @returns {Iterator<number, number>}
 */
const getIterator = (root) => {
  /** @type {TreeNode[]} */
  const stack = [];

  return {
    next: () => {
      while (root !== null || stack.length > 0) {
        if (root !== null) {
          stack.push(root);
          root = root.left;
        } else {
          const curr = stack.pop();
          root = curr.right;
          return { value: curr.val, done: false };
        }
      }

      return { done: true };
    },
  };
};

/**
 * @param {TreeNode} root1
 * @param {TreeNode} root2
 * @return {number[]}
 */
const getAllElements = function (root1, root2) {
  const result = [];

  const iter1 = getIterator(root1);
  const iter2 = getIterator(root2);

  let next1 = iter1.next();
  let next2 = iter2.next();

  while (!next1.done || !next2.done) {
    if (!next1.done && !next2.done) {
      if (next1.value < next2.value) {
        result.push(next1.value);
        next1 = iter1.next();
      } else {
        result.push(next2.value);
        next2 = iter2.next();
      }
    } else if (next1.done) {
      result.push(next2.value);
      next2 = iter2.next();
    } else {
      result.push(next1.value);
      next1 = iter1.next();
    }
  }

  return result;
};
```

### Solutions

[1305. All Elements in Two Binary Search Trees - Solution](https://leetcode.com/problems/all-elements-in-two-binary-search-trees/solution/)
