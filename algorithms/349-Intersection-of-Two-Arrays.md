## [349. Intersection of Two Arrays](https://leetcode.com/problems/intersection-of-two-arrays/#/description)

### Description

Given two arrays, write a function to compute their intersection.

**Example:**
Given _nums1_ = `[1, 2, 2, 1]`, _nums2_ = `[2, 2]`, return `[2]`.

**Note:**

- Each element in the result must be unique.
- The result can be in any order.

**Difficult:** `Easy`

**Tags:** `Binary Search` `Hash Table` `Two Pointers` `Sort`

### Solution One

`Two Pointer` `Sort`

Maybe it's better to change line 12 to `while (i < length && j < length)`

the length is `min(nums1.size(), nums2.size())`

```c++
class Solution {
public:
    vector<int> intersection(vector<int>& nums1, vector<int>& nums2) {
        vector<int> res;
        if (nums1.empty() || nums2.empty()) {
            return res;
        }
        sort(nums1.begin(), nums1.end());
        sort(nums2.begin(), nums2.end());
        size_t i = 0;
        size_t j = 0;
        while (i < nums1.size() && j < nums2.size()) {
            if (nums1[i] == nums2[j]) {
                res.push_back(nums1[i]);
                do {
                    i++;
                } while (i < nums1.size() && nums1[i] == nums1[i-1]);
                do {
                    j++;
                } while (j < nums2.size() && nums2[j] == nums2[j-1]);
            } else if (nums1[i] < nums2[j]) {
                i++;
            } else {
                j++;
            }
        }
        return res;
    }
};
```

### Solution Two

`Set`

```c++
class Solution {
public:
    vector<int> intersection(vector<int>& nums1, vector<int>& nums2) {
        set<int> res;
        // Use longer vector to construct set
        // Because longer vector may have more overlap elements
        if (nums1.size() < nums2.size()) {
            return intersection(nums2, nums1);
        }
        set<int> s(nums1.begin(), nums1.end());
        for (const auto &i : nums2) {
            if (s.find(i) != s.end()) {
                res.insert(i);
            }
        }
        return vector<int>(res.begin(), res.end());
    }
};
```

### Solution Three - In Top Solutions

```c++
class Solution {
public:
    vector<int> intersection(vector<int>& nums1, vector<int>& nums2) {
        vector<int> res;
        unordered_set<int> existed;
        for (int num : nums1) {
            existed.insert(num);
        }
        for (int num : nums2) {
            if (existed.count(num) > 0) {
                res.push_back(num);
                existed.erase(num);
            }
        }
        return res;
    }
};
```
