## [690. Employee Importance](https://leetcode.com/problems/employee-importance/description/)

### Description

You are given a data structure of employee information, which includes the employee's **unique id**, his **importance value** and his **direct** subordinates' id.

For example, employee 1 is the leader of employee 2, and employee 2 is the leader of employee 3. They have importance value 15, 10 and 5, respectively. Then employee 1 has a data structure like [1, 15, [2]], and employee 2 has [2, 10, [3]], and employee 3 has [3, 5, []]. Note that although employee 3 is also a subordinate of employee 1, the relationship is **not direct**.

Now given the employee information of a company, and an employee id, you need to return the total importance value of this employee and all his subordinates.

**Example 1:**

```
Input: [[1, 5, [2, 3]], [2, 3, []], [3, 3, []]], 1
Output: 11
Explanation:
Employee 1 has importance value 5, and he has two direct subordinates: employee 2 and employee 3. They both have importance value 3. So the total importance value of employee 1 is 5 + 3 + 3 = 11.

```

**Note:**

1. One employee has at most one **direct** leader and may have several subordinates.
2. The maximum number of employees won't exceed 2000.



**Difficult:** `Easy`

**Tags:** `Hash Table` `Depth-first Search` `Breadth-first Search`



### Solution One - In Top Solutions

`DFS`

```c++
class Solution {
public:
    int getImportance(vector<Employee*> employees, int id) {
        map<int, Employee*> hash;
        for (auto iter : employees)
            hash[iter->id] = iter;

        return helper(hash[id], hash);
    }

private:
    int helper(Employee *leader, map<int, Employee*> &hash) {
        int sum = leader->importance;
        for (auto iter : leader->subordinates)
            sum += helper(hash[iter], hash);
        return sum;
    }
};
```



### Solution Two - In Top Solutions

`BFS`

```c++
class Solution {
public:
    int getImportance(vector<Employee*> employees, int id) {
        int sum = 0;
        unordered_map<int, Employee*> hash;
        queue<Employee*> q;
        
        for (auto iter : employees)
            hash[iter->id] = iter;
        q.push(hash[id]);

        while (!q.empty()) {
            auto e = q.front();
            sum += e->importance;
            for (auto iter : e->subordinates)
                q.push(hash[iter]);
            q.pop();
        }

        return sum;
    }
};
```



