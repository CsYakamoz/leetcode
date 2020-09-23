## [19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/#/description)

### Description

Given a linked list, remove the $n^{th}$ node from the end of list and return its head.

For example,

```
Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
```

**Note:**
Given *n* will always be valid.
Try to do this in one pass.

**Difficult:** `Medium`

**Tags:** `Linked List` `Two Pointers`

### Solution One

首先遍历一遍 List，得到 List 的长度 length

接着再次遍历 List，length 递减到等于 n 时，此时 current 指向要被删除的结点，frontNode 指向 current 的前一个结点

```c++
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        int length = 0;
        ListNode *current = head;
        ListNode *start = new ListNode(0);
        start->next = head;
        while (current != nullptr)
        {
            ++length;
            current = current->next;
        }
        ListNode *frontNode = start;
        current = start->next;
        while (length != n)
        {
            frontNode = frontNode->next;
            current = current->next;
            length--;
        }
        frontNode->next = current->next;
        delete current;
        return start->next;
    }
};
```

### Solution Two - In Top Solutions

`Two Pointers`

创建两个指针：fast、slow

fast 先移动，移动到与 slow 的距离为 n（ n 为有效值）

接着同时移动 fast、slow，当 fast 指向链尾时，slow 指向要被删除结点的前一结点

此时删除要被删除的结点，并返回 head 即可

for 循环 + while 循环总共遍历了 List 一次

```c++
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode *start = new ListNode(0);
        ListNode *slow = start, *fast = start;
        slow->next = head;
        for (int i = 0; i <= n; i++)
        {
            fast = fast->next;
        }
        while (fast != nullptr)
        {
            slow = slow->next;
            fast = fast->next;
        }
        ListNode *del = slow->next;
        slow->next = del->next;
        delete del;
        return start->next;
    }
};
```
