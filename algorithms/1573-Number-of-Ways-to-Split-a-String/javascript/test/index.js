const func = require('../index');
const assert = require('power-assert');

describe('1573. Number of Ways to Split a String', () => {
    it('s = "10101" should be 4', () => {
        assert.deepStrictEqual(func('10101'), 4);
    });

    it('s = "1001" should be 0', () => {
        assert.deepStrictEqual(func('1001'), 0);
    });

    it('s = "0000" should be 3', () => {
        assert.deepStrictEqual(func('0000'), 3);
    });

    it('s = "000000000000" should be 55', () => {
        assert.deepStrictEqual(func('000000000000'), 55);
    });

    it('s = "100100010100110" should be 12', () => {
        assert.deepStrictEqual(func('100100010100110'), 12);
    });
});
