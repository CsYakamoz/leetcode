const func = require('../index');
const assert = require('power-assert');

describe('1184. Distance Between Bus Stops', () => {
    it('distance = [1,2,3,4], start = 0, destination = 1 should be 1', () => {
        assert.deepStrictEqual(func([1, 2, 3, 4], 0, 1), 1);
    });

    it('distance = [1,2,3,4], start = 0, destination = 2 should be 3', () => {
        assert.deepStrictEqual(func([1, 2, 3, 4], 0, 2), 3);
    });

    it('distance = [1,2,3,4], start = 0, destination = 3 should be 4', () => {
        assert.deepStrictEqual(func([1, 2, 3, 4], 0, 3), 4);
    });
});
