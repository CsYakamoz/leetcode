## [2. Add Two Numbers](https://leetcode.com/problems/add-two-numbers/#/description)

### Description

You are given two linked lists representing two non-negative numbers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

**Input:** (2 -> 4 -> 3) + (5 -> 6 -> 4)
**Output:** 7 -> 0 -> 8

**Difficulty:** `Medium`

**Tags:** `Linked List` `Math`

### Solution One

考虑三种情况，即可

1. 两个链表长度相等，并且最后一次相加时值没有大于 9
2. 两个链表长度相等，并且最后一次相加时值大于 9
3. 两个链表长度不相等

```c++
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode *head = new ListNode(0);
        ListNode *tmp = head;
        bool carry = false;
        while( l1 != NULL || l2 != NULL || carry) {
            int result = 0;
            int x1 = 0;
            int x2 = 0;
            if (l1 != NULL)
            {
                x1 = l1->val;
                l1 = l1->next;
            }
            if (l2 != NULL)
            {
                x2 = l2->val;
                l2 = l2->next;
            }
            if (carry)
            {
                ++result;
                carry = false;
            }
            result = result + x1 + x2;
            if(result >= 10) {
                carry = true;
                result = result % 10;
            }
            tmp->next = new ListNode(result);
            tmp = tmp->next;
        }
        return head->next;
    }
};
```

### Solution Two - In Top Solutions

思路同上，不过这里使用 `sum` 代替上面的 `carry`

```c++
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode *head = new ListNode(0);
        ListNode *tmp = head;
        int sum = 0;
        while (l1 || l2 || sum)
        {
            if (l1)
            {
                sum += l1->val;
                l1 = l1->next;
            }
            if (l2)
            {
                sum += l2->val;
                l2 = l2->next;
            }
            tmp->next = new ListNode(sum % 10);
            tmp = tmp->next;
            sum /= 10;
        }
        return head->next;
    }
};
```
