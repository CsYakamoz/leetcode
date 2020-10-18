## [994. Rotting Oranges](https://leetcode.com/problems/rotting-oranges/)

### Description

In a given grid, each cell can have one of three values:

- the value `0` representing an empty cell;
- the value `1` representing a fresh orange;
- the value `2` representing a rotten orange.

Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return `-1` instead.

**Example 1:**

**![img](https://assets.leetcode.com/uploads/2019/02/16/oranges.png)**

```
Input: [[2,1,1],[1,1,0],[0,1,1]]
Output: 4
```

**Example 2:**

```
Input: [[2,1,1],[0,1,1],[1,0,1]]
Output: -1
Explanation:  The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.
```

**Example 3:**

```
Input: [[0,2]]
Output: 0
Explanation:  Since there are already no fresh oranges at minute 0, the answer is just 0.
```

**Note:**

1. `1 <= grid.length <= 10`
2. `1 <= grid[0].length <= 10`
3. `grid[i][j]` is only `0`, `1`, or `2`.

**Difficulty:** `Medium`

**Tags:** `Breadth-first Search`

### Solution One

```javascript
/**
 * @param {number[][]} grid
 * @return {number}
 */
const orangesRotting = function (grid) {
  /** @type {Set<number>} */
  const rottingIdx = new Set();
  /** @type {Set<number>} */
  const freshIdx = new Set();

  const rowLength = grid.length;
  const colLength = grid[0].length;
  for (let i = 0; i < rowLength; i++) {
    for (let j = 0; j < colLength; j++) {
      const val = grid[i][j];
      if (val === 1) {
        freshIdx.add(i * colLength + j);
      } else if (val === 2) {
        rottingIdx.add(i * colLength + j);
      }
    }
  }

  // there are already no fresh oranges at minute 0
  if (freshIdx.size === 0) {
    return 0;
  }

  let minutes = 0;
  while (freshIdx.size !== 0) {
    let canRotten = false;

    const idxList = Array.from(rottingIdx.values());
    for (const idx of idxList) {
      const [row, col] = [Math.floor(idx / colLength), idx % colLength];

      // up
      const upIdx = idx - colLength;
      if (row > 0 && freshIdx.has(upIdx)) {
        canRotten = true;
        rottingIdx.add(upIdx);
        freshIdx.delete(upIdx);
      }

      // down
      const downIdx = idx + colLength;
      if (row < rowLength - 1 && freshIdx.has(downIdx)) {
        canRotten = true;
        rottingIdx.add(downIdx);
        freshIdx.delete(downIdx);
      }

      // left
      const leftIdx = idx - 1;
      if (col > 0 && freshIdx.has(leftIdx)) {
        canRotten = true;
        rottingIdx.add(leftIdx);
        freshIdx.delete(leftIdx);
      }

      // right
      const rightIdx = idx + 1;
      if (col < grid[0].length - 1 && freshIdx.has(rightIdx)) {
        canRotten = true;
        rottingIdx.add(rightIdx);
        freshIdx.delete(rightIdx);
      }

      rottingIdx.delete(idx);
    }

    if (canRotten) {
      minutes++;
    } else {
      break;
    }
  }

  // Return -1 if there are some fresh oranges that never be rotten
  return freshIdx.size === 0 ? minutes : -1;
};
```

### Solutions

[994. Rotting Oranges - Solution](https://leetcode.com/problems/rotting-oranges/solution/)
