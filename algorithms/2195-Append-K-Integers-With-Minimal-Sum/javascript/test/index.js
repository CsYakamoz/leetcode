const func = require('../index');
const { deepStrictEqual } = require('power-assert');

const testList = [
    {
        nums: [1, 4, 25, 10, 25],
        k: 2,
        output: 5,
    },
    {
        nums: [5, 6],
        k: 6,
        output: 25,
    },
    {
        nums: [5, 6, 9, 10, 12, 16],
        k: 10,
        output: 78,
    },
    {
        nums: [99, 74, 32, 41, 231, 534, 123, 55, 31],
        k: 1000,
        output: 508325,
    },
    {
        nums: [
            123,
            1423,
            4321,
            3121,
            3,
            123,
            12,
            3,
            4,
            1,
            2,
            3,
            1,
            231,
            12,
            3,
            1,
            23,
            13,
            13,
        ],
        k: 100000,
        output: 508325,
    },
];

describe('2195. Append K Integers With Minimal Sum', () => {
    for (const [idx, test] of Object.entries(testList)) {
        it(`TestCase[${idx}]`, () => {
            deepStrictEqual(func(test.nums, test.k), test.output);
        });
    }
});

