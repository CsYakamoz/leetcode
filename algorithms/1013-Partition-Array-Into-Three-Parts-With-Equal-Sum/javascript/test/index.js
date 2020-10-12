const func = require('../index');
const assert = require('power-assert');

describe('1013. Partition Array Into Three Parts With Equal Sum', function() {
    it('[0,2,1,-6,6,-7,9,1,2,0,1] should be true', () => {
        assert.deepStrictEqual(func([0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1]), true);
    });

    it('[0,2,1,-6,6,7,9,-1,2,0,1] should be false', () => {
        assert.deepStrictEqual(func([0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1]), false);
    });

    it('[3,3,6,5,-2,2,5,1,-9,4] should be true', () => {
        assert.deepStrictEqual(func([3, 3, 6, 5, -2, 2, 5, 1, -9, 4]), true);
    });
});
