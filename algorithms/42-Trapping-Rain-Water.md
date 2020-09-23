## [42. Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/#/description)

### Description

Given _n_ non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.

For example,
Given `[0,1,0,2,1,0,1,3,2,1,2,1]`, return `6`.

![img](http://www.leetcode.com/static/images/problemset/rainwatertrap.png)

The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. **Thanks Marcos** for contributing this image!

**Difficult:** `Hard`

**Tags:** `Array` `Stack` `String`

### Solution One

首先将`front`，`end`分别定位到第一个、最后一个极大值，对于例图而言，则分别是`[1,1]`，`[10,2]`，原因是第一个极大值的左边和最后一个极大值的右边不能形成凹形状。

接着，对于`front`这个位置的 bar，需要右侧 bar `back`来形成凹形状

如果存在高度比`front`高的 bar，则只能是由它来组成凹形状

如果不存在，则在寻找过程中，找到一个最大值的 bar 来组成凹形状

**注意**：`back - front`可能等于`1`

跟着计算面积加到`area`

最后`back = front`，再重复上述过程

```c++
class Solution {
public:
    int trap(vector<int>& height) {
        // If the number of bar is less than 3, it can not trap rainwater
        if (height.size() < 3) {
            return 0;
        }
        int area = 0;
        size_t front = 0;
        size_t end = height.size() - 1;
        while (front < height.size() - 1 && height[front] < height[front + 1]) {
            front++;	// Point to the first local maximum
        }
        while (end > 1 && height[end] < height[end-1]) {
            end--;		// Point to the last local maximum
        }
        while (front < end) {
            size_t i = front + 1;
            size_t back = i;
            while (i <= end) {
                if (height[i] < height[front]) {
                    back = height[i] > height[back] ? i : back;
                } else {
                    back = i;
                    break;
                }
                i++;
            }
            int h = min(height[front], height[back]);
            size_t j = front + 1;
            while (j < back) {
                area += h - height[j];
                j++;
            }
            front = back;
        }
        return area;
    }
};
```

### Solution Two - In Top Solutions

[Sharing my Java code: O(n) time, O(1) space](https://discuss.leetcode.com/topic/5819/sharing-my-java-code-o-n-time-o-1-space)

```java
public int trap(int[] A) {
    if (A.length < 3) return 0;

    int ans = 0;
    int l = 0, r = A.length - 1;

    // find the left and right edge which can hold water
    while (l < r && A[l] <= A[l + 1]) l++;
    while (l < r && A[r] <= A[r - 1]) r--;

    while (l < r) {
        int left = A[l];
        int right = A[r];
        if (left <= right) {
            // add volum until an edge larger than the left edge
            while (l < r && left >= A[++l]) {
                ans += left - A[l];
            }
        } else {
            // add volum until an edge larger than the right volum
            while (l < r && A[--r] <= right) {
                ans += right - A[r];
            }
        }
    }
    return ans;
}
```

### Solution Three - In Top Solutions

[7 lines C / C++](https://discuss.leetcode.com/topic/18731/7-lines-c-c)

**xthron explanation**:

I have spent a lot of time understanding this solution. I would try to summarise my thoughts here:

Consider this example: [3, X, 5], and we are working after the first iteration. That means l = 1 and r = 2

What is the information that we need to figure out the amount of water that can be on X? We need only two things:

- Maximum water that can be on X
- Value of X

Maximum water is not dependent on the value of X. It has to be obtained from the previous step. In this case, that has to be 3. Notice that the minimum of height[l] and height [r] of the previous step is also 3. This information is tracked by level. And X can only hold water if level is greater than X.

There can be three cases of X.

- X > level and X < 5: then X cannot hold any water, but we need to update level.
  For X = 4:
  lower = 4.
  level = max(4, 3) = 4
  water += (4 - 4) += 0
- X < level, then X can hold water, and we don't have to update level since for the next step maximum water will still be level units.
  For X = 1:
  lower = 1,
  level = max(1, 3) = 3
  water += (3 -1) += 2
- X > 5 > level: In this case lower will change to 5. And level also needs to change. since, if there were more bins between 5 and X, then from this steps point of view, maximum water permissible water will be 5 units.
  For X = 8:
  lower = 5
  level = max(5, 3) = 5
  water += (5 - 5) = 0

Hope this helps.

```c++
class Solution {
public:
    int trap(vector<int>& height) {
        int l = 0, r = height.size()-1, level = 0, water = 0;
        while (l < r) {
            int lower = height[height[l] < height[r] ? l++ : r--];
            level = max(level, lower);
            water += level - lower;
        }
        return water;
    }
};
```
