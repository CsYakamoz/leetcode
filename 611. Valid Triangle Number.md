## [611. Valid Triangle Number](https://leetcode.com/problems/valid-triangle-number/description/)

### Description

Given an array consists of non-negative integers, your task is to count the number of triplets chosen from the array that can make triangles if we take them as side lengths of a triangle.

**Example 1:**

```
Input: [2,2,3,4]
Output: 3
Explanation:
Valid combinations are: 
2,3,4 (using the first 2)
2,3,4 (using the second 2)
2,2,3
```

**Note:**

1. The length of the given array won't exceed 1000.
2. The integers in the given array are in the range of [0, 1000].



**Difficult:** `Medium`

**Tags:** `Array`



### Solution One

[611. Valid Triangle Number - Solution - #Linear Search](https://leetcode.com/problems/valid-triangle-number/solution/#approach-3-linear-scan-accepted)

```c++
class Solution {
public:
    int triangleNumber(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        int res = 0;
        int length = nums.size();
        for (int i = 0; i < length - 2; i++)
        {
            if (nums[i] == 0)
            {
                continue;
            }
            int k = i + 2;
            for (int j = i + 1; j < length - 1; j++)
            {
                while (k < length && nums[i] + nums[j] > nums[k])
                {
                    k++;
                }
                res += k - j - 1;
            }
        }
        return res;
    }
};
```



### Solution Tow

[611. Valid Triangle Number - Solution - #Binary Search](https://leetcode.com/problems/valid-triangle-number/solution/#approach-2-using-binary-search-accepted)

```c++
class Solution {
public:
    int triangleNumber(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        int res = 0;
        int length = nums.size();
        for (int i = 0; i < length - 2; i++)
        {
            if (nums[i] == 0)
            {
                continue;
            }
            int k = i + 2;
            for (int j = i + 1; j < length - 1; j++)
            {
                k = binarySearch(nums, k, length - 1, nums[i] + nums[j]);
                res += k - j - 1;
            }
        }
        return res;
    }

private:
    int binarySearch(vector<int> &nums, int left, int right, int target)
    {
        while (left <= right)
        {
            int mid = (right - left) / 2 + left;
            if (nums[mid] >= target)
            {
                right = mid - 1;
            }
            else
            {
                left = mid + 1;
            }
        }
        return left;
    }
};
```



### Solution Three - In Top Solutions

`Two Pointers`

在已排序好的数组中，选择`a, b, c`（其中a < b < c）

其中对于任意一个 c，a - b < c 这个是肯定满足的，所以我们需要找到  a + b > c

while 循环中，我们已经初始化`left, right`分别为`0, i - 1`，即 c 的左侧

>  **注意：**这里 a、b、c 相当于 nums[left]、nums[right]、nums[i]，a++ 和 b-- 相当于 left++ 和 right--

然后思路为：当确定 b、c 时，寻找 a， 使得 a + b > c

如果`nums[left] + nums[right] > nums[i]`，则代表对于当前 b、c，已经找到一个范围的 a 来满足 a + b > c，该范围有`right - left`个，然后当前 b 已经搜索完毕，所以 `b--`

而此时 a 保持不变（即不变回`0`）的原因是：

(a - 1) + b <= c 同时 a + b > c， 此时 b--，会有 (a - 1) + (b - 1) < c，故得到 (a 的左侧范围) + (b - 1) < c，所以 a 保持不变

```c++
class Solution {
public:
    int triangleNumber(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        int res = 0;
        int length = nums.size();
        for (int i = length - 1; i >= 2; i--)
        {
            int left = 0;
            int right = i - 1;
            while (left < right)
            {
                if (nums[left] + nums[right] > nums[i])
                {
                    res += right - left;
                    right--;
                }
                else
                {
                    left++;
                }
            }
        }
        return res;
    }
};
```



