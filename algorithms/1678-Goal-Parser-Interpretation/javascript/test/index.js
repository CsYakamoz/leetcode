const func = require('../index');
const { deepStrictEqual } = require('power-assert');

const testList = [
    {
        input: 'G()(al)',
        output: 'Goal',
    },
    {
        input: 'G()()()()(al)',
        output: 'Gooooal',
    },
    {
        input: '(al)G(al)()()G',
        output: 'alGalooG',
    },
];

describe('1678. Goal Parser Interpretation', () => {
    for (const [idx, test] of Object.entries(testList)) {
        it(`TestCase[${idx}]`, () => {
            deepStrictEqual(func(test.input), test.output);
        });
    }
});
