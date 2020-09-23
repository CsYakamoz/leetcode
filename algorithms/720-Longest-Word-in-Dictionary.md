## [720. Longest Word in Dictionary](https://leetcode.com/problems/longest-word-in-dictionary/description/)

### Description

Given a list of strings `words` representing an English Dictionary, find the longest word in `words` that can be built one character at a time by other words in `words`. If there is more than one possible answer, return the longest word with the smallest lexicographical order.

**Example 1:**

```
Input:
words = ["w","wo","wor","worl", "world"]
Output: "world"
Explanation:
The word "world" can be built one character at a time by "w", "wo", "wor", and "worl".
```

**Example 2:**

```
Input:
words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
Output: "apple"
Explanation:
Both "apply" and "apple" can be built from other words in the dictionary. However, "apple" is lexicographically smaller than "apply".
```

**Note:**

- All the strings in the input will only contain lowercase letters.
- The length of `words` will be in the range `[1, 1000]`.
- The length of `words[i]` will be in the range `[1, 30]`.

**Difficult:** `Easy`

**Tags:** `Hash Table`

### Solution One

```c++
class Solution
{
  private:
    struct TrieNode
    {
        bool isEndOfWord;
        vector<TrieNode *> next;
        TrieNode() : isEndOfWord(false)
        {
            for (size_t i = 0; i < 26; i++)
                next.push_back(nullptr);
        }
    };

    void insert(const string &word)
    {
        TrieNode *cur = &root;

        for (auto c : word)
        {
            size_t idx = c - 'a';
            if (cur->next[idx] == nullptr)
                cur->next[idx] = new TrieNode();
            cur = cur->next[idx];
        }

        cur->isEndOfWord = true;
    }


    void dfs(TrieNode *root, string &str)
    {
        if (str.size() > res.size())
            res = str;

        for (size_t i = 0; i < 26; i++)
        {
            TrieNode *cur = root->next[i];
            if (cur != nullptr && cur->isEndOfWord)
            {
                str.push_back('a' + i);
                dfs(cur, str);
                str.pop_back();
            }
        }
    }

    TrieNode root;
    string res;

  public:
    string longestWord(vector<string> &words)
    {
        string str;

        for (auto &word : words)
            insert(word);

        dfs(&root, str);

        return res;
    }
};
```

### Solution Two - In Top Solutions

[[Java/C++ Clean Code]](https://discuss.leetcode.com/topic/109643/java-c-clean-code)

```c++
class Solution
{
  public:
    string longestWord(vector<string> &words)
    {
        unordered_set<string> hash;
        string res;

        sort(words.begin(), words.end());

        for (auto &word : words)
        {
            if (word.size() == 1 || hash.count(word.substr(0, word.size() - 1)))
            {
                if (res.size() < word.size())
                    res = word;
                hash.insert(word);
            }
        }

        return res;
    }
};
```
