## [437. Path Sum III](https://leetcode.com/problems/path-sum-iii/)

### Description

You are given a binary tree in which each node contains an integer value.

Find the number of paths that sum to a given value.

The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).

The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.

**Example:**

```
root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

Return 3. The paths that sum to 8 are:

1.  5 -> 3
2.  5 -> 2 -> 1
3. -3 -> 11
```

**Difficult:** `Medium`

**Tags:** `Tree`

### Solution One

```javascript
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */

class IdxToLastSum {
  /**
   * @param {number} target
   */
  constructor(target) {
    this.count = 0;
    this.target = target;
    /** @type {number[]} */
    this.idxToLastSum = [];
  }

  push(val) {
    this.count = this.target === val ? 1 : 0;
    this.idxToLastSum = this.idxToLastSum.map((sum) => {
      sum += val;
      if (sum === this.target) {
        this.count++;
      }

      return sum;
    });

    this.idxToLastSum.push(val);
  }

  pop() {
    const val = this.idxToLastSum.pop();
    this.idxToLastSum = this.idxToLastSum.map((sum) => (sum -= val));
  }
}

const helper = (root, _) => {
  if (root === null) {
    return;
  }

  _.valList.push(root.val);
  _.res += _.valList.count;
  helper(root.left, _);
  helper(root.right, _);
  _.valList.pop();
};

/**
 * @param {TreeNode} root
 * @param {number} sum
 * @return {number}
 */
const pathSum = function (root, sum) {
  const _ = {
    valList: new IdxToLastSum(sum),
    res: 0,
  };

  helper(root, _);

  return _.res;
};
```

### Solutions

[437. Path Sum III - Solution](https://leetcode.com/problems/path-sum-iii/solution/)
