## [599. Minimum Index Sum of Two Lists](https://leetcode.com/problems/minimum-index-sum-of-two-lists/#/description)

### Description

Suppose Andy and Doris want to choose a restaurant for dinner, and they both have a list of favorite restaurants represented by strings.

You need to help them find out their **common interest** with the **least list index sum**. If there is a choice tie between answers, output all of them with no order requirement. You could assume there always exists an answer.

**Example 1:**

```
Input:
["Shogun", "Tapioca Express", "Burger King", "KFC"]
["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
Output: ["Shogun"]
Explanation: The only restaurant they both like is "Shogun".

```

**Example 2:**

```
Input:
["Shogun", "Tapioca Express", "Burger King", "KFC"]
["KFC", "Shogun", "Burger King"]
Output: ["Shogun"]
Explanation: The restaurant they both like and have the least index sum is "Shogun" with index sum 1 (0+1).

```

**Note:**

1. The length of both lists will be in the range of [1, 1000].
2. The length of strings in both lists will be in the range of [1, 30].
3. The index is starting from 0 to the list length minus 1.
4. No duplicates in both lists.

**Difficulty:** `Easy`

**Tags:** `Hash Table`

### Solution One

首先将`list1`添加到 map 中

接着，对于`list2`中的任意元素，查看在 map 中是否存在

若存在，根据当前索引和`indexSum`与其索引和的比较，进行相应操作

```c++
class Solution {
public:
    vector<string> findRestaurant(vector<string>& list1, vector<string>& list2) {
        if (list1.size() > list2.size())
        {
            return findRestaurant(list2, list1);
        }
        unordered_map<string, int> hash;
        vector<string> res;
        int indexSum = INT_MAX;
        for (size_t i = 0; i < list1.size(); i++)
        {
            hash[list1[i]] = i;
        }
        for (size_t i = 0; i < list2.size(); i++)
        {
            auto iter = hash.find(list2[i]);
            if (iter != hash.end())
            {	// list2[i] is common interest
                if (indexSum > iter->second + i)
                {	// indexSum is not least list index sum
                    indexSum = iter->second + i;
                    res.clear();
                    res.push_back(list2[i]);
                }
                else if (indexSum == iter->second + i)
                {
                    res.push_back(list2[i]);
                }
            }
        }
        return res;
    }
};
```

### Other Solutions

[599. Minimum Index Sum of Two Lists - Solution](https://leetcode.com/problems/minimum-index-sum-of-two-lists/solution/)

个人觉得 Approach #2 Without Using HashMap 这思路很新奇
