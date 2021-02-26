/**
 * @param {number[]} arr
 * @return {number[][]}
 */
const minimumAbsDifference = function (arr) {
    arr.sort((a, b) => a - b);

    const dict = new Map();
    let minDiff = Number.MAX_SAFE_INTEGER;
    for (let i = 1, len = arr.length; i < len; i++) {
        const diff = arr[i] - arr[i - 1];
        minDiff = Math.min(minDiff, diff);

        if (!dict.has(diff)) {
            dict.set(diff, []);
        }

        const list = dict.get(diff);
        list.push([arr[i - 1], arr[i]]);
    }

    return dict.get(minDiff) || [];
};

module.exports = minimumAbsDifference;
