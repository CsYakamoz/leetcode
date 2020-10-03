## [1544. Make The String Great](https://leetcode.com/problems/make-the-string-great/)

### Description

Given a string `s` of lower and upper case English letters.

A good string is a string which doesn't have **two adjacent characters** `s[i]` and `s[i + 1]` where:

- `0 <= i <= s.length - 2`
- `s[i]` is a lower-case letter and `s[i + 1]` is the same letter but in upper-case or **vice-versa**.

To make the string good, you can choose **two adjacent** characters that make the string bad and remove them. You can keep doing this until the string becomes good.

Return _the string_ after making it good. The answer is guaranteed to be unique under the given constraints.

**Notice** that an empty string is also good.

**Example 1:**

```
Input: s = "leEeetcode"
Output: "leetcode"
Explanation: In the first step, either you choose i = 1 or i = 2, both will result "leEeetcode" to be reduced to "leetcode".
```

**Example 2:**

```
Input: s = "abBAcC"
Output: ""
Explanation: We have many possible scenarios, and all lead to the same answer. For example:
"abBAcC" --> "aAcC" --> "cC" --> ""
"abBAcC" --> "abBA" --> "aA" --> ""
```

**Example 3:**

```
Input: s = "s"
Output: "s"
```

**Constraints:**

- `1 <= s.length <= 100`
- `s` contains only lower and upper case English letters.

**Difficult:** `Easy`

**Tags:** `String` `Stack`

### Solution One

```javascript
/**
 * @param {string[]} arr
 * @return {string}
 */
const lastElem = (arr) => arr[arr.length - 1];

/**
 * @param {string} c1
 * @param {string} c2
 * @return {boolean}
 */
const needToRemove = (c1, c2) =>
  Math.abs(c1.charCodeAt(0) - c2.charCodeAt(0)) === 32;

/**
 * @param {string} s
 * @return {string}
 */
const makeGood = function (s) {
  if (s.length < 2) {
    return s;
  }

  const stack = [s[0]];

  for (let i = 1, len = s.length; i < len; i++) {
    const c = s[i];

    if (stack.length === 0) {
      stack.push(c);
      continue;
    }

    if (needToRemove(c, lastElem(stack))) {
      stack.pop();
      continue;
    }

    stack.push(c);
  }

  return stack.join('');
};
```
