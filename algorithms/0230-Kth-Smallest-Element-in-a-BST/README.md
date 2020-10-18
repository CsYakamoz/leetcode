## [230. Kth Smallest Element in a BST](https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/)

### Description

Given a binary search tree, write a function `kthSmallest` to find the **k**th smallest element in it.

**Note: **
You may assume k is always valid, 1 ? k ? BST's total elements.

**Follow up:**
What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently? How would you optimize the kth Smallest routine?

**Difficulty:** `Medium`

**Tags:** `Binary Search` `Tree`

### Solution One

相当于中序遍历循环实现

```c++
class Solution {
public:
    int kthSmallest(TreeNode* root, int k) {
        int th = 0;
        stack<TreeNode*> s;
        int res;
        while (!s.empty() || root != nullptr)
        {
            if (root != nullptr)
            {
                s.push(root);
                root = root->left;
            }
            else
            {
                root = s.top();
                s.pop();
                th++;
                if (th == k)
                {
                    res = root->val;
                    break;
                }
                root = root->right;
            }
        }
        return res;
    }
};
```

### Solution Two

思路同样是中序遍历，但按照个人所知，一般是循环比递归快，但是对于这道题而言，递归比循环快，不知道为什么。

```c++
class Solution {
public:
    int kthSmallest(TreeNode* root, int k) {
        int th = 0;
        helper(root, th, k);
        return res;
    }

private:
    int res = 0;
    bool helper(TreeNode *root, int &th, const int &k)
    {
        if (root == nullptr) return false;

        if (helper(root->left, th, k)) return true;
        th++;
        if (th == k)
        {
            res = root->val;
            return true;
        }
        if (helper(root->right, th, k)) return true;
        return false;
    }
};
```
