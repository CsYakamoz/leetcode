/**
 * @param {number} n
 * @return {number}
 */
const tribonacci = function(n) {
    const constant = [0, 1, 1];
    if (n < 3) {
        return constant[n];
    }

    let [a, b, c, res] = [...constant, null];
    for (let i = 3; i <= n; i++) {
        res = a + b + c;
        a = b;
        b = c;
        c = res;
    }

    return res;
};

module.exports = tribonacci;
