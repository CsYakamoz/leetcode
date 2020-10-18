## [36. Valid Sudoku](https://leetcode.com/problems/valid-sudoku/#/description)

### Description

Determine if a Sudoku is valid, according to: [Sudoku Puzzles - The Rules](http://sudoku.com.au/TheRules.aspx).

The Sudoku board could be partially filled, where empty cells are filled with the character `'.'`.

![img](http://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png)

A partially filled sudoku which is valid.

**Note:**
A valid Sudoku board (partially filled) is not necessarily solvable. Only the filled cells need to be validated.

**Difficulty:** `Medium`

**Tags:** `Hash Table`

### Solution One

第一个循环检查各行是否满足条件

第二个循环检车各列是否满足条件

第三个循环检查各个 3 x 3 矩阵是否满足条件

```c++
class Solution {
public:
    bool isValidSudoku(vector<vector<char>>& board) {
        for (size_t i = 0; i < 9; i++)
        {
            if (!checkRow(board[i],i))
            {
                return false;
            }
        }
        for (size_t j = 0; j < 9; j++)
        {
            if (!checkCol(board, j))
            {
                return false;
            }
        }
        for (size_t i = 0; i < 9; i = i + 3)
        {
            for (size_t j = 0; j < 9; j = j + 3)
            {
                if (!checkSquare(board, i, j))
                {
                    return false;
                }
            }
        }
        return true;
    }
private:
    bool checkRow(vector<char> &row, const int &r) {
        set<char> s;
        for (size_t i = 0; i < 9; i++)
        {
            if (row[i] != '.')
            {
                if (s.find(row[i]) == s.end())
                {
                    s.insert(row[i]);
                }
                else
                {
                    return false;
                }
            }
        }
        return true;
    }
    bool checkCol(vector<vector<char>> &board,const int &c) {
        set<char> s;
        for (size_t i = 0; i < 9; i++)
        {
            if (board[i][c] != '.')
            {
                if (s.find(board[i][c]) == s.end())
                {
                    s.insert(board[i][c]);
                }
                else
                {
                    return false;
                }
            }
        }
        return true;
    }
    bool checkSquare(vector<vector<char>> &board, const int &r, const int &c) {
        set<char> s;
        for (size_t i = r; i < r + 3; i++)
        {
            for (size_t j = c; j < c + 3; j++)
            {
                if (board[i][j] != '.')
                {
                    if (s.find(board[i][j]) == s.end())
                    {
                        s.insert(board[i][j]);
                    }
                    else
                    {
                        return false;
                    }
                }
            }
        }
        return true;
    }
};
```

### Solution Two

`Hash Table`

[判断数独是否合法](http://blog.csdn.net/witnessai1/article/details/49205847)

```c++
class Solution {
public:
    bool isValidSudoku(vector<vector<char>>& board) {
        vector<int> checkRow(10);
        vector<int> checkCol(10);
        vector<int> checkSubBoard(10);
        for (size_t i = 0; i < 9; i++)
        {
            fill(checkRow.begin(), checkRow.end(), 0);
            fill(checkCol.begin(), checkCol.end(), 0);
            fill(checkSubBoard.begin(), checkSubBoard.end(), 0);
            for (size_t j = 0; j < 9; j++)
            {
                if (!isValid(checkRow, board[i][j] - '0')
                    || !isValid(checkCol, board[j][i] - '0')
                    || !isValid(checkSubBoard, board[3 * (i/3) + j / 3][3 * (i%3) + j % 3] - '0'))
                {
                    return false;
                }
            }
        }
        return true;
    }
private:
    bool isValid(vector<int> &v, const int &val)
    {
        if (val < 0)
        {
            return true;
        }
        if (v[val] == 1)
        {
            return false;
        }
        v[val] = 1;
        return true;
    }
};
```
