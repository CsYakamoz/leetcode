## [125. Valid Palindrome](https://leetcode.com/problems/valid-palindrome/#/description)

### Description

Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

For example,
`"A man, a plan, a canal: Panama"` is a palindrome.
`"race a car"` is *not* a palindrome.

**Note:**
Have you consider that the string might be empty? This is a good question to ask during an interview.

For the purpose of this problem, we define empty string as valid palindrome.



**Difficult:** `Easy`

**Tags:** `Two Pointers` `String`



### Solution One

`Two Pointers`

```c++
class Solution {
public:
    bool isPalindrome(string s) {
        int i = 0;
        int j = s.size() - 1;
        while (i < j)
        {
            while (i < j && !isalnum(s[i]))
            {
                i++;
            }
            while (i < j && !isalnum(s[j]))
            {
                j--;
            }
            if (toupper(s[i++]) != toupper(s[j--]))
            {
                return false;
            }
        }
        return true;
    }
};
```



### Solution Two - In Top Solutions

```c++
class Solution {
public:
    bool isPalindrome(string s) {
        string str = "";
        for(char c:s){
            if((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >='0') && (c <= '9')){
                if(c >= 'A' && c <= 'Z'){
                    c = c - 'A' + 'a';
                }
                str += c;
            }

        }

        int i = 0 , j = str.length()-1;
        while(i<= j && str[i] == str[j]){
            i++;
            j--;
        }
        
        if(i >= j)
            return true;
        else
            return false;
    }
};
```



