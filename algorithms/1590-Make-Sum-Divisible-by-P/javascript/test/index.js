const func = require('../index');
const { deepStrictEqual } = require('power-assert');

describe('1590. Make Sum Divisible by P', () => {
    it('nums = [3,1,4,2], p = 6 should be 1', () => {
        deepStrictEqual(func([3, 1, 4, 2], 6), 1);
    });

    it('nums = [6,3,5,2], p = 9 should be 2', () => {
        deepStrictEqual(func([6, 3, 5, 2], 9), 2);
    });

    it('nums = [1,2,3], p = 3 should be 0', () => {
        deepStrictEqual(func([1, 2, 3], 3), 0);
    });

    it('nums = [1,2,3], p = 7 should be -1', () => {
        deepStrictEqual(func([1, 2, 3], 7), -1);
    });

    it('nums = [1000000000,1000000000,1000000000], p = 3 should be 0', () => {
        deepStrictEqual(func([1000000000, 1000000000, 1000000000], 3), 0);
    });

    it('test case 128', () => {
        const { nums, p } = require('../../test-case/128.json');
        deepStrictEqual(func(nums, p), -1);
    });

    it('test case 138', () => {
        const { nums, p } = require('../../test-case/138.json');
        deepStrictEqual(func(nums, p), 4008);
    });

    it('test case 112', () => {
        const { nums, p } = require('../../test-case/112.json');
        deepStrictEqual(func(nums, p), -1);
    });

    it('test case 110', () => {
        const { nums, p } = require('../../test-case/110.json');
        deepStrictEqual(func(nums, p), 7);
    });

    it('test case 113', () => {
        const { nums, p } = require('../../test-case/113.json');
        deepStrictEqual(func(nums, p), 2);
    });

    it('test case 141', () => {
        const { nums, p } = require('../../test-case/141.json');
        deepStrictEqual(func(nums, p), 11);
    });
});
