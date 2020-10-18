## [58. Length of Last Word](https://leetcode.com/problems/length-of-last-word/#/description)

### Description

Given a string _s_ consists of upper/lower-case alphabets and empty space characters `' '`, return the length of last word in the string.

If the last word does not exist, return 0.

**Note:** A word is defined as a character sequence consists of non-space characters only.

For example,
Given _s_ = `"Hello World"`,
return `5`.

**Difficulty:** `Easy`

**Tags:** `String`

### Solution One

```c++
class Solution {
public:
    int lengthOfLastWord(string s) {
        istringstream iss(s);
        string res;
        while (iss >> res)
        {

        }
        return res.size();
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    int lengthOfLastWord(string s) {
        int len = 0, tail = s.size() - 1;
        // skip traling white space
        while( tail >= 0 && s[tail] == ' ') tail--;
        while( tail >= 0 && s[tail] != ' '){
            len++; tail--;
        }
        return len;
    }
};
```
