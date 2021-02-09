## [1736. Latest Time by Replacing Hidden Digits](https://leetcode.com/problems/latest-time-by-replacing-hidden-digits/)

### Description

<p>You are given a string <code>time</code> in the form of <code> hh:mm</code>, where some of the digits in the string are hidden (represented by <code>?</code>).</p>

<p>The valid times are those inclusively between <code>00:00</code> and <code>23:59</code>.</p>

<p>Return <em>the latest valid time you can get from</em> <code>time</code><em> by replacing the hidden</em> <em>digits</em>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> time = &quot;2?:?0&quot;
<strong>Output:</strong> &quot;23:50&quot;
<strong>Explanation:</strong> The latest hour beginning with the digit &#39;2&#39; is 23 and the latest minute ending with the digit &#39;0&#39; is 50.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> time = &quot;0?:3?&quot;
<strong>Output:</strong> &quot;09:39&quot;
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> time = &quot;1?:22&quot;
<strong>Output:</strong> &quot;19:22&quot;
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>time</code> is in the format <code>hh:mm</code>.</li>
	<li>It is guaranteed that you can produce a valid time from the given string.</li>
</ul>

**Difficulty:** `Easy`

**Tags:** `String` `Greedy`

### Solution One

```javascript
const isQuestionMark = (c) => c === '?';

/**
 * @param {string} time
 * @return {string}
 */
const maximumTime = function (time) {
  const list = time.split('');

  if (isQuestionMark(list[0]) && isQuestionMark(list[1])) {
    list[0] = '2';
    list[1] = '3';
  } else if (isQuestionMark(list[0])) {
    list[0] = parseInt(list[1]) > 3 ? '1' : '2';
  } else if (isQuestionMark(list[1])) {
    list[1] = parseInt(list[0]) === 2 ? '3' : '9';
  }

  if (isQuestionMark(list[3])) {
    list[3] = '5';
  }

  if (isQuestionMark(list[4])) {
    list[4] = '9';
  }

  return list.join('');
};
```

### Solution Two - In Top Solutions

link: [C++ 4 rules](https://leetcode.com/problems/latest-time-by-replacing-hidden-digits/discuss/1032018/C%2B%2B-4-rules)

```c++
string maximumTime(string time) {
    time[0] = time[0] != '?' ? time[0] : (time[1] == '?' || time[1] <= '3') ? '2' : '1';
    time[1] = time[1] != '?' ? time[1] : time[0] == '2' ? '3' : '9';
    time[3] = time[3] != '?' ? time[3] : '5';
    time[4] = time[4] != '?' ? time[4] : '9';
    return time;
}
```
