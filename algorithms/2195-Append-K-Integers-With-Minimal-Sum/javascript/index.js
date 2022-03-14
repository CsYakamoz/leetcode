/**
 * left-closed, right-open, e.g., [begin, end)
 */
class Interval {
    /**
     * @param {number} begin
     * @param {number} end
     */
    constructor(begin, end) {
        this.begin = begin;
        this.end = end;
    }
}

/**
 * find the first element that its end is bigger than the target
 *
 * @param {Interval[]} intervalList
 * @param {number} target
 * @returns {number}
 */
const binarySearch = function (intervalList, target) {
    let left = 0;
    let right = intervalList.length;

    while (left < right) {
        const mid = parseInt((right - left) / 2) + left;
        if (intervalList[mid].end > target) {
            right = mid;
        } else {
            left = mid + 1;
        }
    }

    return right;
};

/**
 * @param {number[]} nums
 * @returns {Interval[]}
 */
const createIntervalList = function (nums) {
    /** @type {Interval[]} */
    const list = [];

    for (const n of nums) {
        const idx = binarySearch(list, n);

        if (idx == list.length) {
            list.push(new Interval(n, n + 1));
        } else {
            // const target = list[idx];
            if (n >= list[idx].begin) {
                continue;
            }

            list.splice(idx, 0, new Interval(n, n + 1));
        }

        if (idx > 0) {
            const [left, right] = [list[idx - 1], list[idx]];
            if (left.end == right.begin) {
                left.end = right.end;
                list.splice(idx, 1);
            }
        }
        if (idx < list.length - 1) {
            const [left, right] = [list[idx], list[idx + 1]];
            if (left.end == right.begin) {
                left.end = right.end;
                list.splice(idx + 1, 1);
            }
        }
    }

    return list;
};

const calcSum = function (a, n) {
    return a * n + (n * (n - 1)) / 2;
};

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
const minimalKSum = function (nums, k) {
    let result = 0;
    let start = 1;
    let remain = k;

    const list = createIntervalList(nums);
    for (const i of list) {
        if (start < i.begin) {
            const length = Math.min(i.begin - start, remain);

            result += calcSum(start, length);
            remain -= length;

            if (remain == 0) {
                break;
            }
        }

        start = i.end;
    }

    if (remain > 0) {
        const last = list[list.length - 1];
        result += calcSum(last.end, remain);
    }

    return result;
};

module.exports = minimalKSum;

