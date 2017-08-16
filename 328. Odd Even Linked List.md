## [328. Odd Even Linked List](https://leetcode.com/problems/odd-even-linked-list/description/)

### Description

Given a singly linked list, group all odd nodes together followed by the even nodes. Please note here we are talking about the node number and not the value in the nodes.

You should try to do it in place. The program should run in O(1) space complexity and O(nodes) time complexity.

**Example:**
Given `1->2->3->4->5->NULL`,
return `1->3->5->2->4->NULL`.

**Note:**
The relative order inside both the even and odd groups should remain as it was in the input. 
The first node is considered odd, the second node even and so on ...



**Difficult:** `Medium`

**Tags:** `Linked List`



### Solution One

24 ~ 26 行可以改为`cur = next != nullptr ? next->next : next`

```c++
class Solution {
public:
    ListNode* oddEvenList(ListNode* head) {
        if (!head || !head->next)
        {
            return head;
        }
        ListNode *oddTail = head;
        ListNode *evenTail = oddTail->next;
        ListNode *evenHead = evenTail;
        ListNode *cur = evenTail->next;
        while (cur)
        {
            ListNode *next = cur->next;
            // change the order
            oddTail->next = cur;
            cur->next = evenHead;
            evenTail->next = next;
            // oddTail points to the tail of the odd groups
            // evenTail points to the tail of the even groups
            oddTail = oddTail->next;
            evenTail = evenTail->next;

            cur = next;
            if (cur)
                cur = cur->next;
        }
        return head;
    }
};
```



### Solution Two - In Top Solutions

在 **One** 中，每次修改顺序，都会将 odd 组的尾部连接到 even 组的头部

这里则不，每次都只将 odd node 添加到 odd 组的尾部，even node 同理

最后才将 odd 组的尾部连接到 even 组的头部

```c++
class Solution {
public:
    ListNode* oddEvenList(ListNode* head) {
        if (!head || !head->next || !head->next->next)
        {
            return head;
        }
        ListNode *h1 = head;
        ListNode *p = head;
        ListNode *h2 = head->next;
        ListNode *q = h2;
        while (p->next && p->next->next)
        {
            p->next = q->next;
            q->next = p->next->next;
            p = p->next;
            q = q->next;
        }
        p->next = h2;
        return h1;
    }
};
```



