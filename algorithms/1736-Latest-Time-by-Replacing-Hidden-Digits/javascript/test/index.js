const func = require('../index');
const { deepStrictEqual } = require('power-assert');

const testList = [
    { input: '2?:?0', output: '23:50' },
    { input: '0?:3?', output: '09:39' },
    { input: '1?:22', output: '19:22' },
];

describe('1736. Latest Time by Replacing Hidden Digits', () => {
    for (const [idx, test] of Object.entries(testList)) {
        it(`TestCase[${idx}]`, () => {
            deepStrictEqual(func(test.input), test.output);
        });
    }
});
