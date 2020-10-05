## [211. Add and Search Word - Data structure design](https://leetcode.com/problems/add-and-search-word-data-structure-design/description/)

### Description

Design a data structure that supports the following two operations:

```
void addWord(word)
bool search(word)

```

search(word) can search a literal word or a regular expression string containing only letters `a-z` or `.`. A `.` means it can represent any one letter.

For example:

```
addWord("bad")
addWord("dad")
addWord("mad")
search("pad") -> false
search("bad") -> true
search(".ad") -> true
search("b..") -> true

```

**Note:**
You may assume that all words are consist of lowercase letters `a-z`.

You should be familiar with how a Trie works. If not, please work on this problem: [Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/) first.

**Difficult:** `Medium`

**Tags:** `Backtracking` `Design` `Trie`

### Solution One

```c++
class WordDictionary {
private:
    struct TrieNode {
        TrieNode *childNode[26];
        bool isEndofWord;
        TrieNode() : isEndofWord(false) { memset(childNode, 0, sizeof(childNode)); }
    };

    TrieNode *root;

    bool search(string &word, size_t idx, TrieNode *root)
    {
        auto cur = root;

        while (idx < word.size() && word[idx] != '.') {
            size_t index = word[idx] - 'a';
            cur = cur->childNode[index];
            if (cur == nullptr)
                return false;
            idx++;
        }

        if (idx == word.size())
            return cur->isEndofWord;

        for (size_t i = 0; i < 26; i++)
            if (cur->childNode[i] != nullptr && search(word, idx + 1, cur->childNode[i]))
                return true;

        return false;
    }
public:
    /** Initialize your data structure here. */
    WordDictionary(): root(new TrieNode()) { }

    /** Adds a word into the data structure. */
    void addWord(string word)
    {
        auto cur = root;
        for (size_t i = 0; i < word.size(); i++) {
            size_t index = word[i] - 'a';
            if (cur->childNode[index] == nullptr)
                cur->childNode[index] = new TrieNode();
            cur = cur->childNode[index];
        }

        cur->isEndofWord = true;
    }

    /** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
    bool search(string word)
    {
        return search(word, 0, root);
    }
};
```

### Solution Two - In Top Solutions

```c++
class WordDictionary {
public:
    /** Initialize your data structure here. */
    WordDictionary() {

    }

    /** Adds a word into the data structure. */
    void addWord(string word) {
        map[(int)word.length()].insert(word);
    }

    /** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
    bool search(string word) {
        int len = word.length();

        for (auto& s: map[len]) {
            bool flag = true;
            for (int i=0; i<s.length(); i++) {
                if (word[i] == '.')
                    continue;
                if (word[i] != s[i]) {
                    flag = false;
                    break;
                }
            }
            if (flag)
                return true;
        }

        return false;
    }

private:
    unordered_map<int, unordered_set<string>> map;
};
```
