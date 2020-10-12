const func = require('../index');
const assert = require('power-assert');

describe('1047. Remove All Adjacent Duplicates In String', function() {
    it('abbaca should be ca', function() {
        assert.deepStrictEqual(func('abbaca'), 'ca');
    });

    it('aaaa should be empty', function() {
        assert.deepStrictEqual(func('aaaa'), '');
    });

    it('aaabaaa should be aba', function() {
        assert.deepStrictEqual(func('aaabaaa'), 'aba');
    });

    it('abaabaaba should be aba', function() {
        assert.deepStrictEqual(func('abaabaaba'), 'aba');
    });

    it('abacabacaba should be abacabacaba', function() {
        assert.deepStrictEqual(func('abacabacaba'), 'abacabacaba');
    });
});
