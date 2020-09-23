## [141. Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/description/)

### Description

Given a linked list, determine if it has a cycle in it.

Follow up:
Can you solve it without using extra space?

**Difficult:** `Easy`

**Tags:** `Linked List` `Two Pointers`

### Solution One

[**Floyd 判圈算法**](https://zh.wikipedia.org/wiki/Floyd%E5%88%A4%E5%9C%88%E7%AE%97%E6%B3%95#.E7.AE.97.E6.B3.95.E6.8F.8F.E8.BF.B0)

```c++
class Solution {
public:
    bool hasCycle(ListNode *head) {
        if (!head)
        {
            return false;
        }
        ListNode *slow = head;
        ListNode *fast = head;
        do
        {
            slow = slow->next;
            fast = fast->next;
            if (fast)
            {
                fast = fast->next;
            }
        } while (fast != nullptr && slow != fast);
        if (fast == nullptr)
        {
            return false;
        }
        else
        {
            return true;
        }
    }
};
```
