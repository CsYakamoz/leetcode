## [496. Next Greater Element I](https://leetcode.com/problems/next-greater-element-i/#/description)

### Description

You are given two arrays **(without duplicates)** `nums1` and `nums2` where `nums1`’s elements are subset of `nums2`. Find all the next greater numbers for `nums1`'s elements in the corresponding places of `nums2`.

The Next Greater Number of a number **x** in `nums1` is the first greater number to its right in `nums2`. If it does not exist, output -1 for this number.

**Example 1:**

```
Input: nums1 = [4,1,2], nums2 = [1,3,4,2].
Output: [-1,3,-1]
Explanation:
    For number 4 in the first array, you cannot find the next greater number for it in the second array, so output -1.
    For number 1 in the first array, the next greater number for it in the second array is 3.
    For number 2 in the first array, there is no next greater number for it in the second array, so output -1.
```

**Example 2:**

```
Input: nums1 = [2,4], nums2 = [1,2,3,4].
Output: [3,-1]
Explanation:
    For number 2 in the first array, the next greater number for it in the second array is 3.
    For number 4 in the first array, there is no next greater number for it in the second array, so output -1.
```

**Note:**

1. All elements in `nums1` and `nums2` are unique.
2. The length of both `nums1` and `nums2` would not exceed 1000.

**Difficult:** `Easy`

**Tags:** `Stack`

### Solution One

对于`findNums`中任意一个元素`findNums[i]`，在`nums`找到对应位置，再从相应位置往右找第一个比`findNums[i]`大的元素

```c++
class Solution {
public:
    vector<int> nextGreaterElement(vector<int>& findNums, vector<int>& nums) {
        vector<int> result(findNums.size());
        for (size_t i = 0; i < findNums.size(); i++)
        {
            int val = -1;;
            auto iter = find(nums.cbegin(), nums.cend(), findNums[i]);
            if (iter != nums.cend())
            {
                iter++;
                while (iter != nums.cend())
                {
                    if (*iter > findNums[i])
                    {
                        val = *iter;
                        break;
                    }
                    ++iter;
                }
            }
            result[i] = val;
        }
        return result;
    }
};
```

### Solution Two

将`nums`以值 - 索引对添加到 Map 中

然后对于`findNums`中的任意元素`findNums[i]`，可以容易获得其在`nums`中对应位置，此时往右找第一个比`findNums[i]`大的元素

```c++
class Solution {
public:
    vector<int> nextGreaterElement(vector<int>& findNums, vector<int>& nums) {
        vector<int> result(findNums.size());
        map<int, int> m;
        for (size_t i = 0; i < nums.size(); i++)
        {
            m[nums[i]] = i;
        }
        for (size_t i = 0; i < findNums.size(); i++)
        {
            int val = -1;
            int j = m[findNums[i]];
            ++j;
            while (j != nums.size())
            {
                if (nums[j] > findNums[i])
                {
                    val = nums[j];
                    break;
                }
                j++;
            }
            result[i] = val;
        }
        return result;
    }
};
```

### Solution Three:

思路与**Two**相似，不过这里是将`findNums`以 值-索引对添加到 Map 中

```c++
class Solution {
public:
    vector<int> nextGreaterElement(vector<int>& findNums, vector<int>& nums) {
        vector<int> result(findNums.size());
        map<int, int> m;
        for (size_t i = 0; i < findNums.size(); i++)
        {
            m[findNums[i]] = i;
        }
        for (size_t i = 0; i < nums.size(); i++)
        {
            if (m.find(nums[i]) != m.cend())
            {
                int j = i + 1;
                int val = -1;
                while (j != nums.size())
                {
                    if (nums[j] > nums[i])
                    {
                        val = nums[j];
                        break;
                    }
                    j++;
                }
                result[ m[ nums[i] ] ] = val;
            }
        }
        return result;
    }
};
```

### Solution Four: In Top Solutions

关键代码为第 7 行到第 15 行

题目要求是对于`findNums`中的元素，在`nums`中对应位置的右侧寻找第一个比`findNums[i]`大的元素

可以转换为：

对于`nums[i]`这个元素，在该元素右侧寻找第一个比`nums[i]`大的元素

1. 如果栈非空且`val`大于栈顶，说明`val`是第一个比栈顶大的数，所以有`m[s.top()] = val`，这里使用循环的原因是栈顶找到了，所以出栈，新的栈顶可能也小于`val`，此时`val`也是第一个比栈顶大的数
2. 否则意味着栈为空或者`val`小于栈顶，不做任何事
3. 将`val`进栈

```c++
class Solution {
public:
    vector<int> nextGreaterElement(vector<int>& findNums, vector<int>& nums) {
        map<int, int> m;
        stack<int> s;
        vector<int> result(findNums.size(), -1);
        for (auto const &val : nums)
        {
            while (!s.empty() && s.top() < val)
            {
                m[s.top()] = val;
                s.pop();
            }
            s.push(val);
        }
        for (size_t i = 0; i < findNums.size(); i++)
        {
            if (m.find(findNums[i]) != m.end())
            {
                result[i] = m[findNums[i]];
            }
        }
        return result;
    }
};
```
