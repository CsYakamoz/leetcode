const func = require('../index');
const { deepStrictEqual } = require('power-assert');

describe('1184. Distance Between Bus Stops', () => {
    it('distance = [1,2,3,4], start = 0, destination = 1 should be 1', () => {
        deepStrictEqual(func([1, 2, 3, 4], 0, 1), 1);
    });

    it('distance = [1,2,3,4], start = 0, destination = 2 should be 3', () => {
        deepStrictEqual(func([1, 2, 3, 4], 0, 2), 3);
    });

    it('distance = [1,2,3,4], start = 0, destination = 3 should be 4', () => {
        deepStrictEqual(func([1, 2, 3, 4], 0, 3), 4);
    });
});
