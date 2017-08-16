## [142. Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/description/)

### Description

Given a linked list, return the node where the cycle begins. If there is no cycle, return `null`.

**Note:** Do not modify the linked list.

**Follow up**:
Can you solve it without using extra space?



**Difficult:** `Medium`

**Tags:** `Linked List` `Two Pointers`



### Solution One

[Floyd 判圈算法](https://zh.wikipedia.org/wiki/Floyd%E5%88%A4%E5%9C%88%E7%AE%97%E6%B3%95#.E7.AE.97.E6.B3.95)

```c++
class Solution {
public:
    ListNode *detectCycle(ListNode *head) {
        if (head == nullptr)
        {
            return head;
        }
        ListNode *slow = head;
        ListNode *fast = head;
        do
        {
            slow = slow->next;
            fast = fast->next;
            if (fast) fast = fast->next;
        } while (fast != nullptr && fast != slow);
        if (fast == nullptr)
        {   // There is no cycle
            return nullptr;
        }
        slow = head;
        while (slow != fast)
        {
            slow = slow->next;
            fast = fast->next;
        }
        return slow;
    }
};
```



