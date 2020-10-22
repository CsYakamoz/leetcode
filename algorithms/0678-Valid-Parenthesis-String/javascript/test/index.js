const func = require('../index');
const { deepStrictEqual } = require('power-assert');

describe('678. Valid Parenthesis String', () => {
    it('s = "()" should be true', () => {
        deepStrictEqual(func('()'), true);
    });

    it('s = "(*)" should be true', () => {
        deepStrictEqual(func('(*)'), true);
    });

    it('s = "(*))" should be true', () => {
        deepStrictEqual(func('(*))'), true);
    });

    it('s = "(*)*" should be true', () => {
        deepStrictEqual(func('(*)*'), true);
    });

    it('s = "(*(*)*)(*(**)()" should be true', () => {
        deepStrictEqual(func('(*(*)*)(*(**)()'), true);
    });

    it('s = "())()))*()" should be false', () => {
        deepStrictEqual(func('())()))*()'), false);
    });

    it('s = "*()(())*()(()()((()(()()*)(*(())((((((((()*)(()(*)" should be false', () => {
        deepStrictEqual(
            func('*()(())*()(()()((()(()()*)(*(())((((((((()*)(()(*)'),
            false
        );
    });
});
