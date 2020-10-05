## [1047. Remove All Adjacent Duplicates In String](https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/)

### Description

Given a string `S` of lowercase letters, a _duplicate removal_ consists of choosing two adjacent and equal letters, and removing them.

We repeatedly make duplicate removals on S until we no longer can.

Return the final string after all such duplicate removals have been made. It is guaranteed the answer is unique.

**Example 1:**

```
Input: "abbaca"
Output: "ca"
Explanation:
For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal, and this is the only possible move.  The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca".
```

**Note:**

1. `1 <= S.length <= 20000`
2. `S` consists only of English lowercase letters.

**Difficult:** `Easy`

**Tags:** `Stack`

### Solution One

```javascript
/**
 * @param {string} S
 * @return {string}
 */
var removeDuplicates = function (S) {
  const stack = [];

  for (const c of S) {
    stack.push(c);

    while (stack.length > 1) {
      const last1 = stack.pop();
      const last2 = stack.pop();
      if (last1 !== last2) {
        stack.push(last2);
        stack.push(last1);
        break;
      }
    }
  }

  return stack.join('');
};
```