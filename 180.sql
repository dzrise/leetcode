/*
 180. Consecutive Numbers
Medium
Topics
Companies
SQL Schema
Pandas Schema
Table: Logs

+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| num         | varchar |
+-------------+---------+
In SQL, id is the primary key for this table.
id is an autoincrement column.


Find all numbers that appear at least three times consecutively.

Return the result table in any order.

The result format is in the following example.



Example 1:

Input:
Logs table:
+----+-----+
| id | num |
+----+-----+
| 1  | 1   |
| 2  | 1   |
| 3  | 1   |
| 4  | 2   |
| 5  | 1   |
| 6  | 2   |
| 7  | 2   |
+----+-----+
Output:
+-----------------+
| ConsecutiveNums |
+-----------------+
| 1               |
+-----------------+
Explanation: 1 is the only number that appears consecutively for at least three times.
 */

SELECT DISTINCT lg1.num as ConsecutiveNums
FROM Logs lg1, Logs lg2, Logs lg3
WHERE
    lg2.id = lg1.id - 1
  AND lg3.id = lg2.id - 1
  AND lg1.num = lg2.num
  AND lg2.num = lg3.num