/**
 * @param {number[]} arr
 * @return {boolean}
 */
const uniqueOccurrences = function (arr) {
    const dict = new Map();

    arr.forEach((val) => {
        if (!dict.has(val)) {
            dict.set(val, 0);
        }

        dict.set(val, dict.get(val) + 1);
    });

    const keys = Array.from(dict.values());
    return keys.length === new Set(keys).size;
};

module.exports = uniqueOccurrences;
