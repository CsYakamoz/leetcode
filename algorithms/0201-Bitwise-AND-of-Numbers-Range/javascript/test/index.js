const func = require('../index');
const { deepStrictEqual } = require('power-assert');

describe('201. Bitwise AND of Numbers Range', function() {
    it('[5,7] should be 4', function() {
        deepStrictEqual(func(5, 7), 4);
    });

    it('[0, 1] should be 0', function() {
        deepStrictEqual(func(0, 1), 0);
    });

    it('[1024, 1025] should be 1024', function() {
        deepStrictEqual(func(1024, 1025), 1024);
    });

    it('[2147483646, 2147483647] should be 2147483646', function() {
        deepStrictEqual(func(2147483646, 2147483647), 2147483646);
    });

    it('[2147483645, 2147483647] should be 2147483646', function() {
        deepStrictEqual(func(2147483645, 2147483647), 2147483644);
    });
});
