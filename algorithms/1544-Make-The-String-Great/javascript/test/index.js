const func = require('../index');
const assert = require('power-assert');

describe('1544. Make The String Great', () => {
    it('s = "leEeetcode" should be "leetcode"', () => {
        assert.deepStrictEqual(func('leEeetcode'), 'leetcode');
    });

    it('s = "abBAcC" should be ""', () => {
        assert.deepStrictEqual(func('abBAcC'), '');
    });
});
