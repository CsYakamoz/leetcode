## [593. Valid Square](https://leetcode.com/problems/valid-square/description/)

### Description

Given the coordinates of four points in 2D space, return whether the four points could construct a square.

The coordinate (x,y) of a point is represented by an integer array with two integers.

**Example:**

```
Input: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
Output: True

```

Note:

1. All the input integers are in the range [-10000, 10000].
2. A valid square has four equal sides with positive length and four equal angles (90-degree angles).
3. Input points have no order.



**Difficult:** `Medium`

**Tags:** `Math`



### Solution One

```c++
class Solution {
public:
    bool validSquare(vector<int>& p1, vector<int>& p2, vector<int>& p3, vector<int>& p4) {
        // 取 p1 p2 p3 三个点组成三角形
        // diagonal1 为斜边的向量
        vector<int> diagonal1{ p1[0] - p2[0], p1[1] - p2[1] };
        double longEdge = (p1[0] - p2[0]) * (p1[0] - p2[0]) + (p1[1] - p2[1]) * (p1[1] - p2[1]);
        double shortEdge = (p1[0] - p3[0]) * (p1[0] - p3[0]) + (p1[1] - p3[1]) * (p1[1] - p3[1]);
        double thirdEdge = (p2[0] - p3[0]) * (p2[0] - p3[0]) + (p2[1] - p3[1]) * (p2[1] - p3[1]);

        // 检查 p1 p2 p3 能否形成等腰直角三角形
        if (shortEdge != longEdge)
        {   // 得到腰和斜边的长度，此时第三条边的长度必须等于腰的长度
            if (shortEdge > longEdge)
            {
                swap(shortEdge, longEdge);
                diagonal1[0] = p1[0] - p3[0];
                diagonal1[1] = p1[1] - p3[1];
            }
            if (thirdEdge != shortEdge || shortEdge + shortEdge != longEdge)
                return false;
        }
        else
        {   // 只得到腰的长度，此时第三条边的长度必须大于腰的长度
            if (thirdEdge <= shortEdge || shortEdge + shortEdge != thirdEdge)
                return false;
            longEdge = thirdEdge;
            diagonal1[0] = p2[0] - p3[0];
            diagonal1[1] = p2[1] - p3[1];
        }

       // 检查 p4 与其它三个点的距离是否存在有 longEdge
        vector<vector<int>> points = { p1,p2,p3 };
        vector<int> p;
        bool found = false;
        for (size_t i = 0; i < points.size(); i++)
        {
            double length = (p4[0] - points[i][0]) * (p4[0] - points[i][0]) +
                (p4[1] - points[i][1]) * (p4[1] - points[i][1]);
            if (length != shortEdge && length != longEdge) return false;
            if (length == longEdge)
            {
                found = true;
                p = points[i];
                break;
            }
        }
        if (found == false) return false;

        // 检查对角线是否垂直
        vector<int> diagonal2{ p4[0] - p[0], p4[1] - p[1] };
        if (diagonal1[0] * diagonal2[0] + diagonal1[1] * diagonal2[1] != 0)
            return false;
        return true;
    }
};
```



### Solution Two - In Top Solutions

[C++ 3 lines (unordered_set)](https://discuss.leetcode.com/topic/89985/c-3-lines-unordered_set)

> If we calculate all distances between 4 points, 4 smaller distances should be equal (sides), and 2 larger distances should be equal too (diagonals). As an optimization, we can compare squares of the distances, so we do not have to deal with the square root and precision loss.
>
> Therefore, our set will only contain 2 unique distances in case of square (beware of the zero distance though).

```c++
int d(vector<int>& p1, vector<int>& p2) {
    return (p1[0] - p2[0]) * (p1[0] - p2[0]) + (p1[1] - p2[1]) * (p1[1] - p2[1]);
}
bool validSquare(vector<int>& p1, vector<int>& p2, vector<int>& p3, vector<int>& p4) {
    unordered_set<int> s({ d(p1, p2), d(p1, p3), d(p1, p4), d(p2, p3), d(p2, p4), d(p3, p4) });
    return !s.count(0) && s.size() == 2;
}
```



### Solution Three - In Top Solutions

```c++
class Solution {
public:
    bool validSquare(vector<int>& p1, vector<int>& p2, vector<int>& p3, vector<int>& p4) {
        unordered_map<long, int> mp;
        mp[distance(p1, p2)]++;
        mp[distance(p2, p3)]++;
        mp[distance(p3, p4)]++;
        mp[distance(p4, p1)]++;
        mp[distance(p1, p3)]++;
        mp[distance(p2, p4)]++;
        if(mp.size() != 2) return false;
        auto it = mp.begin();
        int t1 = it->second;
        it++;
        int t2 = it->second;
        return (t1 == 2 && t2 == 4) || (t1 == 4 && t2 == 2);
    }
    long distance(vector<int> a, vector<int> b) {
        return pow(a[0]-b[0], 2) + pow(a[1]-b[1], 2);
    }
};
```

