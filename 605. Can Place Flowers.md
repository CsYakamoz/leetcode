## [605. Can Place Flowers](https://leetcode.com/problems/can-place-flowers/#/description)

### Description

Suppose you have a long flowerbed in which some of the plots are planted and some are not. However, flowers cannot be planted in adjacent plots - they would compete for water and both would die.

Given a flowerbed (represented as an array containing 0 and 1, where 0 means empty and 1 means not empty), and a number **n**, return if **n** new flowers can be planted in it without violating the no-adjacent-flowers rule.

**Example 1:**

```
Input: flowerbed = [1,0,0,0,1], n = 1
Output: True
```

**Example 2:**

```
Input: flowerbed = [1,0,0,0,1], n = 2
Output: False
```

**Note:**

1. The input array won't violate no-adjacent-flowers rule.
2. The input array size is in the range of [1, 20000].
3. **n** is a non-negative integer which won't exceed the input array size.



**Difficult:** `Easy`

**Tags:** `Array`



### Solution One

考虑两种情况：

`length`代表空盆子的数量

1. `flowerbed`一朵花都没有，此时可放`length / 2 + length % 2`
2. `flowerbed`至少有一朵花：则要考虑三个部分：
   1. 第一朵花的左侧，此时可放`length / 2`
   2. 最后一朵花的右侧，此时可放`length / 2`
   3. 中间两朵花之间，此时可放`length / 2 - !(length % 2)`

此题太多坑了- -||

```c++
class Solution {
public:
    bool canPlaceFlowers(vector<int>& flowerbed, int n) {
        int count = n;
        size_t i = 0;
        size_t lastFlower = flowerbed.size() - 1;
        while (i < flowerbed.size() && flowerbed[i] == 0)
        {	// make i points to the first flower
            i++;
        }
        if (i == flowerbed.size())
        {	// If no flower
            count -= i / 2 + i % 2;
            return count <= 0;
        }
        else
        {
            count -= i / 2;
        }
        while (lastFlower > i && flowerbed[lastFlower] == 0)
        {	// make it points to the last flower
            lastFlower--;
        }
        count -= (flowerbed.size() - lastFlower - 1) / 2;
        while (i < lastFlower && count > 0)
        {
            size_t j = i + 2;
            while (j < lastFlower && flowerbed[j] == 0)
            {	
                j++;
            }
            // The length of empty flowerbed between two flowers
            int length = j - i - 1;
            count -= length / 2 - !(length % 2);
            i = j;
        }
        return count <= 0;
    }
};
```



### Solution Two - In Top Solutions

[simplest c++ code](https://discuss.leetcode.com/topic/91376/simplest-c-code)

```c++
class Solution {
public:
    bool canPlaceFlowers(vector<int>& flowerbed, int n) {
        flowerbed.insert(flowerbed.begin(),0);
        flowerbed.push_back(0);
        for(int i = 1; i < flowerbed.size()-1; ++i)
        {
            if(flowerbed[i-1] + flowerbed[i] + flowerbed[i+1] == 0)
            {
                --n;
                ++i;
            }
                
        }
        return n <=0;
    }
};
```



