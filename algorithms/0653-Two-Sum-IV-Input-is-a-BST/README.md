## [653. Two Sum IV - Input is a BST](https://leetcode.com/problems/two-sum-iv-input-is-a-bst/description/)

### Description

Given a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their sum is equal to the given target.

**Example 1:**

```
Input:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 9

Output: True

```

**Example 2:**

```
Input:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 28

Output: False
```

**Difficulty:** `Easy`

**Tags:** `Tree`

### Solution One

```c++
class Solution {
public:
    bool findTarget(TreeNode* root, int k) {
        TreeNode *node = nullptr;
        queue<TreeNode*> q;

        if (root)
            q.push(root);

        while (!q.empty())
        {
            node = q.front();
            q.pop();
            if (search(root, k - node->val, node))
                return true;
            if (node->left)
                q.push(node->left);
            if (node->right)
                q.push(node->right);
        }

        return false;
    }

    bool search(TreeNode *root, int val, TreeNode *&node)
    {
        while (root != nullptr)
        {
            if (val == root->val)
                return node != root;
            else if (val < root->val)
                root = root->left;
            else
                root = root->right;
        }
        return false;
    }
};
```

### Solution Two - In Top Solutions

[[C++\] Clean Code - O(n) time O(lg n) space - BinaryTree Iterator](https://discuss.leetcode.com/topic/98391/c-clean-code-o-n-time-o-lg-n-space-binarytree-iterator)

> **O(n) time O(n) space - Inorder Vector**

```c++
class Solution {
public:
    bool findTarget(TreeNode* root, int k) {
        vector<int> nums;
        inorder(root, nums);
        return findTargetInSortedArray(nums, k);
    }

private:
    void inorder(TreeNode* node, vector<int>& nums) {
        if (!node) return;
        inorder(node->left, nums);
        nums.push_back(node->val);
        inorder(node->right, nums);
    }

    bool findTargetInSortedArray(vector<int> a, int target) {
        for (int i = 0, j = a.size() - 1; i < j;) {
            int sum = a[i] + a[j];
            if (sum == target) {
                return true;
            }
            else if (sum < target) {
                i++;
            }
            else {
                j--;
            }
        }

        return false;
    }
};
```

> **O(n) time O(lg n) space - BinaryTree Iterator**

```c++
class BSTIterator {
    stack<TreeNode*> s;
    TreeNode* node;
    bool forward;
public:
    BSTIterator(TreeNode *root, bool forward) : node(root), forward(forward) {};
    bool hasNext() {
        return node != nullptr || !s.empty();
    }
    int next() {
        while (node || !s.empty()) {
            if (node) {
                s.push(node);
                node = forward ? node->left : node->right;
            }
            else {
                node = s.top();
                s.pop();
                int nextVal = node->val;
                node = forward ? node->right : node->left;
                return nextVal;
            }
        }

        return -1;  // never reach & not necessary
    }
};
class Solution {
public:
    bool findTarget(TreeNode* root, int k) {
        if (!root) return false;
        BSTIterator l(root, true);
        BSTIterator r(root, false);
        for (int i = l.next(), j = r.next(); i < j;) {
            int sum = i + j;
            if (sum == k) {
                return true;
            }
            else if (sum < k) {
                i = l.next();
            }
            else {
                j = r.next();
            }
        }
        return false;
    }
};
```

> Simplified using operator overloading

```c++
class BSTIterator {
    stack<TreeNode*> s;
    TreeNode* node;
    bool forward;
    int cur;
public:
    BSTIterator(TreeNode *root, bool forward) : node(root), forward(forward) { (*this)++; };
    int operator*() { return cur; }
    void operator++(int) {
        while (node || !s.empty()) {
            if (node) {
                s.push(node);
                node = forward ? node->left : node->right;
            }
            else {
                node = s.top();
                s.pop();
                cur = node->val;
                node = forward ? node->right : node->left;
                break;
            }
        }
    }
};
class Solution {
public:
    bool findTarget(TreeNode* root, int k) {
        if (!root) return false;
        BSTIterator l(root, true);
        BSTIterator r(root, false);
        while (*l < *r) {
            if (*l + *r == k) return true;
            else if (*l + *r < k) l++;
            else r++;
        }
        return false;
    }
};
```
