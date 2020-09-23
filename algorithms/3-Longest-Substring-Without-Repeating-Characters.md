## [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/#/description)

### Description

Given a string, find the length of the **longest substring** without repeating characters.

**Examples:**

Given `"abcabcbb"`, the answer is `"abc"`, which the length is 3.

Given `"bbbbb"`, the answer is `"b"`, with the length of 1.

Given `"pwwkew"`, the answer is `"wke"`, with the length of 3. Note that the answer must be a **substring**, `"pwke"` is a *subsequence* and not a substring.



**Difficult:** `Medium`

**Tags:** `Hash Table` `Two Pointers` `String`



### Solution One

将某个字符$C_i$添加到`result`

判定$C_j$(j > i)是否与`result`中的某个字符相等，若不相等，则将其添加到`result`

否则根据`result`长度与`max`来判断是否更新`max`

```c++
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        if (s.size() == 0 || s.size() == 1)
        {
            return s.size();
        }
        size_t max = 0;
        bool flag = false;
        vector<int> result;
        for (size_t i = 0; i < s.size(); i++)
        {
            result.clear();
            result.push_back(s[i]);
            for (size_t j = i+1; j < s.size(); j++)
            {
                for (size_t k = 0; k < result.size(); k++)
                {
                    if (s[j] == result[k])
                    {
                        flag = true;
                        break;
                    }
                }
                if (flag)
                {
                    flag = false;
                    break;
                }
                result.push_back(s[j]);
            }
            if (max < result.size())
            {
                max = result.size();
            }
        }
        return max;
    }
};
```



### Solution Two

`Hash Table` `Two Pointers`

使用Map的方法，以字符为key，下标为value

对于一个字符序列$C_i,\cdots,C_{j-1}C_j,C_{j+1},\cdots,C_{k-1},C_k,C_{k+1}$(i < j < k)

假设$C_i$~$C_{k-1}$没有重复字符，而$C_k$与$C_i$~$C_{k-1}$中的$C_j$重复，则比较$C_i,\cdots,C_{k-1}$的长度与当前最长长度`length`，以此判断是否更新`length`，接着将Map中以$C_i,\cdots,C_{j-1}$为key的节点删去，同时更新以$C_j$字符为key对应的value（即更新value值为k）

然后下次检测的是$C_{k+1}$是否与$C_{j+1}$~$C_k$中的某个字符重复

```c++
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        map<char, int> m;
        int length = 0;
        string::size_type i = 0, j = i;
        string::size_type temp = 0;
        while (i != s.size())
        {
            if (m.find(s[i]) == m.end())
            {
                m.insert(map<char, int>::value_type(s[i], i));
            }
            else
            {
                length = length < i - j ? i - j : length;
                temp = m.find(s[i])->second;
                while (j != temp)
                {
                    m.erase(s[j]);
                    ++j;
                }
                ++j;
                m.find(s[i])->second = i;
            }
            ++i;
        }
        length = length < i - j ? i - j : length;
        return length;
    }
};
```



### Solution Three - In Top Solutions

`Hash Table` `Two Pointers`

与**Solution Two**类似，只是没有在Map中删除以$C_i,\cdots,C_{j-1}$为key的结点

```c++
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        vector<int> v(256, -1);
        int start = -1;
        int maxLength = 0;
        for (int i = 0; i < s.size(); i++)
        {
            if (v[s[i]] > start)
            {
                start = v[s[i]];
            }
            v[s[i]] = i;
            maxLength = max(maxLength, i - start);
        }
        return maxLength;
    }
};
```

与上面类似，不过这个是找到重复字符才更新`maxLength`，而上面是每次循环都更新

```c++
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        vector<int> flag(256, -1);
        int start = 0, maxLength = 0;
        for (int i = 0; i != s.size(); i++) {
            if (flag[s[i]] >= start) {
                maxLength = max(i - start, maxLength);
                start = flag[s[i]] + 1;
            }
            flag[s[i]] = i;
        }
        return max(s.size() - start, maxLength);
    }
};
```


