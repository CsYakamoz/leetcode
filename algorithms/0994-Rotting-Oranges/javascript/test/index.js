const func = require('../index');
const { deepStrictEqual } = require('power-assert');

describe('994. Rotting Oranges', () => {
    it('[[2,1,1],[1,1,0],[0,1,1]] should be 4', () => {
        deepStrictEqual(
            func([
                [2, 1, 1],
                [1, 1, 0],
                [0, 1, 1],
            ]),
            4
        );
    });

    it('[[2,1,1],[0,1,1],[1,0,1]] should be -1', () => {
        deepStrictEqual(
            func([
                [2, 1, 1],
                [0, 1, 1],
                [1, 0, 1],
            ]),
            -1
        );
    });

    it('[[0,2]] should be 0', () => {
        deepStrictEqual(func([[0, 2]]), 0);
    });

    it('[[1,2]] should be 1', () => {
        deepStrictEqual(func([[1, 2]]), 1);
    });
});
