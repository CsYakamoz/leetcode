const func = require('../index');
const { deepStrictEqual } = require('power-assert');
const {
    utils: { preOrderTraversal },
} = require('@lib/javascript');

describe('1008. Construct Binary Search Tree from Preorder Traversal', () => {
    it('[8,5,1,7,10,12]', function() {
        const input = [8, 5, 1, 7, 10, 12];
        deepStrictEqual(preOrderTraversal(func(input)), input);
    });

    it('[8,5,1,0,2,7,6,10,9,12,11,13]', function() {
        const input = [8, 5, 1, 0, 2, 7, 6, 10, 9, 12, 11, 13];
        deepStrictEqual(preOrderTraversal(func(input)), input);
    });
});
