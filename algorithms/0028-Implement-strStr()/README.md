## [28. Implement strStr()](https://leetcode.com/problems/implement-strstr/#/description)

### Description

Implement strStr().

Returns the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

**Difficulty:** `Easy`

**Tags:** `String` `Tow Pointers`

### Solution One

[从头到尾彻底理解 KMP（2014 年 8 月 22 日版）](http://blog.csdn.net/v_july_v/article/details/7041827)

```c++
class Solution {
public:
    int strStr(string haystack, string needle) {
        if (needle.empty())
        {
            return 0;
        }
        else if (haystack.empty())
        {
            return -1;
        }
        getNext(needle);	// Generate next array
        int posP = 0, posT = 0;
        int lengthP = needle.size();
        int lengthT = haystack.size();
        while (posP < lengthP && posT < lengthT)
        {
            if (posP == -1 || haystack[posT] == needle[posP])
            {
                posP++; posT++;
            }
            else
            {
                posP = next[posP];
            }
        }
        if (posP < lengthP)
        {
            return -1;
        }
        else
        {
            return posT - lengthP;
        }
    }
private:
    vector<int> next;
    void getNext(const string &needle)
    {
        int j = 0;
        int k = -1;
        next.resize(needle.size());
        next[0] = -1;
        while (j < needle.size() - 1)
        {
            if (k == -1 || needle[j] == needle[k])
            {
                j++; k++;
                next[j] = k;
            }
            else
            {
                k = next[k];
            }
        }
    }
};
```
