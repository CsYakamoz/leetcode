## [434. Number of Segments in a String](https://leetcode.com/problems/number-of-segments-in-a-string/#/description)

### Description

Count the number of segments in a string, where a segment is defined to be a contiguous sequence of non-space characters.

Please note that the string does not contain any **non-printable** characters.

**Example:**

```
Input: "Hello, my name is John"
Output: 5
```



**Difficult:** `Easy`

**Tags:** `String`



### Solution One

```c++
class Solution {
public:
    int countSegments(string s) {
        size_t i = 0;
        int num = 0;
        while (i < s.size())
        {
            if (isspace(s[i]))
            {
                i++;
            }
            else
            {
                while (i < s.size() && !isspace(s[i]))
                {
                    i++;
                }
                num++;
                i++;
            }
        }
        return num;
    }
};
```



### Solution Two

```c++
class Solution {
public:
    int countSegments(string s) {
        int res = 0;
        s.push_back(' ');
        for(int i = 1; i < s.size(); ++i)
          if(s[i] == ' ' && s[i-1] != ' ') ++res;
        return res;
    }
};
```



