## [648. Replace Words](https://leetcode.com/problems/replace-words/description/)

## Description

In English, we have a concept called `root`, which can be followed by some other words to form another longer word - let's call this word `successor`. For example, the root `an`, followed by `other`, which can form another word `another`.

Now, given a dictionary consisting of many roots and a sentence. You need to replace all the `successor` in the sentence with the `root`forming it. If a `successor` has many `roots` can form it, replace it with the root with the shortest length.

You need to output the sentence after the replacement.

**Example 1:**

```
Input: dict = ["cat", "bat", "rat"]
sentence = "the cattle was rattled by the battery"
Output: "the cat was rat by the bat"

```

**Note:**

1. The input will only have lower-case letters.
2. 1 <= dict words number <= 1000
3. 1 <= sentence words number <= 1000
4. 1 <= root length <= 100
5. 1 <= sentence words length <= 1000

**Difficult:** `Medium`

**Tags:** `Hash Table` `Trie`

### Solution One

```c++
class Solution {
private:
    struct TrieNode {
        bool isEndOfWord;
        TrieNode *next[26];
        TrieNode() {
            isEndOfWord = false;
            memset(next, 0, sizeof(next));
        }
    };


    void insert(const string &str) {
        TrieNode *cur = &root;

        for (auto c : str) {
            size_t idx = c - 'a';
            if (cur->next[idx] == nullptr) {
                cur->next[idx] = new TrieNode();
            }
            cur = cur->next[idx];
        }

        cur->isEndOfWord = true;
    }

    string search(const string &str) {
        TrieNode *cur = &root;
        string res;
        bool find = true;

        for (auto c : str) {
            size_t idx = c - 'a';
            if (cur->next[idx] == nullptr) {
                find = false;
                break;
            }
            res.push_back(c);
            cur = cur->next[idx];
            if (cur->isEndOfWord) {
                break;
            }
        }

        return find ? res : str;
    }

    TrieNode root;

public:
    string replaceWords(vector<string>& dict, string sentence) {
        istringstream iss(sentence);
        string res;
        string str;

        for (auto &str : dict) {
            insert(str);
        }

        while (iss >> str) {
            res += search(str) + ' ';
        }
        res.pop_back();

        return res;
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    string replaceWords(vector<string>& dict, string sentence) {
        string result, cur;
        vector<vector<string>> aux(26);
        sort(dict.begin(), dict.end(), [](const string &a, const string &b) {
            return a.size() < b.size();
        });

        for (auto word : dict) {
            if (word.empty())
                continue;
            aux[word[0]-'a'].push_back(word);
        }

        istringstream iss(sentence);

        string word;
        while (iss>>word) {
            for (auto prefix : aux[word[0]-'a']) {
                int len = prefix.size();
                if (word.substr(0,len) == prefix) {
                    word = prefix;
                    break;
                }
            }
            result += word +' ';
        }

        result.pop_back();
        return result;
    }
};
```

### Solution Three - In Top Solutions

[C++ trie with optimizations (50 ms)](https://discuss.leetcode.com/topic/96835/c-trie-with-optimizations-50-ms)
