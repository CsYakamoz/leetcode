## [109. Convert Sorted List to Binary Search Tree](https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/description/)

### Description

Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.

**Difficult:** `Medium`

**Tags:** `Depth-first Search` `Linked List`

### Solution One

```c++
class Solution {
public:
    TreeNode* sortedListToBST(ListNode* head) {
        if (head == nullptr)
            return nullptr;

        return BST(head);
    }

    TreeNode* BST(ListNode *head)
    {
        if (head == nullptr)
            return nullptr;

        if (head->next == nullptr)
            return new TreeNode(head->val);

        ListNode *slow = head;
        ListNode *fast = head;
        ListNode *prev = nullptr;
        while (fast && fast->next)
        {
            prev = slow;
            slow = slow->next;
            fast = fast->next->next;
        }

        TreeNode *root = new TreeNode(slow->val);
        prev->next = nullptr;
        root->left = BST(head);
        root->right = BST(slow->next);
        return root;
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:

    TreeNode* sortedListToBST(ListNode* head) {
        if (!head) return NULL;
        ListNode *prev=NULL,* mid = head, *f = head->next;
        while (f && f->next) {
            prev=mid;
            mid=mid->next;
            f=f->next->next;
        }
        if (prev) {
            prev->next=NULL;
        }
        TreeNode * root = new TreeNode(mid->val);
        if (mid != head) {
            root->left=sortedListToBST(head);
        }
        root->right=sortedListToBST(mid->next);
        return root;
    }
};
```

### Solution Three - In Top Solutions

[Share my code with O(n) time and O(1) space](https://discuss.leetcode.com/topic/3286/share-my-code-with-o-n-time-and-o-1-space)

> count is a function to calculate the size of list.
>
> Key words: inorder traversal.

```java
class Solution {
public:
    ListNode *list;
    int count(ListNode *node){
        int size = 0;
        while (node) {
            ++size;
            node = node->next;
        }
        return size;
    }

    TreeNode *generate(int n){
        if (n == 0)
            return NULL;
        TreeNode *node = new TreeNode(0);
        node->left = generate(n / 2);
        node->val = list->val;
        list = list->next;
        node->right = generate(n - n / 2 - 1);
        return node;
    }

    TreeNode *sortedListToBST(ListNode *head) {
        this->list = head;
        return generate(count(head));
    }
};
```
