## [203. Remove Linked List Elements](https://leetcode.com/problems/remove-linked-list-elements/#/description)

### Description

Remove all elements from a linked list of integers that have value **val**.

**Example**
**_Given_:** 1 --> 2 --> 6 --> 3 --> 4 --> 5 --> 6, **val** = 6
**_Return_:** 1 --> 2 --> 3 --> 4 --> 5

**Difficult:** `Easy`

**Tags:** `Linked List`

### Solution One

[LINUS：利用二级指针删除单向链表](http://coolshell.cn/articles/8990.html)

```c++
class Solution {
public:
    ListNode* removeElements(ListNode* head, int val) {
        ListNode **curr = &head;
        while (*curr)
        {
            ListNode *entry = *curr;
            if (entry->val == val)
            {
                *curr = entry->next;
                delete entry;
            }
            else
            {
                curr = &entry->next;
            }
        }
        return head;
    }
};
```
