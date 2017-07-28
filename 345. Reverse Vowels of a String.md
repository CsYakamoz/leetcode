## [345. Reverse Vowels of a String](https://leetcode.com/problems/reverse-vowels-of-a-string/#/description)

### Description

Write a function that takes a string as input and reverse only the vowels of a string.

**Example 1:**

Given s = "hello", return "holle".

**Example 2:**
Given s = "leetcode", return "leotcede".

**Note:**
The vowels does not include the letter "y".



**Difficult:** `Easy`

**Tags:** `Two Pointers` `String`



### Solution One

```c++
class Solution {
public:
    string reverseVowels(string s) {
        set<char> dict({ 'a','e','i','o','u','A','E','I','O','U' });
        vector<int> index;
        stack<char> stack;
        for (size_t i = 0; i < s.size(); i++)
        {
            if (dict.find(s[i]) != dict.end())
            {
                stack.push(s[i]);
                index.push_back(i);
            }
        }
        int i = 0;
        while (!stack.empty())
        {
            s[index[i++]] = stack.top();
            stack.pop();
        }
        return s;
    }
};
```



### Solution Two

```c++
class Solution
{
  public:
    string reverseVowels(string s)
    {
        int front = 0;
        int back = s.size() - 1;
        set<char> dict{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'};
        while (front < back)
        {
            while (front < back && dict.find(s[front]) == dict.end())
            {
                front++;
            }
            while (front < back && dict.find(s[back]) == dict.end())
            {
                back--;
            }
            swap(s[front++], s[back--]);
        }
        return s;
    }
};
```


