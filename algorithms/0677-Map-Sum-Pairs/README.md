## [677. Map Sum Pairs](https://leetcode.com/problems/map-sum-pairs/description/)

### Description

Implement a MapSum class with `insert`, and `sum` methods.

For the method `insert`, you'll be given a pair of (string, integer). The string represents the key and the integer represents the value. If the key already existed, then the original key-value pair will be overridden to the new one.

For the method `sum`, you'll be given a string representing the prefix, and you need to return the sum of all the pairs' value whose key starts with the prefix.

**Example 1:**

```
Input: insert("apple", 3), Output: Null
Input: sum("ap"), Output: 3
Input: insert("app", 2), Output: Null
Input: sum("ap"), Output: 5
```

**Difficult:** `Medium`

**Tags:** `Trie`

### Solution One

```c++
class MapSum {
private:
    struct TrieNode {
        TrieNode *childNode[26];
        bool isEndofWord;
        int val;
        TrieNode() :isEndofWord(false), val(0) { memset(childNode, 0, sizeof(childNode)); }
    };

    TrieNode *root;

    void helper(TrieNode *root, int &val)
    {
        if (root == nullptr) return;

        if (root->isEndofWord)
            val += root->val;

        for (size_t i = 0; i < 26; i++)
            helper(root->childNode[i], val);
    }
public:
    /** Initialize your data structure here. */
    MapSum(): root(new TrieNode()) {}

    void insert(string key, int val)
    {
        // empty string is not a prefix and its val is always zero
        if (key.empty()) return;

        auto cur = root;

        for (size_t i = 0; i < key.size(); i++) {
            size_t index = key[i] - 'a';
            if (cur->childNode[index] == nullptr)
                cur->childNode[index] = new TrieNode();
            cur = cur->childNode[index];
        }

        cur->isEndofWord = true;
        cur->val = val;
    }

    int sum(string prefix)
    {
        int res = 0;
        auto cur = root;

        for (size_t i = 0; i < prefix.size() && cur != nullptr; i++) {
            if (!cur->isEndofWord)
                res += cur->val;
            size_t index = prefix[i] - 'a';
            cur = cur->childNode[index];
        }

        helper(cur, res);
        return res;
    }
};
```

### Solution Two - In Top Solutions

```c++
class TrieNode{
public:
    unordered_map<char, TrieNode*> children;
    bool is_word;
    int val;

    TrieNode(int n){
        is_word = false;
        val = n;
    }
};


class MapSum {
public:
    TrieNode* root;
    unordered_map<string, int> table;

    /** Initialize your data structure here. */
    MapSum() {
        root = new TrieNode(0);
    }


    void insert(string key, int val) {
        int diff;
        if(table.find(key) == table.end()){
            table[key] = val;
            diff = val;
        }
        else{
            diff = val - table[key];
        }

        TrieNode* p = root;
        for(int i = 0; i < key.size(); i++){
            if(p->children.find(key[i]) == p->children.end()){
                TrieNode* temp = new TrieNode(0);
                (p->children)[key[i]] = temp;
            }
            (p->children)[key[i]]->val += diff;
            p = (p->children)[key[i]];
        }
        p->is_word = true;

    }

    int sum(string prefix) {
        TrieNode* p = root;

        for(int i = 0; i < prefix.size(); i++){
            if(p->children.find(prefix[i]) == p->children.end()){
                return 0;
            }
            p = (p->children)[prefix[i]];
        }
        return p->val;
        /*int sum = 0;
        for(auto iter = table.begin(); iter != table.end(); iter++){
            if(iter->first.find(prefix) == 0){
                sum+=iter->second;
            }
        }
        return sum;*/
    }
};
```
