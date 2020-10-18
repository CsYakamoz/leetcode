## [212. Word Search II](https://leetcode.com/problems/word-search-ii/description/)

### Description

Given a 2D board and a list of words from the dictionary, find all words in the board.

Each word must be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

For example,
Given **words** = `["oath","pea","eat","rain"]` and **board** =

```
[
  ['o','a','a','n'],
  ['e','t','a','e'],
  ['i','h','k','r'],
  ['i','f','l','v']
]
```

Return `["eat","oath"]`.

**Note:**
You may assume that all inputs are consist of lowercase letters `a-z`.

You would need to optimize your backtracking to pass the larger test. Could you stop backtracking earlier?

If the current candidate does not exist in all words' prefix, you could stop backtracking immediately. What kind of data structure could answer such query efficiently? Does a hash table work? Why or why not? How about a Trie? If you would like to learn how to implement a basic trie, please work on this problem: [Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/) first.

**Difficulty:** `Medium`

**Tags:** `Backtracking` `Trie`

### Solution One

```c++
class Solution
{
private:
    struct TrieNode
    {
        bool isEndOfWord;
        vector<shared_ptr<TrieNode>> next;
        TrieNode() : isEndOfWord(false)
        {
            for (size_t i = 0; i < 26; i++)
                next.push_back(nullptr);
        }
    };

    void insert(const string &word)
    {
        shared_ptr<TrieNode> cur(root);

        for (auto c : word)
        {
            size_t idx = c - 'a';
            if (cur->next[idx] == nullptr)
                cur->next[idx] = make_shared<TrieNode>();
            cur = cur->next[idx];
        }

        cur->isEndOfWord = true;
    }

    void search(int i, int j, shared_ptr<TrieNode> root,
        string &str, vector<vector<char>> &board)
    {
        if (i < 0 || j < 0 || i >= m || j >= n || board[i][j] == 0)
            return;

        shared_ptr<TrieNode> cur(root->next[board[i][j] - 'a']);
        if (cur == nullptr)
            return;

        str.push_back(board[i][j]);
        board[i][j] = 0;

        if (cur->isEndOfWord)
            res.insert(str);

        search(i - 1, j, cur, str, board);  // Up
        search(i + 1, j, cur, str, board);  // Down
        search(i, j - 1, cur, str, board);  // Left
        search(i, j + 1, cur, str, board);  // Right

        board[i][j] = str.back();
        str.pop_back();
    }

    shared_ptr<TrieNode> root = make_shared<TrieNode>();
    int m, n;
    unordered_set<string> res;

public:
    vector<string> findWords(vector<vector<char>> &board, vector<string> &words)
    {
        res.clear();

        for (auto &word : words)
            insert(word);

        string str;
        m = board.size();
        n = board[0].size();

        for (int i = 0; i < m; i++)
            for (int j = 0; j < n; j++)
            {
                search(i, j, root, str, board);
            }

        return vector<string>(res.begin(), res.end());
    }
};
```

### Solution Two - In Top Solutions

```c++
class TrieNode
{
public:
	TrieNode()
	{
		isWord = false;
		val = 0;
		for(int i = 0; i < 26; i++)
		{
			children[i] = NULL;
		}
	}

	int val;
	bool isWord;
	TrieNode* children[26];
};

class Solution {
public:
	TrieNode* buildTrie(vector<string>& words)
	{
		TrieNode* root = new TrieNode();
		for(int i = 0; i < words.size(); i++)
		{
			TrieNode *p = root;
			for(int j = 0; j < words[i].size(); j++)
			{
				if(p->children[words[i][j] - 'a'] == NULL) p->children[words[i][j] - 'a'] = new TrieNode();
				p = p->children[words[i][j] - 'a'];
			}
			p->isWord = true;
			p->val = i;
		}

		return root;
	}

	void dfs(vector<vector<char>>& board, int i, int j, int row, int col, TrieNode* root, vector<string>& words, vector<string>& res)
	{
		if(board[i][j] == 'X') return;
		if(root->children[board[i][j] - 'a'] == NULL) return;

		char t = board[i][j];
		if(root->children[t - 'a']->isWord)
		{
			res.push_back(words[root->children[t - 'a']->val]);
			root->children[t - 'a']->isWord = false;
		}
		board[i][j] = 'X';
		if(i > 0) dfs(board, i - 1, j, row, col, root->children[t - 'a'], words, res);
		if((i + 1) < row) dfs(board, i + 1, j, row, col, root->children[t - 'a'], words, res);
		if(j > 0) dfs(board, i, j - 1, row, col, root->children[t - 'a'], words, res);
		if((j + 1) < col) dfs(board, i, j + 1, row, col, root->children[t - 'a'], words, res);
		board[i][j] = t;
	}

	vector<string> findWords(vector<vector<char>>& board, vector<string>& words) {
		vector<string> res;
		int row = board.size();
		int col = board[0].size();
		if(row == 0) return res;
		if(col == 0) return res;
		int wordLen = words.size();
		if(wordLen == 0) return res;
		TrieNode *root = buildTrie(words);
		for(int i = 0; i < row; i++)
		{
			for(int j = 0; j < col; j++)
			{
				dfs(board, i, j, row, col, root, words, res);
			}
		}

		return res;
	}
};
```
