const func = require('../index');
const assert = require('power-assert');

describe('678. Valid Parenthesis String', () => {
    it('s = "()" should be true', () => {
        assert.equal(func('()'), true);
    });

    it('s = "(*)" should be true', () => {
        assert.equal(func('(*)'), true);
    });

    it('s = "(*))" should be true', () => {
        assert.equal(func('(*))'), true);
    });

    it('s = "(*)*" should be true', () => {
        assert.equal(func('(*)*'), true);
    });

    it('s = "(*(*)*)(*(**)()" should be true', () => {
        assert.equal(func('(*(*)*)(*(**)()'), true);
    });

    it('s = "())()))*()" should be false', () => {
        assert.equal(func('())()))*()'), false);
    });

    it('s = "*()(())*()(()()((()(()()*)(*(())((((((((()*)(()(*)" should be false', () => {
        assert.equal(
            func('*()(())*()(()()((()(()()*)(*(())((((((((()*)(()(*)'),
            false
        );
    });
});
