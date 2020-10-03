const func = require('../index');
const assert = require('power-assert');

describe('1137. N-th Tribonacci Number', () => {
    it('n = 4 should be 4', function() {
        assert.equal(func(4), 4);
    });

    it('n = 25 should be 1389537', function() {
        assert.equal(func(25), 1389537);
    });

    it('n = 37 should be 2082876103', function() {
        assert.equal(func(37), 2082876103);
    });
});
