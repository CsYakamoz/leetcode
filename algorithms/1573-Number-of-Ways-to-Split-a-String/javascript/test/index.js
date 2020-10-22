const func = require('../index');
const { deepStrictEqual } = require('power-assert');

describe('1573. Number of Ways to Split a String', () => {
    it('s = "10101" should be 4', () => {
        deepStrictEqual(func('10101'), 4);
    });

    it('s = "1001" should be 0', () => {
        deepStrictEqual(func('1001'), 0);
    });

    it('s = "0000" should be 3', () => {
        deepStrictEqual(func('0000'), 3);
    });

    it('s = "000000000000" should be 55', () => {
        deepStrictEqual(func('000000000000'), 55);
    });

    it('s = "100100010100110" should be 12', () => {
        deepStrictEqual(func('100100010100110'), 12);
    });
});
