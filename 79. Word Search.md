## [79. Word Search](https://leetcode.com/problems/word-search/description/)

### Description

Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

For example,
Given **board** = 

```
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]
```

**word** = `"ABCCED"`, -> returns `true`,

**word** = `"SEE"`, -> returns `true`,

**word** = `"ABCB"`, -> returns `false`.



**Difficult:** `Medium`

**Tags:** `Array` `Backtracking`



### Solution One

```c++
class Solution
{
  public:
    bool exist(vector<vector<char> > &board, string word)
    {
        if (board.empty() || board[0].empty())
            return false;

        m = board.size();
        n = board[0].size();

        for (int i = 0; i < m; i++)
            for (int j = 0; j < n; j++)
            {
                if (find(i, j, 0, board, word))
                    return true;
            }

        return false;
    }

  private:
    int m, n;
    bool find(int i, int j, size_t len, vector<vector<char> > &board, const string &word)
    {        
        if (len == word.size())
            return true;

        if (i < 0 || j < 0 || i >= m || j >= n || board[i][j] == 0)
            return false;

        if (board[i][j] != word[len++])
        {
            return false;
        }

        char c = board[i][j];
        board[i][j] = 0;
        bool res = find(i - 1, j, len, board, word) || find(i + 1, j, len, board, word) ||
                    find(i, j - 1, len, board, word) || find(i, j + 1, len, board, word);
        board[i][j] = c;

        return res;
    }
};
```



