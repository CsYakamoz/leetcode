## [15. 3Sum](https://leetcode.com/problems/3sum/)

### Description

Given an array *S* of *n* integers, are there elements *a*, *b*, *c* in *S* such that *a* + *b* + *c* = 0? Find all unique triplets in the array which gives the sum of zero.

**Note:** The solution set must not contain duplicate triplets.

```
For example, given array S = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
```

**Difficult:** `Medium`

**Tags:** `Array` `Two Pointers`

### Solution One

该思路因算法效率太低，无法通过 LeetCode 测试

```c++
class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int>> result;
        for (size_t i = 0; i < nums.size()-2; i++)
        {
            for (size_t j = i+1; j < nums.size()-1; j++)
            {
                for (size_t k = j+1; k < nums.size(); k++)
                {
                    if (nums[i] + nums[j] + nums[k] == 0)
                    {
                        vector<int> temp{ nums[i], nums[j], nums[k] };
                        sort(temp.begin(), temp.end());
                        result.push_back(temp);
                    }
                }
            }
        }
        sort(result.begin(), result.end());		// 排序
        auto end_unique = unique(result.begin(), result.end());
        result.erase(end_unique, result.end());	// 删除重复的
        return result;
    }
};
```

### Solution Two - In Top Solutions

`Sotr` `Two Pointers`

首先排序，接着以`-nums[i]`为目标（`target`），在`[i+1, nums.size() - 1]`中寻找两个数的和为`-nums[i]`

使用前后指针（front，back），其和为`sum = nums[front] + nums[back]`

若和大于`target`，则后指针往前移，因为和要变小，只能后指针前移（原因：已排序）

若和小于`target`，则前指针往后移，因为和要变大，只能前指针后移（原因：已排序）

若等于`target`，则将对应索引添加到`result`

```c++
class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int>> result;
          // 排序
        sort(nums.begin(), nums.end());
        for (size_t i = 0; i < nums.size(); i++)
        {
            int target = -nums[i];			// 目标
            int front = i + 1;				// 前下标
            int back = nums.size() - 1;		// 后下标
            while (front < back)
            {
                int sum = nums[front] + nums[back];
                if (sum == target)
                {
                    vector<int> triplet{ nums[i],nums[front],nums[back] };
                    result.push_back(triplet);
                      // 跳过相同的数
                    while (front < back && nums[front] == triplet[1])
                    {
                        ++front;
                    }
                      // 跳过相同的数
                    while (front < back && nums[back] == triplet[2])
                    {
                        --back;
                    }
                }
                else if (sum < target)	// 前下标后移
                {
                    ++front;
                }
                else					// 后下标前移
                {
                    --back;
                }
            }
              // 跳过相同的数
              // 三个数的和为 0，相同的数只可能且最多有两个
              // 若存在，则上面 nums[i] 和 nums[front] 已经有了（原因：已排序！！！)
              // 故此处跳过
            while (i + 1 < nums.size() && nums[i] == nums[i+1])
            {
                ++i;
            }
        }
        return result;
    }
};
```
