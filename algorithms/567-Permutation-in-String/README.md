## [567. Permutation in String](https://leetcode.com/problems/permutation-in-string/description/)

### Description

Given two strings **s1** and **s2**, write a function to return true if **s2** contains the permutation of **s1**. In other words, one of the first string's permutations is the **substring** of the second string.

**Example 1:**

```
Input:s1 = "ab" s2 = "eidbaooo"
Output:True
Explanation: s2 contains one permutation of s1 ("ba").

```

**Example 2:**

```
Input:s1= "ab" s2 = "eidboaoo"
Output: False

```

**Note:**

1. The input strings only contain lower case letters.
2. The length of both given strings is in range [1, 10,000].

**Difficult:** `Medium`

**Tags:** `Two Pointers`

### Solution One

思路来源于 [438. Find All Anagrams in a String](https://leetcode.com/problems/find-all-anagrams-in-a-string/description/) 中的 Top Solutions - [Shortest/Concise JAVA O(n) Sliding Window Solution](https://discuss.leetcode.com/topic/64434/shortest-concise-java-o-n-sliding-window-solution)

```c++
class Solution {
public:
    bool checkInclusion(string s1, string s2) {
        int length = s1.size();
        vector<int> hash(26, 0);

        for (auto i : s1)
            hash[i - 'a']++;

        size_t left = 0;
        size_t right = 0;
        int count = length;
        while (right < s2.size())
        {
            if (hash[s2[right++] - 'a']-- > 0)
                count--;

            if (count == 0)
                return true;

            if (right - left == length && hash[s2[left++] - 'a']++ >= 0)
                count++;
        }
        return false;
    }
};
```

### Solution Two

[Here is a 10-line template that can solve most 'substring' problems](https://discuss.leetcode.com/topic/30941/here-is-a-10-line-template-that-can-solve-most-substring-problems/2)

```c++
class Solution {
public:
    bool checkInclusion(string s1, string s2) {
        vector<int> hash(128, 0);

        for (auto c : s1) {
            hash[c]++;
        }

        int left = 0;
        int right = 0;
        int counter = s1.size();

        while (right < s2.size()) {
            if (hash[s2[right++]]-- > 0) {
                --counter;
            }

            while(counter == 0) {
                if (right - left == s1.size()) {
                    return true;
                }

                if (hash[s2[left++]]++ == 0) {
                    counter++;
                }
            }
        }

        return false;
    }
};
```
