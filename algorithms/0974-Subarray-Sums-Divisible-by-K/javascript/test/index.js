const func = require('../index');
const { deepStrictEqual } = require('power-assert');

describe('974. Subarray Sums Divisible by K', () => {
    it('A = [4,5,0,-2,-3,1], K = 5 should be 7', () => {
        deepStrictEqual(func([4, 5, 0, -2, -3, 1], 5), 7);
    });
});
