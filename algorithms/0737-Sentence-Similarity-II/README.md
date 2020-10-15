## [737. Sentence Similarity II](https://leetcode.com/problems/sentence-similarity-ii/description/)

### Description

Given two sentences `words1, words2` (each represented as an array of strings), and a list of similar word pairs `pairs`, determine if two sentences are similar.

For example, `words1 = ["great", "acting", "skills"]` and `words2 = ["fine", "drama", "talent"]` are similar, if the similar word pairs are `pairs = [["great", "good"], ["fine", "good"], ["acting","drama"], ["skills","talent"]]`.

Note that the similarity relation **is** transitive. For example, if "great" and "good" are similar, and "fine" and "good" are similar, then "great" and "fine" **are similar**.

Similarity is also symmetric. For example, "great" and "fine" being similar is the same as "fine" and "great" being similar.

Also, a word is always similar with itself. For example, the sentences `words1 = ["great"], words2 = ["great"], pairs = []` are similar, even though there are no specified similar word pairs.

Finally, sentences can only be similar if they have the same number of words. So a sentence like `words1 = ["great"]` can never be similar to `words2 = ["doubleplus","good"]`.

**Note:**

- The length of `words1` and `words2` will not exceed `1000`.
- The length of `pairs` will not exceed `2000`.
- The length of each `pairs[i]` will be `2`.
- The length of each `words[i]` and `pairs[i][j]` will be in the range `[1, 20]`.

**Difficult:** `Medium`

**Tags:** `Depth-first Search` `Union Find`

### Solution One - In Top Solutions

```c++
class Solution
{
  public:
    bool areSentencesSimilarTwo(vector<string> &words1, vector<string> &words2, vector<pair<string, string>> pairs)
    {
        if (words1.size() != words2.size())
            return false;

        unordered_map<string, string> hash;
        UF(pairs, hash);

        size_t length = words1.size();

        for (size_t i = 0; i < length; i++)
        {
            if (words1[i] == words2[i])
                continue;

            if (find(hash, words1[i]) != find(hash, words2[i]))
                return false;
        }

        return true;
    }

  private:
    void UF(const vector<pair<string, string>> &pairs, unordered_map<string, string> &hash)
    {
        for (auto &p : pairs)
        {
            string fRoot = find(hash, p.first);
            string sRoot = find(hash, p.second);
            if (fRoot != sRoot)
                hash[fRoot] = sRoot; // Or hash[sRoot] = fRoot;
        }
    }

    string find(unordered_map<string, string> &hash, string word)
    {
        if (hash.find(word) == hash.end())
            return hash[word] = word;

        while (word != hash[word])
            word = hash[word];

        return word;
    }
};
```

### Solution Two - In Top Solutions

[C++ DFS solution](https://discuss.leetcode.com/topic/112146/c-dfs-solution)

```c++
class Solution {
public:
    bool areSentencesSimilarTwo(vector<string>& words1, vector<string>& words2, vector<pair<string, string>> pairs) {
        if (words1.size() != words2.size()) return false;
        unordered_map<string, unordered_set<string>> p;
        for (auto &vp : pairs) {
            p[vp.first].emplace(vp.second);
            p[vp.second].emplace(vp.first);
        }
        for (int i = 0; i < words1.size(); i++) {
            unordered_set<string> visited;
            if (isSimilar(words1[i], words2[i], p, visited)) continue;
            else return false;
        }
        return true;
    }

    bool isSimilar(string& s1, string& s2, unordered_map<string, unordered_set<string>>& p, unordered_set<string>& visited) {
        if (s1 == s2) return true;

        visited.emplace(s1);
        for (auto s : p[s1]) {
            if (!visited.count(s) && isSimilar(s, s2, p, visited))
                return true;
        }

        return false;
    }
};
```
