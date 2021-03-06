## [151. Reverse Words In a String](https://leetcode.com/problems/reverse-words-in-a-string/#/description)

### Description

Given an input string, reverse the string word by word.

For example,

Given s = "`the sky is blue`",

return "`blue is sky the`".

**Update (2015-02-12):**
For C programmers: Try to solve it _in-place_ in _O_(1) space.

Clarification:

- What constitutes a word?
  A sequence of non-space characters constitutes a word.
- Could the input string contain leading or trailing spaces?
  Yes. However, your reversed string should not contain leading or trailing spaces.
- How about multiple spaces between two words?
  Reduce them to a single space in the reversed string.

**Difficulty:** `Medium`

**Tags:** `String`

### Solution One

```c++
class Solution {
public:
    void reverseWords(string &s) {
        istringstream iss(s);
        string str;
        vector<string> v;
        while (iss >> str)
        {
            v.push_back(str);
        }
        s = "";
        for (int i = v.size() - 1; i >= 0; i--)
        {
            s += v[i] + ' ';
        }
        s.pop_back();
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    void reverseWords(string &s) {
        reverse(s.begin(), s.end());	// reverse whole string

        size_t i = s.find_first_not_of(' ');
        if (i == string::npos)
        {	// there are no word in string
            s = "";
            return;
        }
        size_t storeIndex = 0;
        size_t wordStart;		// the first character's index of every words
        size_t wordLength;

        while (i < s.size())
        {
            // skips spaces in front of the word
            while (i < s.size() && s[i] == ' ')
            {
                i++;
            }
            if (i == s.size())
            {	// avoid adding extra space
                break;
            }
            // copy word
            wordStart = storeIndex;
            wordLength = 0;
            while (i < s.size() && s[i] != ' ')
            {
                wordLength++;
                s[storeIndex++] = s[i++];
            }
            reverse(&s[wordStart], &s[wordStart + wordLength]);	// reverse word
            s[storeIndex++] = ' ';	// single space between two words
        }
        s = s.substr(0, storeIndex - 1);	// delete the trailing space
    }
};
```
