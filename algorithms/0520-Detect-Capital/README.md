## [520. Detect Capital](https://leetcode.com/problems/detect-capital/#/description)

### Description

Given a word, you need to judge whether the usage of capitals in it is right or not.

We define the usage of capitals in a word to be right when one of the following cases holds:

1. All letters in this word are capitals, like "USA".
2. All letters in this word are not capitals, like "leetcode".
3. Only the first letter in this word is capital if it has more than one letter, like "Google".

Otherwise, we define that this word doesn't use capitals in a right way.

**Example 1:**

```
Input: "USA"
Output: True
```

**Example 2:**

```
Input: "FlaG"
Output: False
```

**Note:** The input will be a non-empty word consisting of uppercase and lowercase latin letters.

**Difficulty:** `Easy`

**Tags:** `String`

### Solution One

要么全大写，要么全小写，要么只有第一个大写

`count == 0`对应全小写

`count == word.size()`对应全大写

`count == 1 && isupper(word[0])`对应只有第一个大写

```c++
class Solution {
public:
    bool detectCapitalUse(string word) {
        int count = 0;
        for (auto &c : word)
            if (isupper(c))
            {
                ++count;
            }
        if (count == 0 || count == word.size() || (count == 1 && isupper(word[0])))
        {
            return true;
        }
        else
        {
            return false;
        }
    }
};
```
