const func = require('../index');
const { deepStrictEqual } = require('power-assert');
const {
    utils: { arrayToBinaryTree },
} = require('@lib/javascript');

describe('783. Minimum Distance Between BST Nodes', () => {
    it('[4,2,6,1,3,null,null] should be 2', () => {
        deepStrictEqual(
            func(arrayToBinaryTree([4, 2, 6, 1, 3, null, null])),
            1
        );
    });

    it('[8,4,11,-3,6,null,null] should be 2', () => {
        deepStrictEqual(
            func(arrayToBinaryTree([8, 4, 11, -3, 6, null, null])),
            2
        );
    });
});
