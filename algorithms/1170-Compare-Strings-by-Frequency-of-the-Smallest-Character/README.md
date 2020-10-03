## [1170. Compare Strings by Frequency of the Smallest Character](https://leetcode.com/problems/compare-strings-by-frequency-of-the-smallest-character/)

### Description

Let's define a function `f(s)` over a non-empty string `s`, which calculates the frequency of the smallest character in `s`. For example, if `s = "dcce"` then `f(s) = 2` because the smallest character is `"c"` and its frequency is 2.

Now, given string arrays `queries` and `words`, return an integer array `answer`, where each `answer[i]` is the number of words such that `f(queries[i])` < `f(W)`, where `W` is a word in `words`.

**Example 1:**

```
Input: queries = ["cbd"], words = ["zaaaz"]
Output: [1]
Explanation: On the first query we have f("cbd") = 1, f("zaaaz") = 3 so f("cbd") < f("zaaaz").
```

**Example 2:**

```
Input: queries = ["bbb","cc"], words = ["a","aa","aaa","aaaa"]
Output: [1,2]
Explanation: On the first query only f("bbb") < f("aaaa"). On the second query both f("aaa") and f("aaaa") are both > f("cc").
```

**Constraints:**

- `1 <= queries.length <= 2000`
- `1 <= words.length <= 2000`
- `1 <= queries[i].length, words[i].length <= 10`
- `queries[i][j]`, `words[i][j]` are English lowercase letters.

**Difficult:** `Easy`

**Tags:** `Array` `String`

### Solution One

```javascript
/**
 * @param {string} text
 * @returns {number}
 */
const calc = (text) => {
  let freq = 1;
  let smallestChar = text.charCodeAt(0);

  for (let i = 1, length = text.length; i < length; i++) {
    const char = text.charCodeAt(i);
    if (char < smallestChar) {
      smallestChar = char;
      freq = 1;
    } else if (char === smallestChar) {
      freq += 1;
    }
  }

  return freq;
};

/**
 * @param {string} query
 * @param {number[]} sortedWords
 */
const findSuitableElem = (query, sortedWords) => {
  const f = calc(query);

  let [left, right] = [0, sortedWords.length];
  while (left < right) {
    const mid = Math.floor((right - left) / 2) + left;
    if (sortedWords[mid] > f) {
      right = mid;
    } else {
      left = mid + 1;
    }
  }

  return sortedWords.length - right;
};

/**
 * @param {string[]} queries
 * @param {string[]} words
 * @return {number[]}
 */
const numSmallerByFrequency = function (queries, words) {
  const sortedWords = words.map(calc).sort((a, b) => a - b);

  return queries.map((query) => findSuitableElem(query, sortedWords));
};
```
