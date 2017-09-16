## [386. Lexicographical Numbers](https://leetcode.com/problems/lexicographical-numbers/description/)

### Description

Given an integer *n*, return 1 - *n* in lexicographical order.

For example, given 13, return: [1,10,11,12,13,2,3,4,5,6,7,8,9].

Please optimize your algorithm to use less time and space. The input size may be as large as 5,000,000.



**Difficult:** `Medium`

**Tags:** 



### Solution One

```c++
class Solution {
public:
    vector<int> lexicalOrder(int n) {
        for (size_t i = 1; i <= n && i < 10; i++)
        {
            res.push_back(i);
            helper(i * 10, n);
        }
        return res;
    }

private:
    vector<int> res;
    
    void helper(int i, const int n)
    {
        if (i > n) return;

        for (size_t j = 0; i <= n && j < 10; i++, j++)
        {
            res.push_back(i);
            helper(i * 10, n);
        }
    }
};
```



### Solution Two - In Top Solutions

[Simple Java DFS Solution](https://discuss.leetcode.com/topic/55377/simple-java-dfs-solution)

> The idea is pretty simple. If we look at the order we can find out we just keep adding digit from 0 to 9 to every digit and make it a tree.
> Then we visit every node in pre-order. 
> ​       1        2        3    ...
> ​      /\        /\       /\
>    10 ...19  20...29  30...39   ....

```java
public class Solution {
    public List<Integer> lexicalOrder(int n) {
        List<Integer> res = new ArrayList<>();
        for(int i=1;i<10;++i){
          dfs(i, n, res); 
        }
        return res;
    }
    
    public void dfs(int cur, int n, List<Integer> res){
        if(cur>n)
            return;
        else{
            res.add(cur);
            for(int i=0;i<10;++i){
                if(10*cur+i>n)
                    return;
                dfs(10*cur+i, n, res);
            }
        }
    }
}
```



### Solution Three - In Top Solutions

```c++
class Solution {
public:
    vector<int> lexicalOrder(int n) {
        vector<int> st(n);
        int i;
        int temp=1;
        for(i=0;i<n;i++)
        {
            st[i]=temp;
            if(temp*10<=n)
                temp=temp*10;
            else 
            {
                if(temp>=n)
                    temp=temp/10;
                temp=temp+1;
                while(temp%10==0)
                    temp=temp/10;
            }
        }
        return st;
    }
};
```

