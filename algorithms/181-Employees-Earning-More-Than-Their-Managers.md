## [181. Employees Earning More Than Their Managers](https://leetcode.com/problems/employees-earning-more-than-their-managers/description/)

### Description

The `Employee` table holds all employees including their managers. Every employee has an Id, and there is also a column for the manager Id.

```
+----+-------+--------+-----------+
| Id | Name  | Salary | ManagerId |
+----+-------+--------+-----------+
| 1  | Joe   | 70000  | 3         |
| 2  | Henry | 80000  | 4         |
| 3  | Sam   | 60000  | NULL      |
| 4  | Max   | 90000  | NULL      |
+----+-------+--------+-----------+

```

Given the `Employee` table, write a SQL query that finds out employees who earn more than their managers. For the above table, Joe is the only employee who earns more than his manager.

```
+----------+
| Employee |
+----------+
| Joe      |
+----------+
```

**Difficult:** `Easy`

**Tags:**

### Solution One - In Top Solutions

[A straightforward method](https://discuss.leetcode.com/topic/7420/a-straightforward-method)

```mysql
# Write your MySQL query statement below
select E1.Name AS Employee
from Employee as E1, Employee as E2
where E1.ManagerId = E2.Id and E1.Salary > E2.Salary
```

### Solution Two - In Top Solutions

[Two Straightforward way, using 'where' and 'join'](https://discuss.leetcode.com/topic/38081/two-straightforward-way-using-where-and-join)
