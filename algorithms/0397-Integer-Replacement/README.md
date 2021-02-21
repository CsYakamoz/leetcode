## [397. Integer Replacement](https://leetcode.com/problems/integer-replacement/)

### Description

<p>Given a positive integer <code>n</code>,&nbsp;you can apply one of the following&nbsp;operations:</p>

<ol>
	<li>If <code>n</code> is even, replace <code>n</code> with <code>n / 2</code>.</li>
	<li>If <code>n</code> is odd, replace <code>n</code> with either <code>n + 1</code> or <code>n - 1</code>.</li>
</ol>

<p>Return <em>the minimum number of operations needed for <code>n</code> to become <code>1</code></em>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> n = 8
<strong>Output:</strong> 3
<strong>Explanation:</strong> 8 -&gt; 4 -&gt; 2 -&gt; 1
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> n = 7
<strong>Output:</strong> 4
<strong>Explanation: </strong>7 -&gt; 8 -&gt; 4 -&gt; 2 -&gt; 1
or 7 -&gt; 6 -&gt; 3 -&gt; 2 -&gt; 1
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> n = 4
<strong>Output:</strong> 2
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 2<sup>31</sup> - 1</code></li>
</ul>

**Difficulty:** `Medium`

**Tags:** `Math` `Bit Manipulation`

### Solution One

> 如果不理解的话, 建议看 Solution Two

首先通过 `helper` 函数找到连续的 `1` 以及总共有多少个 `0`

> 例如:
> 对于 n = 100000000, 其二进制为 101111101011110000100000000, 则 helper 函数返回
> operation: 15
> intervalList: [[8, 9], [13, 17], [18, 19], [20, 25], [26, 27]]

接着进行如下逻辑: 当前 `interval` 是否为最后一个?

> 下文中的 `count` 指当前 `interval` 中连续的 `1` 的数量

- 是:

  - count == 1: 什么也不做;
  - count == 2: 两步操作, `n-1` 和 `n/2`;
  - count > 2: `count+1` 步操作, `n+1` 和 `count` 次 `n/2`;

- 否:

  - 当前 `interval` 中结尾是否与下一个的开头差 **1**?
    是:

    - count == 1: 两步操作, `n-1` 和 `n/2`;
    - count > 1: `count+1` 步操作, `n+1` 和 `count` 次 `n/2`, 主要目的是将多个 1 变成单个 1, 且该 1 与下一个 interval 连续;

    否:

    - count <= 2: `2*count` 步操作, (`n-1` 和 `n/2`) 重复 `count` 次;
    - count > 2: `count+2` 步操作, `n+1` 和 `count` 次 `n/2` 和 `n-1`;

```go
func integerReplacement(n int) int {
	operation, intervalList := helper(n)

	length := len(intervalList)
	for i := 0; i < length; i++ {
		if i == length-1 {
			count := intervalList[i][1] - intervalList[i][0]

			if count == 1 {
				break
			} else if count == 2 {
				operation += 2
			} else {
				operation += count + 1
			}
		} else {
			if intervalList[i+1][0]-intervalList[i][1] == 1 {
				count := intervalList[i][1] - intervalList[i][0]
				if count == 1 {
					operation += 2
				} else {
					operation += intervalList[i][1] - intervalList[i][0] + 1
					intervalList[i+1][0]--
					operation--
				}
			} else {
				count := intervalList[i][1] - intervalList[i][0]
				if count > 2 {
					operation += count + 2
				} else {
					operation += count * 2
				}
			}
		}
	}

	return operation
}

func helper(n int) (int, [][2]int) {
	operation, intervalList := 0, make([][2]int, 0)

	i, start := 0, -1
	for curr := n; curr != 0 && i < 31; i++ {
		if n&(1<<i) != 0 {
			if start == -1 {
				start = i
			}
		} else {
			if start != -1 {
				intervalList = append(intervalList, [2]int{start, i})
				start = -1
			}
			operation++
		}

		curr >>= 1
	}

	if start != -1 {
		intervalList = append(intervalList, [2]int{start, i})
	}

	return operation, intervalList
}
```

### Solution Two - In Top Solutions

[A couple of Java solutions with explanations](https://leetcode.com/problems/integer-replacement/discuss/87920/A-couple-of-Java-solutions-with-explanations)

```java
public int integerReplacement(int n) {
    int c = 0;
    while (n != 1) {
        if ((n & 1) == 0) {
            n >>>= 1;
        } else if (n == 3 || Integer.bitCount(n + 1) > Integer.bitCount(n - 1)) {
            --n;
        } else {
            ++n;
        }
        ++c;
    }
    return c;
}
```
