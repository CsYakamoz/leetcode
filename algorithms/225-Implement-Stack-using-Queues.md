## [225. Implement Stack using Queues](https://leetcode.com/problems/implement-stack-using-queues/description/)

### Description

Implement the following operations of a stack using queues.

- push(x) -- Push element x onto stack.
- pop() -- Removes the element on top of the stack.
- top() -- Get the top element.
- empty() -- Return whether the stack is empty.

Notes:

- You must use _only_ standard operations of a queue -- which means only `push to back`, `peek/pop from front`, `size`, and `is empty` operations are valid.
- Depending on your language, queue may not be supported natively. You may simulate a queue by using a list or deque (double-ended queue), as long as you use only standard operations of a queue.
- You may assume that all operations are valid (for example, no pop or top operations will be called on an empty stack).

**Difficult:** `Easy`

**Tags:** `Stack` `Design`

### Solution One

```c++
class MyStack {
public:
    /** Initialize your data structure here. */
    MyStack() {}

    /** Push element x onto stack. */
    void push(int x) {
        q.push(x);
        move();
    }

    /** Removes the element on top of the stack and returns that element. */
    int pop() {
        int val = q.front();
        q.pop();
        return val;
    }

    /** Get the top element. */
    int top() {
        return q.front();
    }

    /** Returns whether the stack is empty. */
    bool empty() {
        return q.empty();
    }

private:
    queue<int> q;

    void move()
    {
        size_t length = q.size() - 1;
        for (size_t i = 0; i < length; i++)
        {
            int val = q.front();
            q.pop();
            q.push(val);
        }
    }
};
```

### Solutions

[225. Implement Stack using Queues - Solution](https://leetcode.com/problems/implement-stack-using-queues/solution/)
