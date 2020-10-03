## [1046. Last Stone Weight](https://leetcode.com/problems/last-stone-weight/)

### Description

We have a collection of stones, each stone has a positive integer weight.

Each turn, we choose the two **heaviest** stones and smash them together. Suppose the stones have weights `x` and `y` with `x <= y`. The result of this smash is:

- If `x == y`, both stones are totally destroyed;
- If `x != y`, the stone of weight `x` is totally destroyed, and the stone of weight `y` has new weight `y-x`.

At the end, there is at most 1 stone left. Return the weight of this stone (or 0 if there are no stones left.)

**Example 1:**

```
Input: [2,7,4,1,8,1]
Output: 1
Explanation:
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of last stone.
```

**Note:**

1. `1 <= stones.length <= 30`
2. `1 <= stones[i] <= 1000`

**Difficult:** `Easy`

**Tags:** `Heap` `Greedy`

### Solution One

1. 以 `stones` 为基础数据，建立最大堆
2. 当最大堆数量大于 1 时，取最大的两个，进行比较
   - 相等，则石头完全毁灭掉；
   - 不相等，则取差值，并塞入最大堆；
   - 重复第二步；

```javascript
class PriorityQueue {
  constructor(rawData, compare) {
    this._data = [null].concat(Array.from(rawData));
    this._compare = compare;

    for (let i = Math.floor(this.length / 2); i > 0; i--) {
      this._sink(i);
    }
  }

  get length() {
    return this._data.length - 1;
  }

  insert(elem) {
    this._data.push(elem);
    this._swim(this.length);
  }

  pop() {
    if (this.length === 0) {
      throw new Error('empty queue');
    }

    const val = this._data[1];

    this._exch(1, this.length);
    this._data.pop();
    this._sink(1);

    return val;
  }

  head() {
    if (this.length === 0) {
      throw new Error('empty queue');
    }

    return this._data[1];
  }

  empty() {
    return this.length === 0;
  }

  _swim(i) {
    while (
      i > 1 &&
      this._compare(this._data[i], this._data[Math.floor(i / 2)]) < 0
    ) {
      this._exch(i, Math.floor(i / 2));
      i = Math.floor(i / 2);
    }
  }

  _sink(i) {
    while (2 * i <= this.length) {
      let j = 2 * i;
      if (
        j < this.length &&
        this._compare(this._data[j], this._data[j + 1]) > 0
      ) {
        j++;
      }

      if (this._compare(this._data[i], this._data[j]) > 0) {
        this._exch(i, j);
        i = j;
      } else {
        break;
      }
    }
  }

  _exch(i, j) {
    const tmp = this._data[j];
    this._data[j] = this._data[i];
    this._data[i] = tmp;
  }
}

/**
 * @param {number[]} stones
 * @return {number}
 */
var lastStoneWeight = function (stones) {
  if (stones.length < 2) {
    return stones.pop() || 0;
  }

  const pq = new PriorityQueue(stones, (a, b) => b - a);

  while (pq.length > 1) {
    const i = pq.pop();
    const j = pq.pop();

    if (i === j) {
      continue;
    }

    pq.insert(Math.abs(i - j));
  }

  return pq.empty() ? 0 : pq.pop();
};
```
