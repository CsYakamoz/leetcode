const func = require('../index');
const assert = require('power-assert');

describe('678. Valid Parenthesis String', () => {
    it('s = "()" should be true', () => {
        assert.deepStrictEqual(func('()'), true);
    });

    it('s = "(*)" should be true', () => {
        assert.deepStrictEqual(func('(*)'), true);
    });

    it('s = "(*))" should be true', () => {
        assert.deepStrictEqual(func('(*))'), true);
    });

    it('s = "(*)*" should be true', () => {
        assert.deepStrictEqual(func('(*)*'), true);
    });

    it('s = "(*(*)*)(*(**)()" should be true', () => {
        assert.deepStrictEqual(func('(*(*)*)(*(**)()'), true);
    });

    it('s = "())()))*()" should be false', () => {
        assert.deepStrictEqual(func('())()))*()'), false);
    });

    it('s = "*()(())*()(()()((()(()()*)(*(())((((((((()*)(()(*)" should be false', () => {
        assert.deepStrictEqual(
            func('*()(())*()(()()((()(()()*)(*(())((((((((()*)(()(*)'),
            false
        );
    });
});
