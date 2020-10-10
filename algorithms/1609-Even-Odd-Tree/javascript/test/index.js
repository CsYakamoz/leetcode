const func = require('../index');
const assert = require('power-assert');
const { arrayToBinaryTree } = require('@javascript/language-helper');

describe('1609. Even Odd Tree', () => {
    it('root = [1,10,4,3,null,7,9,12,8,6,null,null,2] should be true', () => {
        assert.equal(
            func(
                arrayToBinaryTree([
                    1,
                    10,
                    4,
                    3,
                    null,
                    7,
                    9,
                    12,
                    8,
                    6,
                    null,
                    null,
                    2,
                ])
            ),
            true
        );
    });

    it('root = [5,4,2,3,3,7] should be false', () => {
        assert.equal(func(arrayToBinaryTree([5, 4, 2, 3, 3, 7])), false);
    });

    it('root = [5,9,1,3,5,7] should be false', () => {
        assert.equal(func(arrayToBinaryTree([5, 9, 1, 3, 5, 7])), false);
    });

    it('root = [1] should be true', () => {
        assert.equal(func(arrayToBinaryTree([1])), true);
    });

    it('root = [11,8,6,1,3,9,11,30,20,18,16,12,10,4,2,17] should be true', () => {
        assert.equal(
            func(
                arrayToBinaryTree([
                    11,
                    8,
                    6,
                    1,
                    3,
                    9,
                    11,
                    30,
                    20,
                    18,
                    16,
                    12,
                    10,
                    4,
                    2,
                    17,
                ])
            ),
            true
        );
    });
});
