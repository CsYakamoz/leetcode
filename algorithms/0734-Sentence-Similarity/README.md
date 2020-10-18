## [734. Sentence Similarity](https://leetcode.com/problems/sentence-similarity/description/)

### Description

Given two sentences `words1, words2` (each represented as an array of strings), and a list of similar word pairs `pairs`, determine if two sentences are similar.

For example, "great acting skills" and "fine drama talent" are similar, if the similar word pairs are `pairs = [["great", "fine"], ["acting","drama"], ["skills","talent"]]`.

Note that the similarity relation is not transitive. For example, if "great" and "fine" are similar, and "fine" and "good" are similar, "great" and "good" are **not** necessarily similar.

However, similarity is symmetric. For example, "great" and "fine" being similar is the same as "fine" and "great" being similar.

Also, a word is always similar with itself. For example, the sentences `words1 = ["great"], words2 = ["great"], pairs = []` are similar, even though there are no specified similar word pairs.

Finally, sentences can only be similar if they have the same number of words. So a sentence like `words1 = ["great"]` can never be similar to `words2 = ["doubleplus","good"]`.

**Note:**

- The length of `words1` and `words2` will not exceed `1000`.
- The length of `pairs` will not exceed `2000`.
- The length of each `pairs[i]` will be `2`.
- The length of each `words[i]` and `pairs[i][j]` will be in the range `[1, 20]`.

**Difficulty:** `Easy`

**Tags:** `Hash Table`

### Solution One

```c++
class Solution
{
  public:
    bool areSentencesSimilar(vector<string> &words1, vector<string> &words2, vector<pair<string, string>> pairs)
    {
        if (words1.size() != words2.size())
            return false;

        unordered_map<string, vector<string>> hash;

        for (auto p : pairs)
            hash[p.first].push_back(p.second);

        for (size_t i = 0; i < words1.size(); i++)
        {
            if (words1[i] == words2[i])
                continue;

            if (hash.find(words1[i]) != hash.end())
            {
                if (isSimilar(words2[i], hash[words1[i]]))
                    continue;
            }
            if (hash.find(words2[i]) != hash.end())
            {
                if (isSimilar(words1[i], hash[words2[i]]))
                    continue;
            }
            return false;
        }

        return true;
    }

  private:
    bool isSimilar(const string &key, const vector<string> &val)
    {
        for (size_t i = 0; i < val.size(); i++)
            if (key == val[i])
                return true;
        return false;
    }
};
```

### Solution Two - In Top Solutions

[[Java/C++] Clean Code](https://discuss.leetcode.com/topic/112040/java-c-clean-code)
