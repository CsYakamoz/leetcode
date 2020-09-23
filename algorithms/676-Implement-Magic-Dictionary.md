## [676. Implement Magic Dictionary](https://leetcode.com/problems/implement-magic-dictionary/description/)

### Description

Implement a magic directory with `buildDict`, and `search` methods.

For the method `buildDict`, you'll be given a list of non-repetitive words to build a dictionary.

For the method `search`, you'll be given a word, and judge whether if you modify **exactly** one character into **another** character in this word, the modified word is in the dictionary you just built.

**Example 1:**

```
Input: buildDict(["hello", "leetcode"]), Output: Null
Input: search("hello"), Output: False
Input: search("hhllo"), Output: True
Input: search("hell"), Output: False
Input: search("leetcoded"), Output: False
```

**Note:**

1. You may assume that all the inputs are consist of lowercase letters `a-z`.
2. For contest purpose, the test data is rather small by now. You could think about highly efficient algorithm after the contest.
3. Please remember to **RESET** your class variables declared in class MagicDictionary, as static/class variables are **persisted across multiple test cases**. Please see [here](https://leetcode.com/faq/#different-output) for more details.



**Difficult:** `Medium`

**Tags:** `Hash Table` `Trie`



### Solution One

```c++
class MagicDictionary
{
  private:
    unordered_map<size_t, unordered_set<string>> hash;

  public:
    /** Initialize your data structure here. */
    MagicDictionary()
    {
    }

    /** Build a dictionary through a list of words */
    void buildDict(vector<string> dict)
    {
        hash.clear();
        for (auto &word : dict)
        {
            hash[word.size()].insert(word);
        }
    }

    /** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
    bool search(string word)
    {
        for (auto &s : hash[word.size()])
        {
            int count = 0;
            for (size_t i = 0; i < s.size() && count < 2; i++)
            {
                if (s[i] != word[i])
                    count++;
            }
            if (count == 1)
                return true;
        }
        return false;
    }
};

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * MagicDictionary obj = new MagicDictionary();
 * obj.buildDict(dict);
 * bool param_2 = obj.search(word);
 */
```



### Solution Two - In Top Solutions

```c++


class MagicDictionary {
public:
    /** Initialize your data structure here. */
    vector<string> vect;
    MagicDictionary() {
        
    }
    
    ~MagicDictionary() {
        vector<string>().swap(vect);
    }
    
    /** Build a dictionary through a list of words */
    void buildDict(vector<string> dict) {
        vect = dict;
    }
    
    /** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
    bool search(string word) {
        int len = word.size();
        for(auto e:vect){
            if(e.size() == len){
                for(int i=0; i<len; i++){
                    if(e[i] != word[i]){
                        if(e.substr(i+1) == word.substr(i+1)){
                            
                            return true;
                        }
                        break;
                    }
                }
            }
        }
        return false;
    }
};

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * MagicDictionary obj = new MagicDictionary();
 * obj.buildDict(dict);
 * bool param_2 = obj.search(word);
 */
```



