const func = require('../index');
const { deepStrictEqual } = require('power-assert');

const testList = [
    {
        input: [
            [1, 1, 1],
            [1, 0, 1],
            [1, 1, 1],
        ],
        output: [
            [1, 0, 1],
            [0, 0, 0],
            [1, 0, 1],
        ],
    },

    {
        input: [
            [0, 1, 2, 0],
            [3, 4, 5, 2],
            [1, 3, 1, 5],
        ],
        output: [
            [0, 0, 0, 0],
            [0, 4, 5, 0],
            [0, 3, 1, 0],
        ],
    },
];

describe('73. Set Matrix Zeroes', () => {
    for (const [idx, test] of Object.entries(testList)) {
        it(`TestCase[${idx}]`, () => {
            const copyData = JSON.parse(JSON.stringify(test.input));

            func(copyData);

            deepStrictEqual(copyData, test.output);
        });
    }
});
