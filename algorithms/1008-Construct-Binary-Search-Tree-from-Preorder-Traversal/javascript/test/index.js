const func = require('../index');
const assert = require('power-assert');

const { preOrder } = require('utils');

describe('1008. Construct Binary Search Tree from Preorder Traversal', () => {
    it('[8,5,1,7,10,12]', function() {
        const input = [8, 5, 1, 7, 10, 12];
        assert.deepEqual(preOrder(func(input)), input);
    });

    it('[8,5,1,0,2,7,6,10,9,12,11,13]', function() {
        const input = [8, 5, 1, 0, 2, 7, 6, 10, 9, 12, 11, 13];
        assert.deepEqual(preOrder(func(input)), input);
    });
});
