/**
 * @param {number[]} arr1
 * @param {number[]} arr2
 * @returns {number}
 */
const findMinDiff = (arr1, arr2) => {
    let minDiff = Number.MAX_SAFE_INTEGER;
    for (const i of arr1) {
        let begin = 0;
        let end = arr2.length;

        while (begin < end) {
            const mid = Math.floor((end - begin) / 2) + begin;
            if (arr2[mid] > i) {
                end = mid;
            } else {
                begin = mid + 1;
            }
        }
        if (end !== arr2.length) {
            minDiff = Math.min(minDiff, arr2[end] - i);
        }
    }

    return minDiff;
};

/**
 * @param {number[]} nums
 * @param {number} p
 * @return {number}
 */
const minSubarray = function(nums, p) {
    /** @type {Map<number, number[]>} */
    const prefixRemainderDict = new Map();

    const expRemainder = nums.reduce((acc, curr, idx) => {
        const remainder = (acc + curr) % p;
        if (!prefixRemainderDict.has(remainder)) {
            prefixRemainderDict.set(remainder, []);
        }
        prefixRemainderDict.get(remainder).push(idx);

        return remainder;
    }, 0);

    if (expRemainder === 0) {
        return 0;
    }

    let minLength = Number.MAX_SAFE_INTEGER;
    for (const [currRemainder, idxList] of prefixRemainderDict) {
        const anotherRemainder = (expRemainder + currRemainder) % p;

        if (!prefixRemainderDict.has(anotherRemainder)) {
            continue;
        }

        minLength = Math.min(
            minLength,
            findMinDiff(idxList, prefixRemainderDict.get(anotherRemainder))
        );
    }

    if (minLength !== Number.MAX_SAFE_INTEGER) {
        return minLength;
    } else if (prefixRemainderDict.has(expRemainder)) {
        const length = prefixRemainderDict.get(expRemainder)[0] + 1;
        return length !== nums.length ? length : -1;
    } else {
        return -1;
    }
};

module.exports = minSubarray;
