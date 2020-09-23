## [206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/#/description)

### Description

Reverse a singly linked list.

**Hint:**

A linked list can be reversed either iteratively or recursively. Could you implement both?



**Difficult:** `Easy`

**Tags:** `Linked List`



### Solution One

首先检查链表长度是否小于2，若是，则什么都不用做

否则，以下面例子说明思路

node1 -> node2 -> node3 -> node4，对于递归而言

当递归到 node4 时，直接返回 node3，返回到结点 node3 ，执行操做`node4 -> next = node3`

**Note:** 因为最后一个结点是反转后的第一个结点，所以需要一个变量保存，此处使用`res`

此时链表为

```
node1 -> node2 -> node3
                   ↓ ↑
                  node4
```

接着返回到结点 node2，执行操做`node3 -> next = node2`

此时链表为

```
node1 -> node2
          ↓ ↑
node4 -> node3
```

最后返回到结点 node1，执行操作`node2 -> next = node1,  node1 -> next = nullptr`

此时链表为

```
node4 -> node3 -> node2 -> node1 
```

```c++
class Solution {
public:
    ListNode* reverseList(ListNode* head) {
        if (head == nullptr || head->next == nullptr)
        {	// If the length of list is 0 or 1, nothing to do
            return head;
        }
        auto p = helper(head->next);
        p->next = head;
        head->next = nullptr;
        return res;
    }

private:
    ListNode *res = nullptr;

    ListNode* helper(ListNode *node)
    {
        if (node->next)
        {
            auto p = helper(node->next);
            p->next = node;
        }
        else
        {
            res = node;
        }
        return node;
    }
};
```



### Solution Two

思路同**One**，不过这里改为使用 Stack

```c++
class Solution {
public:
    ListNode* reverseList(ListNode* head) {
        if (head == nullptr || head->next == nullptr)
        {	// If the length of list is 0 or 1, nothing to do
            return head;
        }
        ListNode *res = nullptr;
        stack<ListNode*> s;
        while (head->next)
        {
            s.push(head);
            head = head->next;
        }
        res = head;		// head points to the last node
        while (!s.empty())
        {
            auto p = s.top();
            s.pop();
            head->next = p;
            head = p;
        }
        head->next = nullptr;
        return res;
    }
};
```



### Solution Three - In Top Solutions

```c++
class Solution {
public:
    ListNode* reverseList(ListNode* head) {
         ListNode* pre = NULL;
        while (head) {
            ListNode* next = head -> next;
            head -> next = pre;
            pre = head;
            head = next;
        } 
        return pre;
    }
};
```



### Other Solutions

[206. Reverse Linked List - Solution](https://leetcode.com/problems/reverse-linked-list/#/solution)

