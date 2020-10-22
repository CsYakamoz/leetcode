const func = require('../index');
const { deepStrictEqual } = require('power-assert');

describe('1207. Unique Number of Occurrences', () => {
    it('[1,2,2,1,1,3] should be true', () => {
        deepStrictEqual(func([1, 2, 2, 1, 1, 3]), true);
    });

    it('[1,2] should be false', () => {
        deepStrictEqual(func([1, 2]), false);
    });

    it('[-3,0,1,-3,1,1,1,-3,10,0] should be true', () => {
        deepStrictEqual(func([-3, 0, 1, -3, 1, 1, 1, -3, 10, 0]), true);
    });
});
