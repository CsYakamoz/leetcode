## [459. Repeated Substring Pattern](https://leetcode.com/problems/repeated-substring-pattern/#/description)

### Description

Given a non-empty string check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.

You may assume the given string consists of lowercase English letters only and its length will not exceed 10000.

**Example 1:**

```
Input: "abab"
Output: True
Explanation: It's the substring "ab" twice.
```

**Example 2:**

```
Input: "aba"

Output: False

```

**Example 3:**

```
Input: "abcabcabcabc"

Output: True

Explanation: It's the substring "abc" four times. (And the substring "abcabc" twice.)
```



**Difficult:** `Easy`

**Tags:** `String`



### Solution One

如果`s`是重复的，那么重复次数至少为 2，故从 2 开始

我们首先检查前 i 个字符 和 后 i 个字符是否相同

若相同，则检查中间的每 i 个字符是否和前 i 个字符相同

若不相同，则`i = s.size() / ++div`，继续检查

直到 i 为 0

```c++
class Solution {
public:
    bool repeatedSubstringPattern(string s) {
        int div = 2;	// At least repeats 2 times
        int i = s.size() / div;
        while (i != 0)
        {
            if (s.substr(0, i) == s.substr(s.size() - i))
            {
                string sub = s.substr(0, i);
                size_t j = i;
                bool isPattern = true;
                while (j < s.size())
                {
                    if (sub != s.substr(j, i))
                    {
                        isPattern = false;
                        break;
                    }
                    j += i;
                }
                if (isPattern)
                {
                    return true;
                }
            }
            div++;
            i = s.size() / div;
        }
        return false;
    }
};
```



### Solution Two - In Top Solutions

[Java Simple Solution with Explanation](https://discuss.leetcode.com/topic/67992/java-simple-solution-with-explanation)

1. The length of the repeating substring must be a divisor of the length of the input string
2. Search for all possible divisor of `str.length`, starting for `length/2`
3. If `i` is a divisor of `length`, repeat the substring from `0` to `i` the number of times `i` is contained in `s.length`
4. If the repeated substring is equals to the input `str` return `true`

```java
public boolean repeatedSubstringPattern(String str) {
    int l = str.length();
    for(int i=l/2;i>=1;i--) {
        if(l%i==0) {
            int m = l/i;
            String subS = str.substring(0,i);
            StringBuilder sb = new StringBuilder();
            for(int j=0;j<m;j++) {
                sb.append(subS);
            }
            if(sb.toString().equals(str)) return true;
        }
    }
    return false;
}
```



