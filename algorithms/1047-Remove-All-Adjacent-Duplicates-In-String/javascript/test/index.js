const func = require('../index');
const assert = require('power-assert');

describe('1047. Remove All Adjacent Duplicates In String', function() {
    it('abbaca should be ca', function() {
        assert.equal(func('abbaca'), 'ca');
    });

    it('aaaa should be empty', function() {
        assert.equal(func('aaaa'), '');
    });

    it('aaabaaa should be aba', function() {
        assert.equal(func('aaabaaa'), 'aba');
    });

    it('abaabaaba should be aba', function() {
        assert.equal(func('abaabaaba'), 'aba');
    });

    it('abacabacaba should be abacabacaba', function() {
        assert.equal(func('abacabacaba'), 'abacabacaba');
    });
});
