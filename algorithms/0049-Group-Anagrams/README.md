## [49. Group Anagrams](https://leetcode.com/problems/group-anagrams/description/)

### Description

Given an array of strings, group anagrams together.

For example, given: `["eat", "tea", "tan", "ate", "nat", "bat"]`,
Return:

```
[
  ["ate", "eat","tea"],
  ["nat","tan"],
  ["bat"]
]
```

**Note:** All inputs will be in lower-case.

**Difficulty:** `Medium`

**Tags:** `Hash Table` `String`

### Solution One

```c++
class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        unordered_map<string, vector<string>> hash;
        vector<vector<string>> res;

        for (auto &str : strs) {
            string tmp = str;
            sort(tmp.begin(), tmp.end());
            hash[tmp].push_back(str);
        }

        for (auto iter : hash) {
            res.push_back(iter.second);
        }

        return res;
    }
};
```

### Solution Two - In Top Solutions

[Java beat 100%!!! use prime number](https://discuss.leetcode.com/topic/45639/java-beat-100-use-prime-number)

```java
public static List<List<String>> groupAnagrams(String[] strs) {
   int[] prime = {2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103};//最多10609个z

            List<List<String>> res = new ArrayList<>();
            HashMap<Integer, Integer> map = new HashMap<>();
            for (String s : strs) {
                int key = 1;
                for (char c : s.toCharArray()) {
                    key *= prime[c - 'a'];
                }
                List<String> t;
                if (map.containsKey(key)) {
                    t = res.get(map.get(key));
                } else {
                    t = new ArrayList<>();
                    res.add(t);
                    map.put(key, res.size() - 1);
                }
                t.add(s);
            }
            return res;
    }
```
