## [93. Restore IP Address](https://leetcode.com/problems/restore-ip-addresses/description/)

### Description

Given a string containing only digits, restore it by returning all possible valid IP address combinations.

For example:
Given `"25525511135"`,

return `["255.255.11.135", "255.255.111.35"]`. (Order does not matter)

**Difficulty:** `Medium`

**Tags:** `String` `Backtracking`

### Solution One

```c++
class Solution {
public:
    vector<string> restoreIpAddresses(string s) {
        if (s.size() < 4)
            return {};

        vector<string> res;
        string ip;
        dfs(s, 0, 0, ip, res);
        return res;
    }

private:
    void dfs(const string &s, int pos, size_t i, string &ip, vector<string> &res) {
        // End condition
        if (pos == 4) {
            if (i == s.size())
                res.push_back(ip);
            return;
        }

        // Get one digit
        ip.push_back(s[i]);
        if (pos != 3) ip.push_back('.');
        dfs(s, pos + 1, i + 1, ip, res);
        ip.erase(ip.end() - 1 - (pos != 3), ip.end());

        if (s[i] == '0')
            return;

        // Get two digit
        if (i + 2 <= s.size())
        {
            ip.append(s.substr(i, 2));
            if (pos != 3) ip.push_back('.');
            dfs(s, pos + 1, i + 2, ip, res);
            ip.erase(ip.end() - 2 - (pos != 3), ip.end());
        }

        // Get three digit
        if (i + 3 <= s.size())
        {
            int val = stoi(s.substr(i, 3));
            if (val < 256)
            {
                ip.append(s.substr(i, 3));
                if (pos != 3) ip.push_back('.');
                dfs(s, pos + 1, i + 3, ip, res);
                ip.erase(ip.end() - 3 - (pos != 3), ip.end());
            }
        }
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
private:
    void restoreIPAdresses(string addr, int cnt, string remaining, vector<string> &res) {
        int len;

        if (cnt == 3) {
            len = remaining.length();
            if (len > 0 && len <= 3 && remaining[0] != '0' && stoi(remaining) <= 255) {
                res.push_back(addr + "." + remaining);
            } else if (len == 1 && remaining[0] == '0') {
                res.push_back(addr + "." + remaining);
            } else {
                return;
            }
        }

        len = remaining.length();

        for (int i = 1; i <= min(3, len); i++) {
            string s = remaining.substr(0, i);

            if (stoi(s) > 255) break;

            restoreIPAdresses(addr == "" ? s : addr + "." + s, cnt + 1, remaining.substr(i), res);

            if (i == 1 && s[0] == '0') break;
        }
    }

public:
    vector<string> restoreIpAddresses(string s) {
        vector<string> res;
        restoreIPAdresses("", 0, s, res);

        return res;
    }
};
```

### Solution Three - In Top Solutions

```c++
class Solution {
public:
    vector<string> restoreIpAddresses(string s) {
        if(s.size() > 4 * 3 || s.size() < 4) return {};
        vector<string> res;
        func(0, 4, "", res, s);
        return res;
    }

    void func(int start, int part, string out, vector<string>& res, string s) {
        if (start == s.size() && part == 0) {
            out.pop_back();
            res.push_back(out);
        }
        if (part <= 0) return;

        for (int i = start, j = 1; i < s.size() && j <= 3; i++, j++) {
            //if (j > 3) return;
            string tmp = s.substr(start, j);
            if (isValid(tmp)) {
                func(i + 1, part - 1, out + tmp + '.', res, s);
            }
            //if (isMore(tmp)) return;
        }
    }

/*    bool isMore(string str) {
        if (str.size() > 3) return true;
        else return false;
    }*/
    bool isValid(string str) {
        if (str.empty() || str.size() > 3 || (str.size() > 1 && str[0] == '0')) return false;
        else {
            int res = 0;
            for (int i = 0; i < str.size(); i ++) {
                res = res * 10 + str[i] - '0';
            }
            return (res <= 255 && res >=0);
        }

    }
};
```
