## [73. Set Matrix Zeroes](https://leetcode.com/problems/set-matrix-zeroes/)

### Description

<p>Given an&nbsp;<code><em>m</em> x <em>n</em></code> matrix. If an element is <strong>0</strong>, set its entire row and column to <strong>0</strong>. Do it <a href="https://en.wikipedia.org/wiki/In-place_algorithm" target="_blank"><strong>in-place</strong></a>.</p>

<p><strong>Follow up:</strong></p>

<ul>
	<li>A straight forward solution using O(<em>m</em><em>n</em>) space is probably a bad idea.</li>
	<li>A simple improvement uses O(<em>m</em> + <em>n</em>) space, but still not the best solution.</li>
	<li>Could you devise a constant space solution?</li>
</ul>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/08/17/mat1.jpg" style="width: 450px; height: 169px;" />
<pre>
<strong>Input:</strong> matrix = [[1,1,1],[1,0,1],[1,1,1]]
<strong>Output:</strong> [[1,0,1],[0,0,0],[1,0,1]]
</pre>

<p><strong>Example 2:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/08/17/mat2.jpg" style="width: 450px; height: 137px;" />
<pre>
<strong>Input:</strong> matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
<strong>Output:</strong> [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>m == matrix.length</code></li>
	<li><code>n == matrix[0].length</code></li>
	<li><code>1 &lt;= m, n &lt;= 200</code></li>
	<li><code>-2<sup>31</sup> &lt;= matrix[i][j] &lt;= 2<sup>31</sup> - 1</code></li>
</ul>

**Difficulty:** `Medium`

**Tags:** `Array`

### Solution One

```javascript
/**
 * @param {number[][]} matrix
 * @return {void} Do not return anything, modify matrix in-place instead.
 */
const setZeroes = function (matrix) {
  const m = matrix.length;
  const n = matrix[0].length;

  let setRow = false;
  let setCol = false;
  for (let i = 0; i < m; i++) {
    if (matrix[i][0] === 0) {
      setCol = true;
      break;
    }
  }
  for (let j = 0; j < n; j++) {
    if (matrix[0][j] === 0) {
      setRow = true;
      break;
    }
  }

  for (let i = 1; i < m; i++) {
    for (let j = 1; j < n; j++) {
      if (matrix[i][j] === 0) {
        matrix[i][0] = 0;
        matrix[0][j] = 0;
      }
    }
  }

  for (let i = 1; i < m; i++) {
    if (matrix[i][0] === 0) {
      for (let j = 1; j < n; j++) {
        matrix[i][j] = 0;
      }
    }
  }
  for (let j = 1; j < n; j++) {
    if (matrix[0][j] === 0) {
      for (let i = 1; i < m; i++) {
        matrix[i][j] = 0;
      }
    }
  }

  if (setRow) {
    for (let j = 0; j < n; j++) {
      matrix[0][j] = 0;
    }
  }

  if (setCol) {
    for (let i = 0; i < m; i++) {
      matrix[i][0] = 0;
    }
  }
};
```

### Solution Two - In Top Solutions

[Any shorter O(1) space solution?](<https://leetcode.com/problems/set-matrix-zeroes/discuss/26014/Any-shorter-O(1)-space-solution>)

```cpp
void setZeroes(vector<vector<int> > &matrix) {
    int col0 = 1, rows = matrix.size(), cols = matrix[0].size();

    for (int i = 0; i < rows; i++) {
        if (matrix[i][0] == 0) col0 = 0;
        for (int j = 1; j < cols; j++)
            if (matrix[i][j] == 0)
                matrix[i][0] = matrix[0][j] = 0;
    }

    for (int i = rows - 1; i >= 0; i--) {
        for (int j = cols - 1; j >= 1; j--)
            if (matrix[i][0] == 0 || matrix[0][j] == 0)
                matrix[i][j] = 0;
        if (col0 == 0) matrix[i][0] = 0;
    }
}
```
