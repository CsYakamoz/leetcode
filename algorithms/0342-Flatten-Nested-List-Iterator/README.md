## [342. Flatten Nested List Iterator](https://leetcode.com/problems/flatten-nested-list-iterator/description/)

### Description

Given a nested list of integers, implement an iterator to flatten it.

Each element is either an integer, or a list -- whose elements may also be integers or other lists.

**Example 1:**
Given the list `[[1,1],2,[1,1]]`,

By calling _next_ repeatedly until _hasNext_ returns false, the order of elements returned by _next_ should be: `[1,1,2,1,1]`.

**Example 2:**
Given the list `[1,[4,[6]]]`,

By calling _next_ repeatedly until _hasNext_ returns false, the order of elements returned by _next_ should be: `[1,4,6]`.

**Difficulty:** `Medium`

**Tags:** `Stack` `Design`

### Solution One

```c++
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * class NestedInteger {
 *   public:
 *     // Return true if this NestedInteger holds a single integer, rather than a nested list.
 *     bool isInteger() const;
 *
 *     // Return the single integer that this NestedInteger holds, if it holds a single integer
 *     // The result is undefined if this NestedInteger holds a nested list
 *     int getInteger() const;
 *
 *     // Return the nested list that this NestedInteger holds, if it holds a nested list
 *     // The result is undefined if this NestedInteger holds a single integer
 *     const vector<NestedInteger> &getList() const;
 * };
 */
class NestedIterator {
public:
    NestedIterator(vector<NestedInteger> &nestedList) {
        dfs(nestedList);
    }

    int next() {
        int x = q.front();
        q.pop();
        return x;
    }

    bool hasNext() {
       return !q.empty();
    }

private:
    queue<int> q;

    void dfs(const vector<NestedInteger> &nestedList) {
        for (auto &i: nestedList) {
            if (i.isInteger()) {
                q.push(i.getInteger());
            } else {
                dfs(i.getList());
            }
        }
    }
};
```

### Solution Two - In Top Solutions

[Simple Java solution using a stack with explanation](https://discuss.leetcode.com/topic/42042/simple-java-solution-using-a-stack-with-explanation)

> A question before this is the Nested List Weight Sum, and it requires recursion to solve. As it carries to this problem that we will need recursion to solve it. But since we need to access each NestedInteger at a time, we will use a stack to help.
>
> In the constructor, we push all the nestedList into the stack from back to front, so when we pop the stack, it returns the very first element. Second, in the hasNext() function, we peek the first element in stack currently, and if it is an Integer, we will return true and pop the element. If it is a list, we will further flatten it. This is iterative version of flatting the nested list. Again, we need to iterate from the back to front of the list.

```java
public class NestedIterator implements Iterator<Integer> {
    Stack<NestedInteger> stack = new Stack<>();
    public NestedIterator(List<NestedInteger> nestedList) {
        for(int i = nestedList.size() - 1; i >= 0; i--) {
            stack.push(nestedList.get(i));
        }
    }

    @Override
    public Integer next() {
        return stack.pop().getInteger();
    }

    @Override
    public boolean hasNext() {
        while(!stack.isEmpty()) {
            NestedInteger curr = stack.peek();
            if(curr.isInteger()) {
                return true;
            }
            stack.pop();
            for(int i = curr.getList().size() - 1; i >= 0; i--) {
                stack.push(curr.getList().get(i));
            }
        }
        return false;
    }
}
```
