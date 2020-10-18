## [46. Permutations](https://leetcode.com/problems/permutations/#/description)

### Description

Given a collection of **distinct** numbers, return all possible permutations.

For example,
`[1,2,3]` have the following permutations:

```c++
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
```

**Difficulty:** `Medium`

**Tags:** `Backtracking`

### Solution One

```c++
class Solution {
public:
    vector<vector<int>> permute(vector<int>& nums) {
        if (nums.size() < 2)
        {
            return vector<vector<int>>(nums.size(), nums);
        }
        for (size_t i = 0; i < nums.size(); i++)
        {
            vector<bool> isExist(nums.size());
            vector<int> per;
            getPermute(per, i, isExist, nums);
        }
        return result;
    }

private:
    vector<vector<int>> result;

    void getPermute(vector<int> per, size_t &index, vector<bool> isExist, vector<int> &nums)
    {
        per.push_back(nums[index]);
        isExist[index] = true;
        if (per.size() == nums.size())
        {
            result.push_back(per);
            return;
        }
        for (size_t i = 0; i < isExist.size(); i++)
        {
            if (isExist[i] != true)
            {
                getPermute(per, i, isExist, nums);
            }
        }
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    void helper(vector<int>& nums, int end, vector<vector<int>>& ans) {
        if (end == 0) {
            ans.push_back(nums);
            return;
        }

        for (int i = 0; i < end; ++i) {
            swap(nums[i], nums[end - 1]);
            helper(nums, end - 1, ans);
            swap(nums[i], nums[end - 1]);
        }
    }

    vector<vector<int>> permute(vector<int>& nums) {
        vector<vector<int>> ans;
        helper(nums, nums.size(), ans);
        return ans;
    }
};
```

### Solution Three - In Top Solutions

[My elegant recursive C++ solution with inline explanation](https://discuss.leetcode.com/topic/5881/my-elegant-recursive-c-solution-with-inline-explanation)

```c++
class Solution {
public:
    vector<vector<int> > permute(vector<int> &num) {
	    vector<vector<int> > result;

	    permuteRecursive(num, 0, result);
	    return result;
    }

    // permute num[begin..end]
    // invariant: num[0..begin-1] have been fixed/permuted
	void permuteRecursive(vector<int> &num, int begin, vector<vector<int> > &result)	{
		if (begin >= num.size()) {
		    // one permutation instance
		    result.push_back(num);
		    return;
		}

		for (int i = begin; i < num.size(); i++) {
		    swap(num[begin], num[i]);
		    permuteRecursive(num, begin + 1, result);
		    // reset
		    swap(num[begin], num[i]);
		}
    }
};
```
