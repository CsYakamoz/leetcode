## [647. Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/description/)

### Description

Given a string, your task is to count how many palindromic substrings in this string.

The substrings with different start indexes or end indexes are counted as different substrings even they consist of same characters.

**Example 1:**

```
Input: "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".

```

**Example 2:**

```
Input: "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".

```

**Note:**

1. The input string length won't exceed 1000.

**Difficulty:** `Medium`

**Tags:** `Dynamic Programming` `String`

### Solution One

思路来源于题目 [5. Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/#/description) 中的 Top Solutions

```c++
class Solution {
public:
    int countSubstrings(string s) {
        int count = 0;
        for (size_t i = 0; i < s.size();)
        {
            size_t j = i;
            size_t k = i;

            while (k < s.size() - 1 && s[k] == s[k + 1])
                k++;

            i = k + 1;
            int length = i - j;
            // For a substring that constructs by a single character
            count += (1 + length) * length / 2;

            // Extended to both sides
            while (j > 0 && k < s.size() - 1 && s[j - 1] == s[k + 1])
            {
                --j;
                k++;
                count++;
            }
        }
        return count;
    }
};
```

### Solution Two - In Top Solutions

[Java solution, 8 lines, extendPalindrome](https://discuss.leetcode.com/topic/96819/java-solution-8-lines-extendpalindrome)

> Idea is start from each index and try to extend palindrome for both odd and even length.

```java
public class Solution {
    int count = 0;

    public int countSubstrings(String s) {
        if (s == null || s.length() == 0) return 0;

        for (int i = 0; i < s.length(); i++) { // i is the mid point
            extendPalindrome(s, i, i); // odd length;
            extendPalindrome(s, i, i + 1); // even length
        }

        return count;
    }

    private void extendPalindrome(String s, int left, int right) {
        while (left >=0 && right < s.length() && s.charAt(left) == s.charAt(right)) {
            count++; left--; right++;
        }
    }
}
```
