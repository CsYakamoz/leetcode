const func = require('../index');
const assert = require('power-assert');

describe('1170. Compare Strings by Frequency of the Smallest Character', () => {
    it('queries = ["cbd"], words = ["zaaaz"] should be [1]', () => {
        assert.deepEqual(func(['cbd'], ['zaaaz']), [1]);
    });

    it('queries = ["bbb","cc"], words = ["a","aa","aaa","aaaa"] should be [1, 2]', () => {
        assert.deepEqual(func(['bbb', 'cc'], ['a', 'aa', 'aaa', 'aaaa']), [
            1,
            2,
        ]);
    });
});
