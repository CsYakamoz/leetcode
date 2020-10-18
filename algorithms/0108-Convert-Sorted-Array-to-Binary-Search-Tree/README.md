## [108. Convert Sorted Array to Binary Search Tree](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/)

### Description

Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

**Difficulty:** `Easy`

**Tags:** `Tree` `Depth-first Search`

### Solution One

```c++
class Solution {
public:
    TreeNode* sortedArrayToBST(vector<int>& nums) {
        return BST(nums, 0, nums.size());
    }

    TreeNode* BST(vector<int> &nums, int left, int right)
    {
        TreeNode *root = nullptr;
        if (left < right)
        {
            int mid = left + (right - left) / 2;
            root = new TreeNode(nums[mid]);
            root->left = BST(nums, left, mid);
            root->right = BST(nums, mid + 1, right);
        }
        return root;
    }
};
```
