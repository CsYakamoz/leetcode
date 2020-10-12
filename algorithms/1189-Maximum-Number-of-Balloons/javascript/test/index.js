const func = require('../index');
const assert = require('power-assert');

describe('1189. Maximum Number of Balloons', () => {
    it('text = nlaebolko should be 1', () => {
        assert.deepStrictEqual(func('nlaebolko'), 1);
    });

    it('text = loonbalxballpoon should be 2', () => {
        assert.deepStrictEqual(func('loonbalxballpoon'), 2);
    });

    it('text = leetcode', () => {
        assert.deepStrictEqual(func('leetcode'), 0);
    });
});
