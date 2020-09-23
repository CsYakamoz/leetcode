## [182. Duplicate Emails](https://leetcode.com/problems/duplicate-emails/description/)

### Description

Write a SQL query to find all duplicate emails in a table named `Person`.

```
+----+---------+
| Id | Email   |
+----+---------+
| 1  | a@b.com |
| 2  | c@d.com |
| 3  | a@b.com |
+----+---------+

```

For example, your query should return the following for the above table:

```
+---------+
| Email   |
+---------+
| a@b.com |
+---------+

```

**Note**: All emails are in lowercase.



**Difficult:** `Easy`

**Tags:** 



### Solution One

```mysql
# Write your MySQL query statement below
SELECT DISTINCT a.Email
FROM Person AS a, Person AS b
WHERE a.id != b.id AND a.Email = b.Email;
```



### Solutions

[182. Duplicate Emails - Solution](https://leetcode.com/problems/duplicate-emails/solution/)

