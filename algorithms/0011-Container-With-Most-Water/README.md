## [11. Container With Most Water](https://leetcode.com/problems/container-with-most-water/#/description)

### Description

Given _n_ non-negative integers _a1_, _a2_, ..., _an_, where each represents a point at coordinate (_i_, _ai_). _n_ vertical lines are drawn such that the two endpoints of line _i_ is at (_i_, _ai_) and (_i_, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and _n_ is at least 2.

**Difficulty:** `Medium`

**Tags:** `Array` `Two Pointers`

### Solution One In Top Solutions

[Simple and fast c++/c with explanation](https://discuss.leetcode.com/topic/16754/simple-and-fast-c-c-with-explanation)

Start by evaluating the widest container, using the first and the last line. All other possible containers are less wide, so to hold more water, they need to be higher. Thus, after evaluating that widest container, skip lines at both ends that don't support a higher height. Then evaluate that new container we arrived at. Repeat until there are no more possible containers left.

```c++
class Solution {
public:
    int maxArea(vector<int>& height) {
        int water = 0;
        int i = 0, j = height.size() - 1;
        while (i < j) {
            int h = min(height[i], height[j]);
            water = max(water, (j - i) * h);
            while (height[i] <= h && i < j) i++;
            while (height[j] <= h && i < j) j--;
        }
        return water;
    }
};
```

### Solution Two In Top Solutions

与**One**一样，但该方法在 LeetCode 上效率更高

原因**猜测**为：思路一中的第 9、10 行虽然跳过了很多 i、j，但循环条件的判定需要更多的时间。

```c++
class Solution {
public:
    int maxArea(vector<int>& height) {


        int i = 0; int j = height.size()-1;
        int water = 0;

        while (i<j){
            int shortBoard = min(height[i],height[j]);
            water = max(water,shortBoard*(j-i));
            (shortBoard == height[i])?(++i):(--j);
        }
        return water;
    }
};
```
