## [1184. Distance Between Bus Stops](https://leetcode.com/problems/distance-between-bus-stops/)

### Description

A bus has `n` stops numbered from `0` to `n - 1` that form a circle. We know the distance between all pairs of neighboring stops where `distance[i]` is the distance between the stops number `i` and `(i + 1) % n`.

The bus goes along both directions i.e. clockwise and counterclockwise.

Return the shortest distance between the given `start` and `destination` stops.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2019/09/03/untitled-diagram-1.jpg)

```
Input: distance = [1,2,3,4], start = 0, destination = 1
Output: 1
Explanation: Distance between 0 and 1 is 1 or 9, minimum is 1.
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2019/09/03/untitled-diagram-1-1.jpg)

```
Input: distance = [1,2,3,4], start = 0, destination = 2
Output: 3
Explanation: Distance between 0 and 2 is 3 or 7, minimum is 3.
```

**Example 3:**

![img](https://assets.leetcode.com/uploads/2019/09/03/untitled-diagram-1-2.jpg)

```
Input: distance = [1,2,3,4], start = 0, destination = 3
Output: 4
Explanation: Distance between 0 and 3 is 6 or 4, minimum is 4.
```

**Constraints:**

- `1 <= n <= 10^4`
- `distance.length == n`
- `0 <= start, destination < n`
- `0 <= distance[i] <= 10^4`

**Difficult:** `Easy`

**Tags:** `Array`

### Solution One

```javascript
/**
 * @param {number} num
 * @param {number} n
 */
const dec = (num, n) => (num !== 0 ? num - 1 : n - 1);

/**
 * @param {number[]} distance
 * @param {number} start
 * @param {number} destination
 * @return {number}
 */
const distanceBetweenBusStops = function (distance, start, destination) {
  let clockwise = 0;
  let clockwiseStart = start;
  while (clockwiseStart !== destination) {
    clockwise += distance[clockwiseStart];

    clockwiseStart = (clockwiseStart + 1) % distance.length;
  }

  let counterclockwise = 0;
  let counterclockwiseStart = start;
  while (counterclockwiseStart !== destination) {
    counterclockwise += distance[dec(counterclockwiseStart, distance.length)];

    counterclockwiseStart = dec(counterclockwiseStart, distance.length);
  }

  return Math.min(clockwise, counterclockwise);
};
```
