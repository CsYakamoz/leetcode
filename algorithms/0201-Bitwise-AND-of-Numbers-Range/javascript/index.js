/**
 * @param {number} diff
 * @return {number}
 */
const findTarget = (diff) => {
    let target = 1;

    while (target <= diff) {
        target *= 2;
    }

    return ~(target - 1);
};

/**
 * @param {number} m
 * @param {number} n
 * @return {number}
 */
const rangeBitwiseAnd = function(m, n) {
    if (m === n) {
        return m;
    }

    return m & n & findTarget(n - m);
};

module.exports = rangeBitwiseAnd;
