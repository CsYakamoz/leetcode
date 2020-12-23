## [1329. Sort the Matrix Diagonally](https://leetcode.com/problems/sort-the-matrix-diagonally/)

### Description

<p>Given a <code>m * n</code> matrix <code>mat</code>&nbsp;of integers, sort it diagonally in ascending order from the top-left to the bottom-right then return the sorted array.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/01/21/1482_example_1_2.png" style="width: 500px; height: 198px;" />
<pre>
<strong>Input:</strong> mat = [[3,3,1,1],[2,2,1,2],[1,1,1,2]]
<strong>Output:</strong> [[1,1,1,1],[1,2,2,2],[1,2,3,3]]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>m ==&nbsp;mat.length</code></li>
	<li><code>n ==&nbsp;mat[i].length</code></li>
	<li><code>1 &lt;= m, n&nbsp;&lt;= 100</code></li>
	<li><code>1 &lt;= mat[i][j] &lt;= 100</code></li>
</ul>

**Difficulty:** `Medium`

**Tags:** `Array` `Sort`

### Solution One

每个对角线进行插入排序

```go
func diagonalSort(mat [][]int) [][]int {
	rowLen, colLen := len(mat), len(mat[0])

	for i := 1; i < rowLen; i++ {
		for j := 1; j < colLen; j++ {
			row, col := i-1, j-1
			for row >= 0 && col >= 0 && mat[row][col] > mat[row+1][col+1] {
				mat[row][col], mat[row+1][col+1] =
					mat[row+1][col+1], mat[row][col]

				row--
				col--
			}
		}
	}

	return mat
}
```
