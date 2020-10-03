const func = require('../index');
const assert = require('power-assert');

describe('1200. Minimum Absolute Difference', () => {
    it('[4,2,1,3] should be [[1,2],[2,3],[3,4]]', function() {
        assert.deepEqual(func([4, 2, 1, 3]), [
            [1, 2],
            [2, 3],
            [3, 4],
        ]);
    });

    it('[1,3,6,10,15] should be [[1,3]]', () => {
        assert.deepEqual(func([1, 3, 6, 10, 15]), [[1, 3]]);
    });

    it('[3,8,-10,23,19,-4,-14,27] should be [[-14,-10],[19,23],[23,27]]', () => {
        assert.deepEqual(func([3, 8, -10, 23, 19, -4, -14, 27]), [
            [-14, -10],
            [19, 23],
            [23, 27],
        ]);
    });
});
