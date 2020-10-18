## [678. Valid Parenthesis String](https://leetcode.com/problems/valid-parenthesis-string/)

### Description

Given a string containing only three types of characters: '(', ')' and '\*', write a function to check whether this string is valid. We define the validity of a string by these rules:

1. Any left parenthesis `'('` must have a corresponding right parenthesis `')'`.
2. Any right parenthesis `')'` must have a corresponding left parenthesis `'('`.
3. Left parenthesis `'('` must go before the corresponding right parenthesis `')'`.
4. `'*'` could be treated as a single right parenthesis `')'` or a single left parenthesis `'('` or an empty string.
5. An empty string is also valid.

**Example 1:**

```
Input: "()"
Output: True
```

**Example 2:**

```
Input: "(*)"
Output: True
```

**Example 3:**

```
Input: "(*))"
Output: True
```

**Note:**

1. The string size will be in the range [1, 100].

**Difficulty:** `Medium`

**Tags:** `String`

### Solution One

1. 遍历数组，做如下事
   - 遇到 `)`, 则判断是否有 `)` 或 `*` 可用，若无，则直接返回 `false`, 否则对应的栈进行 `pop` 操作
   - 遇到 `(` 或 `*`, 则将下标 `push` 到对应的栈
2. 走到此步时，字符串中只剩下 `(` 和 `*`
   - 对于每一个 `(` 而言，在其后面必须存在 `*`, 才能消掉，故从数组后面开始循环
3. 判断 `leftBrackets` 长度是否为 0

或者看该解答：[Java using 2 stacks. O(n) space and time complexity.](<https://leetcode.com/problems/valid-parenthesis-string/discuss/107572/Java-using-2-stacks.-O(n)-space-and-time-complexity.>)

```javascript
/**
 * @param {any[]} arr
 * @return {any}
 */
const lastElemOrNeg1 = (arr) => (arr.length !== 0 ? arr[arr.length - 1] : -1);

/**
 * @param {string} s
 * @return {boolean}
 */
const checkValidString = function (s) {
  const starStack = [];
  const leftBrackets = [];

  for (let i = 0, len = s.length; i < len; i++) {
    const char = s[i];

    switch (char) {
      case ')':
        {
          if (leftBrackets.length !== 0) {
            leftBrackets.pop();
          } else if (starStack.length !== 0) {
            starStack.pop();
          } else {
            return false;
          }
        }
        break;
      case '*':
        starStack.push(i);
        break;
      case '(':
        leftBrackets.push(i);
        break;
      default:
        throw new Error(`unknown char: ${char}`);
    }
  }

  while (
    leftBrackets.length !== 0 &&
    lastElemOrNeg1(starStack) > lastElemOrNeg1(leftBrackets)
  ) {
    starStack.pop();
    leftBrackets.pop();
  }

  return leftBrackets.length === 0;
};
```

### Solutions

[678. Valid Parenthesis String - Solution](https://leetcode.com/problems/valid-parenthesis-string/solution/)
