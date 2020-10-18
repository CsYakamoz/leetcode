## [383. Ransom Note](https://leetcode.com/problems/ransom-note/tabs/description/)

### Description

Given an arbitrary ransom note string and another string containing letters from all the magazines, write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.

Each letter in the magazine string can only be used once in your ransom note.

**Note:**
You may assume that both strings contain only lowercase letters.

```
canConstruct("a", "b") -> false
canConstruct("aa", "ab") -> false
canConstruct("aa", "aab") -> true
```

**Difficulty:** `Easy`

**Tags:** `String`

### Solution One

```c++
class Solution {
public:
    bool canConstruct(string ransomNote, string magazine) {
        vector<int> letters(26, 0);
        for (auto i : magazine)
            letters[i - 'a']++;
        for (auto i : ransomNote)
        {
            letters[i - 'a']--;
            if (letters[i - 'a'] < 0)
            {
                return false;
            }
        }
        return true;
    }
};
```

### Solution Two - In Top Solutions

**One**的是 29ms，这个是 19ms，目前不知道差别是什么

```c++
class Solution {
public:
    bool canConstruct(string ransomNote, string magazine) {
        vector<int> letters(26, 0);
        for (char ch : magazine) {
            letters[ch-'a']++;
        }

        for (char ch : ransomNote) {
            letters[ch-'a']--;
            if (letters[ch-'a']<0) return false;
        }

        return true;
    }
};
```
