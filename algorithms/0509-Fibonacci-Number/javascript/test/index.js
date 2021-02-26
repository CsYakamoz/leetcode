const func = require('../index');
const { deepStrictEqual } = require('power-assert');

describe('509. Fibonacci Number', () => {
    it('N = 2 should be 1', function () {
        deepStrictEqual(func(2), 1);
    });

    it('N = 3 should be 2', function () {
        deepStrictEqual(func(3), 2);
    });

    it('N = 4 should be 3', function () {
        deepStrictEqual(func(4), 3);
    });
});
