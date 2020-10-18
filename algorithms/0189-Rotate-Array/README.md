## [189. Rotate Array](https://leetcode.com/problems/rotate-array/#/description)

### Description

Rotate an array of _n_ elements to the right by _k_ steps.

For example, with _n_ = 7 and _k_ = 3, the array `[1,2,3,4,5,6,7]` is rotated to `[5,6,7,1,2,3,4]`.

**Note:**
Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.

**Hint:**
Could you do it in-place with O(1) extra space?

Related problem: [Reverse Words in a String II](https://leetcode.com/problems/reverse-words-in-a-string-ii/)

**Difficulty:** `Easy`

**Tags:** `Array`

### Solution One

```c++
class Solution {
public:
    void rotate(vector<int>& nums, int k) {
        size_t n = k % nums.size();
        if (n == 0)
        {	// When n is zero, It means that k is zero or a multiple of nums.size()
            // So we have nothing to do
            return;
        }
        size_t i = 0;
        vector<bool> isViewed(nums.size(), false);
        while (!isViewed[i])
        {
            size_t p = (i + n) % nums.size();
            int preVal = nums[i];
            while (p != i)
            {
                int tmp = nums[p];
                nums[p] = preVal;
                isViewed[p] = true;
                preVal = tmp;
                p = (p+n) % nums.size();
            }
            nums[i] = preVal;
            isViewed[i] = true;
            ++i;
        }
    }
};
```

### Other Solution

[189. Rotate Array - Solution](https://leetcode.com/problems/rotate-array/#/solution)
