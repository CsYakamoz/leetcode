## [234. Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/#/description)

### Description

Given a singly linked list, determine if it is a palindrome.

**Follow up:**
Could you do it in O(n) time and O(1) space?



**Difficult:** `Easy`

**Tags:** `Linked List` `Two Pointers`



### Solution One

将链表中每个结点的值添加到 vector 中，接着检测该 vector 是否 palindrome

```c++
class Solution {
public:
    bool isPalindrome(ListNode* head) {
        if (head == nullptr || head->next == nullptr)
        {	// If the length of list is 0 or 1
            return true;
        }
        vector<int> list_val;
        while (head)
        {
            list_val.push_back(head->val);
            head = head->next;
        }
        size_t i = 0;
        size_t j = list_val.size() - 1;
        while (i < j)
        {
            if (list_val[i++] != list_val[j--])
            {
                return false;
            }
        }
        return true;
    }
};
```



### Solution Two - In Top Solutions

[My easy understand C++ solution](https://discuss.leetcode.com/topic/27605/my-easy-understand-c-solution)

```c++
class Solution {
public:
    ListNode* temp;
    bool isPalindrome(ListNode* head) {
        temp = head;
        return check(head);
    }
    
    bool check(ListNode* p) {
        if (NULL == p) return true;
        bool isPal = check(p->next) & (temp->val == p->val);
        temp = temp->next;
        return isPal;
    }
};
```



###  Solution Three - In Top Solutions

```c++
class Solution {
public:
    bool isPalindrome(ListNode* head) {
        if (head == nullptr)
            return true;
        ListNode *p1 = head, *p2 = head;
        while (p2 != nullptr && p2->next != nullptr)
        {
            p1 = p1->next;
            p2 = p2->next->next;
            if (p1 == p2)
                return false;
        }
        
        ListNode *last = nullptr, *next;
        while (p1 != nullptr)
        {
            next = p1->next;
            p1->next = last;
            last = p1;
            p1 = next;
        }
        p1 = last;
        
        while (p1 != nullptr)
        {
            if (head->val != p1->val)
                return false;
            head = head->next;
            p1 = p1->next;
        }
        return true;
    }
};
```



