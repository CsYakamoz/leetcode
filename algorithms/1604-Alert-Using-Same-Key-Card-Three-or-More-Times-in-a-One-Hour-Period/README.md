## [1604. Alert Using Same Key-Card Three or More Times in a One Hour Period](https://leetcode.com/problems/alert-using-same-key-card-three-or-more-times-in-a-one-hour-period/)

### Description

LeetCode company workers use key-cards to unlock office doors. Each time a worker uses their key-card, the security system saves the worker's name and the time when it was used. The system emits an **alert** if any worker uses the key-card **three or more times** in a one-hour period.

You are given a list of strings `keyName` and `keyTime` where `[keyName[i], keyTime[i]]` corresponds to a person's name and the time when their key-card was used **in a** **single day**.

Access times are given in the **24-hour time format "HH:MM"**, such as `"23:51"` and `"09:49"`.

Return a _list of unique worker names who received an alert for frequent keycard use_. Sort the names in **ascending order alphabetically**.

Notice that `"10:00"` - `"11:00"` is considered to be within a one-hour period, while `"22:51"` - `"23:52"` is not considered to be within a one-hour period.

**Example 1:**

```
Input: keyName = ["daniel","daniel","daniel","luis","luis","luis","luis"], keyTime = ["10:00","10:40","11:00","09:00","11:00","13:00","15:00"]
Output: ["daniel"]
Explanation: "daniel" used the keycard 3 times in a one-hour period ("10:00","10:40", "11:00").
```

**Example 2:**

```
Input: keyName = ["alice","alice","alice","bob","bob","bob","bob"], keyTime = ["12:01","12:00","18:00","21:00","21:20","21:30","23:00"]
Output: ["bob"]
Explanation: "bob" used the keycard 3 times in a one-hour period ("21:00","21:20", "21:30").
```

**Example 3:**

```
Input: keyName = ["john","john","john"], keyTime = ["23:58","23:59","00:01"]
Output: []
```

**Example 4:**

```
Input: keyName = ["leslie","leslie","leslie","clare","clare","clare","clare"], keyTime = ["13:00","13:20","14:00","18:00","18:51","19:30","19:49"]
Output: ["clare","leslie"]
```

**Constraints:**

- `1 <= keyName.length, keyTime.length <= 10^5`
- `keyName.length == keyTime.length`
- `keyTime` are in the format **"HH:MM"**.
- `[keyName[i], keyTime[i]]` is **unique**.
- `1 <= keyName[i].length <= 10`
- `keyName[i] contains only lowercase English letters.`

**Difficult:** `Medium`

**Tags:** `String` `Ordered Map`

### Solution One

```javascript
/**
 * @param {string} str
 * @returns {number[]}
 */
const getHourAndMinute = (str) =>
  [str.slice(0, 2), str.slice(3, 5)].map((item) => parseInt(item));

/**
 * @param {string[]} timeList
 * @returns {boolean}
 */
const check = (timeList) => {
  if (timeList.length < 3) {
    return false;
  }

  let [begin, end] = [0, 1];
  while (end < timeList.length) {
    const [beginHH, beginMM] = getHourAndMinute(timeList[begin]);
    const [endHH, endMM] = getHourAndMinute(timeList[end]);

    if (beginHH === endHH || (endHH - beginHH === 1 && beginMM >= endMM)) {
      end++;
    } else {
      begin++;
    }

    if (end - begin >= 3) {
      return true;
    }
  }

  return end - begin >= 3;
};

/**
 * @param {string[]} keyName
 * @param {string[]} keyTime
 * @return {string[]}
 */
const alertNames = function (keyName, keyTime) {
  /** @type {Map<string, string[]} */
  const dict = new Map();
  for (let i = 0, len = keyName.length; i < len; i++) {
    if (!dict.has(keyName[i])) {
      dict.set(keyName[i], []);
    }

    dict.get(keyName[i]).push(keyTime[i]);
  }

  /** @type {string[]} */
  const result = [];
  for (const [key, value] of dict) {
    if (check(value.sort())) {
      result.push(key);
    }
  }

  return result.sort();
};
```
