const func = require('../index');
const { deepStrictEqual } = require('power-assert');
const {
    utils: { arrayToBinaryTree },
} = require('@lib/javascript');

describe('1305. All Elements in Two Binary Search Trees', () => {
    it('root1 = [2,1,4], root2 = [1,0,3] should be [0,1,1,2,3,4]', () => {
        deepStrictEqual(
            func(arrayToBinaryTree([2, 1, 4]), arrayToBinaryTree([1, 0, 3])),
            [0, 1, 1, 2, 3, 4]
        );
    });

    it('root1 = [0,-10,10], root2 = [5,1,7,0,2] should be [-10,0,0,1,2,5,7,10]', () => {
        deepStrictEqual(
            func(
                arrayToBinaryTree([0, -10, 10]),
                arrayToBinaryTree([5, 1, 7, 0, 2])
            ),
            [-10, 0, 0, 1, 2, 5, 7, 10]
        );
    });

    it('root1 = [], root2 = [5,1,7,0,2] should be [0,1,2,5,7]', () => {
        deepStrictEqual(
            func(arrayToBinaryTree([]), arrayToBinaryTree([5, 1, 7, 0, 2])),
            [0, 1, 2, 5, 7]
        );
    });

    it('root1 = [0,-10,10], root2 = [] should be [-10,0,10]', () => {
        deepStrictEqual(
            func(arrayToBinaryTree([0, -10, 10]), arrayToBinaryTree([])),
            [-10, 0, 10]
        );
    });

    it('root1 = [1,null,8], root2 = [8,1] should be [1,1,8,8]', () => {
        deepStrictEqual(
            func(arrayToBinaryTree([1, null, 8]), arrayToBinaryTree([8, 1])),
            [1, 1, 8, 8]
        );
    });
});
