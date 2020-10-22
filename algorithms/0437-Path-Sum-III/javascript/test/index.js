const func = require('../index');
const { deepStrictEqual } = require('power-assert');
const {
    utils: { arrayToBinaryTree },
} = require('@lib/javascript');

describe('437. Path Sum III', () => {
    it('root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8, should be 3', function() {
        deepStrictEqual(
            func(
                arrayToBinaryTree([10, 5, -3, 3, 2, null, 11, 3, -2, null, 1]),
                8
            ),
            3
        );
    });

    it('root = [1,3, 4, null, 5, 2, -1, -3, 2, 3, 1,2], sum = 3, should be 4', function() {
        deepStrictEqual(
            func(
                arrayToBinaryTree([1, 3, 4, null, 5, 2, -1, -3, 2, 3, 1, 2]),
                3
            ),
            4
        );
    });

    it('root = [1,3, 4, null, 5, 2, -1, -3, 2, 3, 1,2], sum = 4, should be 3', function() {
        deepStrictEqual(
            func(
                arrayToBinaryTree([1, 3, 4, null, 5, 2, -1, -3, 2, 3, 1, 2]),
                4
            ),
            3
        );
    });
});
