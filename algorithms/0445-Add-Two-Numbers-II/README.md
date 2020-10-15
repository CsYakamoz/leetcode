## [445. Add Two Numbers II](https://leetcode.com/problems/add-two-numbers-ii/#/description)

### Description

You are given two **non-empty** linked lists representing two non-negative integers. The most significant digit comes first and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

**Follow up:**
What if you cannot modify the input lists? In other words, reversing the lists is not allowed.

**Example:**

```
Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 8 -> 0 -> 7
```

**Difficult:** `Medium`

**Tags:** `Linked List`

### Solution One

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
        stack<ListNode*> s1;
        stack<ListNode*> s2;
        ListNode *cur1 = l1;
        ListNode *cur2 = l2;
        while (cur1)
        {
            s1.push(cur1);
            cur1 = cur1->next;
        }
        while (cur2)
        {
            s2.push(cur2);
            cur2 = cur2->next;
        }
        bool isCarried = false;
        ListNode *head = new ListNode(0);
        while (!s1.empty() || !s2.empty() || isCarried)
        {
            int val1 = 0;
            int val2 = 0;
            int result = 0;
            if (!s1.empty())
            {
                val1 = s1.top()->val;
                s1.pop();
            }
            if (!s2.empty())
            {
                val2 = s2.top()->val;
                s2.pop();
            }
            if (isCarried)
            {
                result++;
                isCarried = false;
            }
            result += val1 + val2;
            if (result > 9)
            {
                isCarried = true;
                result %= 10;
            }
            ListNode *tmp = new ListNode(result);
            tmp->next = head->next;
            head->next = tmp;
        }
        return head->next;
    }
};
```

### Solution Two - In Top Solutions

[9-liner C++ O(N1+N2) solution with stacks to store digits, no list modification](https://discuss.leetcode.com/topic/76885/9-liner-c-o-n1-n2-solution-with-stacks-to-store-digits-no-list-modification)

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
        stack<int> s1, s2;
        while (l1)
        {
            s1.push(l1->val);
            l1 = l1->next;
        }
        while (l2)
        {
            s2.push(l2->val);
            l2 = l2->next;
        }
        ListNode *result = new ListNode(0), *tmp = nullptr;
        for (int sum = 0; !s1.empty() || !s2.empty(); tmp = new ListNode(sum /= 10), tmp->next = result, result= tmp)
        {
            if (!s1.empty())
            {
                sum += s1.top();
                s1.pop();
            }
            if (!s2.empty())
            {
                sum += s2.top();
                s2.pop();
            }
            result->val = sum % 10;
        }
        return result->val ? result : result->next;
    }
};
```
