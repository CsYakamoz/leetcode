/**
 * @param {number} N
 * @return {number}
 */
const fib = function(N) {
    if (N < 2) {
        return N;
    }

    let [a, b, res] = [0, 1, null];
    for (let i = 2; i <= N; i++) {
        res = a + b;
        a = b;
        b = res;
    }

    return res;
};

module.exports = fib;
