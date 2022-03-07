const func = require('../index');
const { deepStrictEqual } = require('power-assert');

const testList = [
    {
        input: 'K1:L2',
        output: ['K1', 'K2', 'L1', 'L2'],
    },
    {
        input: 'A1:F1',
        output: ['A1', 'B1', 'C1', 'D1', 'E1', 'F1'],
    },
];

describe('2194. Cells in a Range on an Excel Sheet', () => {
    for (const [idx, test] of Object.entries(testList)) {
        it(`TestCase[${idx}]`, () => {
            deepStrictEqual(func(test.input), test.output);
        });
    }
});

