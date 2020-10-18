## [143. Recorder List](https://leetcode.com/problems/reorder-list/description/)

### Description

Given a singly linked list _L_: *L*0?*L*1?…?_L\*\*n_-1?*L*n,
reorder it to: *L*0?_L\*\*n_?*L*1?_L\*\*n_-1?*L*2?_L\*\*n_-2?…

You must do this in-place without altering the nodes' values.

For example,
Given `{1,2,3,4}`, reorder it to `{1,4,2,3}`.

**Difficulty:** `Medium`

**Tags:** `Linked List`

### Solution One

首先找到 链表的中点（do ... while 循环），接着将中点右侧（含中点）的链表反转

接着以`Two Pointers`的思路将左侧链表和反转后的右侧链表合并

```c++
class Solution {
public:
    void reorderList(ListNode* head) {
        if (!head || !head->next || !head->next->next)
        {
            return;
        }
        ListNode *prev = nullptr;
        ListNode *slow = head;
        ListNode *fast = head;
        do
        {
            prev = slow;
            slow = slow->next;
            fast = fast->next;
            if (fast)
                fast = fast->next;
        } while (fast != nullptr);
        prev->next = nullptr;
        ListNode *p1 = head;
        ListNode *p2 = reverseList(slow);
        while (p2)
        {
            ListNode *node1 = p1->next;
            ListNode *node2 = p2->next;
            p1->next = p2;
            p2->next = node1;
            p1 = node1;
            p2 = node2;
        }
    }

    ListNode* reverseList(ListNode *head)
    {
        if (head == nullptr || head->next == nullptr)
        {
            return head;
        }
        ListNode *pre = nullptr;
        while (head)
        {
            ListNode *node = head->next;
            head->next = pre;
            pre = head;
            head = node;
        }
        return pre;
    }
};
```
