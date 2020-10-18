## [232. Implement Queue using Stacks](https://leetcode.com/problems/implement-queue-using-stacks/description/)

### Description

Implement the following operations of a queue using stacks.

- push(x) -- Push element x to the back of queue.
- pop() -- Removes the element from in front of queue.
- peek() -- Get the front element.
- empty() -- Return whether the queue is empty.

Notes:

- You must use _only_ standard operations of a stack -- which means only `push to top`, `peek/pop from top`, `size`, and `is empty` operations are valid.
- Depending on your language, stack may not be supported natively. You may simulate a stack by using a list or deque (double-ended queue), as long as you use only standard operations of a stack.
- You may assume that all operations are valid (for example, no pop or peek operations will be called on an empty queue).

**Difficulty:** `Easy`

**Tags:** `Stack` `Design`

### Solution One - In Top Solutions

[232. Implement Queue using Stacks - Solution](https://leetcode.com/problems/implement-queue-using-stacks/solution/)

```c++
class MyQueue {
public:
    /** Initialize your data structure here. */
    MyQueue(): front(0) {}

    /** Push element x to the back of queue. */
    void push(int x) {
        if (s1.empty())
            front = x;

        s1.push(x);
    }

    /** Removes the element from in front of queue and returns that element. */
    int pop() {
        if (s2.empty())
            while (!s1.empty())
            {
                s2.push(s1.top());
                s1.pop();
            }

        int val = s2.top();
        s2.pop();
        return val;
    }

    /** Get the front element. */
    int peek() {
        if (!s2.empty())
            return s2.top();

        return front;
    }

    /** Returns whether the queue is empty. */
    bool empty() {
        return s1.empty() && s2.empty();
    }

private:
    int front;
    stack<int> s1;
    stack<int> s2;
};
```
