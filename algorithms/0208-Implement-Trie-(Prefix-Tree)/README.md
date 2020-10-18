## [208. Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/description/)

### Description

Implement a trie with `insert`, `search`, and `startsWith` methods.

**Note:**
You may assume that all inputs are consist of lowercase letters `a-z`.

**Difficulty:** `Medium`

**Tags:** `Design` `Trie`

### Solution One

```c++
class Trie {
public:
    /** Initialize your data structure here. */
    Trie(): root(new TrieNode()) {}

    /** Inserts a word into the trie. */
    void insert(string word)
    {
        TrieNode *cur = root;
        for (size_t i = 0; i < word.size(); i++) {
            size_t index = word[i] - 'a';
            if (cur->childNode[index] == nullptr)
                cur->childNode[index] = new TrieNode();
            cur = cur->childNode[index];
        }

        cur->isEndofWord = true;
    }

    /** Returns if the word is in the trie. */
    bool search(string word)
    {
        TrieNode *cur = root;
        for (size_t i = 0; i < word.size(); i++) {
            size_t index = word[i] - 'a';
            cur = cur->childNode[index];
            if (cur == nullptr) return false;
        }

        return cur->isEndofWord;
    }

    /** Returns if there is any word in the trie that starts with the given prefix. */
    bool startsWith(string prefix)
    {
        if (prefix.empty()) return true;

        TrieNode *cur = root;
        for (size_t i = 0; i < prefix.size(); i++) {
            size_t index = prefix[i] - 'a';
            cur = cur->childNode[index];
            if (cur == nullptr) return false;
        }

        return true;
    }

private:
    static constexpr size_t SIZE = 26;

    struct TrieNode {
        TrieNode *childNode[SIZE];
        bool isEndofWord;
        TrieNode():isEndofWord(false)
        {
            for (size_t i = 0; i < SIZE; i++)
                childNode[i] = nullptr;
        }
    };

    TrieNode *root;
};
```

### Solution Two - In Top Solutions

```c++
class TrieNode{
public:
    TrieNode *next[26];
    bool is_word;
    TrieNode(bool b=false):is_word(b){
        memset(next, 0, sizeof(next));
    }
};

class Trie {
    TrieNode *root;
public:
    /** Initialize your data structure here. */
    Trie() {
        root=new TrieNode();
    }

    /** Inserts a word into the trie. */
    void insert(string word) {
        TrieNode *p=root;
        for(int i=0; i<word.size(); i++){
            if(p->next[word[i]-'a']==NULL) p->next[word[i]-'a']=new TrieNode();
            p=p->next[word[i]-'a'];
        }
        p->is_word=true;
    }

    /** Returns if the word is in the trie. */
    bool search(string word) {
        TrieNode *p=find(word);
        return p!=NULL && p->is_word;
    }

    /** Returns if there is any word in the trie that starts with the given prefix. */
    bool startsWith(string prefix) {
        return find(prefix)!=NULL;
    }
private:
    TrieNode* find(string key){
        TrieNode *p=root;
        for(int i=0; i<key.size() && p!=NULL; i++)
            p=p->next[key[i]-'a'];
        return p;
    }
};
```

### Solutions

[208. Implement Trie (Prefix Tree) - Solution](https://leetcode.com/problems/implement-trie-prefix-tree/solution/)
