const func = require('../index');
const assert = require('power-assert');

describe('994. Rotting Oranges', () => {
    it('[[2,1,1],[1,1,0],[0,1,1]] should be 4', () => {
        assert.equal(func([[2, 1, 1], [1, 1, 0], [0, 1, 1]]), 4);
    });

    it('[[2,1,1],[0,1,1],[1,0,1]] should be -1', () => {
        assert.equal(func([[2, 1, 1], [0, 1, 1], [1, 0, 1]]), -1);
    });

    it('[[0,2]] should be 0', () => {
        assert.equal(func([[0, 2]]), 0);
    });

    it('[[1,2]] should be 1', () => {
        assert.equal(func([[1, 2]]), 1);
    });
});
