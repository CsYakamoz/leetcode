## [76. Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/description/)

### Description

Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

For example,
**S** = `"ADOBECODEBANC"`
**T** = `"ABC"`

Minimum window is `"BANC"`.

**Note:**
If there is no such window in S that covers all characters in T, return the empty string `""`.

If there are multiple such windows, you are guaranteed that there will always be only one unique minimum window in S.

**Difficult:** `Hard`

**Tags:** `Hash Table` `Two Pointers` `String`

### Solution One - In Top Solutions

[Here is a 10-line template that can solve most 'substring' problems](https://discuss.leetcode.com/topic/30941/here-is-a-10-line-template-that-can-solve-most-substring-problems)

```c++
class Solution {
public:
    string minWindow(string s, string t) {

        vector<int> hash(128, 0);
        int start = 0;
        int begin = 0;
        int end = 0;
        int counter = t.size();
        int length = INT_MAX;

        for (auto c : t)
            ++hash[c];

        while (end < s.size()) {

            if (hash[s[end++]]-- > 0)
                --counter;

            while (counter == 0) {

                if (length > end - begin) {
                    start = begin;
                    length = end - begin;
                }

                if (hash[s[begin++]]++ == 0)
                    ++counter;
            }
        }

        return length == INT_MAX ? "" : s.substr(start, length);
    }
};
```
