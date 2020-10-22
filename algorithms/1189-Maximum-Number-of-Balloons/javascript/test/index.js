const func = require('../index');
const { deepStrictEqual } = require('power-assert');

describe('1189. Maximum Number of Balloons', () => {
    it('text = nlaebolko should be 1', () => {
        deepStrictEqual(func('nlaebolko'), 1);
    });

    it('text = loonbalxballpoon should be 2', () => {
        deepStrictEqual(func('loonbalxballpoon'), 2);
    });

    it('text = leetcode', () => {
        deepStrictEqual(func('leetcode'), 0);
    });
});
