## [500. Keyboard Row](https://leetcode.com/problems/keyboard-row/description/)

### Description

Given a List of words, return the words that can be typed using letters of **alphabet** on only one row's of American keyboard like the image below.

![American keyboard](https://leetcode.com/static/images/problemset/keyboard.png)

**Example 1:**

```
Input: ["Hello", "Alaska", "Dad", "Peace"]
Output: ["Alaska", "Dad"]

```

**Note:**

1. You may use one character in the keyboard more than once.
2. You may assume the input string will only contain letters of alphabet.

**Difficulty:** `Easy`

**Tags:** `Hash Table`

### Solution One

```c++
class Solution {
public:
    vector<string> findWords(vector<string>& words) {
        vector<string> res;
        vector<string> row{ "qwertyuiop", "asdfghjkl", "zxcvbnm" };
        for (auto &s : words)
        {
            size_t r = 0;
            while (row[r].find(tolower(s[0])) == string::npos)
            {
                r++;
            }
            size_t i = 1;
            for (; i < s.size(); i++)
            {
                if (row[r].find(tolower(s[i])) == string::npos)
                {
                    break;
                }
            }
            if (i == s.size())
            {
                res.push_back(s);
            }
        }
        return res;
    }
};
```

### Solution Two

```c++
class Solution {
public:
    vector<string> findWords(vector<string>& words) {
        vector<string> res;
        vector<unordered_set<char>> row{
            {'q','w','e','r','t','y','u','i','o','p'} ,
            {'a','s','d','f','g','h','j','k','l'} ,
            {'z','x','c','v','b','n','m'}
        };
        for (auto &s : words)
        {
            size_t r = 0;
            while (row[r].find(tolower(s[0])) == row[r].end())
            {
                r++;
            }
            size_t i = 1;
            for (; i < s.size(); i++)
            {
                if (row[r].find(tolower(s[i])) == row[r].end())
                {
                    break;
                }
            }
            if (i == s.size())
            {
                res.push_back(s);
            }
        }
        return res;
    }
};
```

### Solution Three - In Top Solutions

```c++
class Solution {
public:
    vector<string> findWords(vector<string>& words) {
        vector<string> res;
        unordered_set<char> row1{'q','w','e','r','t','y','u','i','o','p'};
        unordered_set<char> row2{'a','s','d','f','g','h','j','k','l'};
        unordered_set<char> row3{'z','x','c','v','b','n','m'};
        for (string word : words) {
            int one = 0, two = 0, three = 0;
            for (char c : word) {
                if (c < 'a') c += 32;
                if (row1.count(c)) one = 1;
                if (row2.count(c)) two = 1;
                if (row3.count(c)) three = 1;
                if (one + two + three > 1) break;
            }
            if (one + two + three == 1) res.push_back(word);
        }
        return res;

    }
};
```
