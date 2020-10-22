const func = require('../index');
const { deepStrictEqual } = require('power-assert');

describe('1046. Last Stone Weight', () => {
    it('[] should be 0', function() {
        deepStrictEqual(func([]), 0);
    });

    it('[1] should be 1', function() {
        deepStrictEqual(func([1]), 1);
    });

    it('[2,7,4,1,8,1] should be 1', function() {
        deepStrictEqual(func([2, 7, 4, 1, 8, 1]), 1);
    });

    it('[2,7,4,1,8,1,1] should be 0', function() {
        deepStrictEqual(func([2, 7, 4, 1, 8, 1, 1]), 0);
    });

    it('[7,6,7,6,9] should 3', function() {
        deepStrictEqual(func([7, 6, 7, 6, 9]), 3);
    });

    it('[2,7,4,1,8,1,1,3,2,1,3,2,4,1,2,3,1] should be 0', function() {
        deepStrictEqual(
            func([2, 7, 4, 1, 8, 1, 1, 3, 2, 1, 3, 2, 4, 1, 2, 3, 1]),
            0
        );
    });

    it('[2,7,4,11,1,3,4,1,7,8,6,4,3,1] should 0', function() {
        deepStrictEqual(func([2, 7, 4, 11, 1, 3, 4, 1, 7, 8, 6, 4, 3, 1]), 0);
    });
});
