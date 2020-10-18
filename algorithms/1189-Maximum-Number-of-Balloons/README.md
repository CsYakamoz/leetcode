## [1189. Maximum Number of Balloons](https://leetcode.com/problems/maximum-number-of-balloons/)

### Description

Given a string `text`, you want to use the characters of `text` to form as many instances of the word **"balloon"** as possible.

You can use each character in `text` **at most once**. Return the maximum number of instances that can be formed.

**Example 1:**

**![img](https://assets.leetcode.com/uploads/2019/09/05/1536_ex1_upd.JPG)**

```
Input: text = "nlaebolko"
Output: 1
```

**Example 2:**

**![img](https://assets.leetcode.com/uploads/2019/09/05/1536_ex2_upd.JPG)**

```
Input: text = "loonbalxballpoon"
Output: 2
```

**Example 3:**

```
Input: text = "leetcode"
Output: 0
```

**Constraints:**

- `1 <= text.length <= 10^4`
- `text` consists of lower case English letters only.

**Difficulty:** `Easy`

**Tags:** `Hash Table` `String`

### Solution One

```javascript
/**
 * @param {string} text
 * @returns {object}
 */
const createCountDict = function (text) {
  return text.split('').reduce((prev, curr) => {
    if (prev[curr] === undefined) {
      prev[curr] = 1;
    } else {
      prev[curr] += 1;
    }

    return prev;
  }, {});
};

/**
 * @param {string} text
 * @return {number}
 */
const maxNumberOfBalloons = function (text) {
  const keyword = 'balloon';
  const keywordDict = createCountDict(keyword);
  const textDict = createCountDict(text);

  let maxNumber = Number.MAX_SAFE_INTEGER;
  for (const c in keywordDict) {
    const division = keywordDict[c];
    const count = Math.floor((textDict[c] || 0) / division);
    maxNumber = Math.min(maxNumber, count);
  }

  return maxNumber;
};
```
