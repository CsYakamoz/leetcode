## [30. Substring with Concatenation of All Words](https://leetcode.com/problems/substring-with-concatenation-of-all-words/description/#)

### Description

You are given a string, **s**, and a list of words, **words**, that are all of the same length. Find all starting indices of substring(s) in **s** that is a concatenation of each word in **words** exactly once and without any intervening characters.

For example, given:
**s**: `"barfoothefoobarman"`
**words**: `["foo", "bar"]`

You should return the indices: `[0,9]`.
(order does not matter).



**Difficult:** `Hard`

**Tags:** `Hash Table` `Two Pointers` `String`



### Solution One

```c++
class Solution
{
  private:
    bool check(const string &str, unordered_map<string, int> hash, const size_t length)
    {
        for (size_t i = 0; i < str.size(); i += length)
        {
            auto iter = hash.find(str.substr(i, length));
            if (iter == hash.end() || iter->second == 0)
                return false;
            else
                iter->second--;
        }

        return true;
    }

  public:
    vector<int> findSubstring(string s, vector<string> &words)
    {
        int counter = 0;
        vector<int> res;
        vector<int> vec(128, 0);
        unordered_map<string, int> hash;

        for (auto &word : words)
        {
            hash[word]++;
            counter += word.size();
            for (auto c : word)
                vec[c]++;
        }

        size_t length = counter;
        size_t left = 0;
        size_t right = 0;

        while (right < s.size())
        {
            if (vec[s[right++]]-- > 0)
                counter--;

            while (counter == 0)
            {
                if (length == right - left &&
                    check(s.substr(left, length), hash, words[0].size()))
                {
                    res.push_back(left);
                }

                if (vec[s[left++]]++ == 0)
                    counter++;
            }
        }
        return res;
    }
};
```



