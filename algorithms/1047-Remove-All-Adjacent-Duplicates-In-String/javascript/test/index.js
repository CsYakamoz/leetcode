const func = require('../index');
const { deepStrictEqual } = require('power-assert');

describe('1047. Remove All Adjacent Duplicates In String', function () {
    it('abbaca should be ca', function () {
        deepStrictEqual(func('abbaca'), 'ca');
    });

    it('aaaa should be empty', function () {
        deepStrictEqual(func('aaaa'), '');
    });

    it('aaabaaa should be aba', function () {
        deepStrictEqual(func('aaabaaa'), 'aba');
    });

    it('abaabaaba should be aba', function () {
        deepStrictEqual(func('abaabaaba'), 'aba');
    });

    it('abacabacaba should be abacabacaba', function () {
        deepStrictEqual(func('abacabacaba'), 'abacabacaba');
    });
});
