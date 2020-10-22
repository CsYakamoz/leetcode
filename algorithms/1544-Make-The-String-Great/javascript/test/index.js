const func = require('../index');
const { deepStrictEqual } = require('power-assert');

describe('1544. Make The String Great', () => {
    it('s = "leEeetcode" should be "leetcode"', () => {
        deepStrictEqual(func('leEeetcode'), 'leetcode');
    });

    it('s = "abBAcC" should be ""', () => {
        deepStrictEqual(func('abBAcC'), '');
    });
});
