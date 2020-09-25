## [82. Remove Duplicates from Sorted List II](https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description/)

### Description

Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only _distinct_ numbers from the original list.

For example,
Given `1->2->3->3->4->4->5`, return `1->2->5`.
Given `1->1->1->2->3`, return `2->3`.

**Difficult:** `Medium`

**Tags:** `Linked List`

### Solution One

利用二级指针

```c++
class Solution {
public:
    ListNode* deleteDuplicates(ListNode* head) {
        ListNode **curr = &head;
        while (*curr)
        {
            ListNode *entry = *curr;
            int val = entry->val;
            int length = 1;
            while (entry->next && entry->next->val == val)
            {
                length++;
                entry = entry->next;
            }
            if (length >= 2)
                *curr = entry->next;
            else
                curr = &entry->next;
        }
        return head;
    }
};
```
