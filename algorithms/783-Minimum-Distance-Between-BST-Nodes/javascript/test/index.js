const func = require('../index');
const assert = require('power-assert');
const { arrayToBinaryTree } = require('utils');

describe('783. Minimum Distance Between BST Nodes', () => {
    it('[4,2,6,1,3,null,null] should be 2', () => {
        assert.equal(func(arrayToBinaryTree([4, 2, 6, 1, 3, null, null])), 1);
    });

    it('[8,4,11,-3,6,null,null] should be 2', () => {
        assert.equal(func(arrayToBinaryTree([8, 4, 11, -3, 6, null, null])), 2);
    });
});
