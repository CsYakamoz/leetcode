const func = require('../index');
const assert = require('power-assert');

describe('1604. Alert Using Same Key-Card Three or More Times in a One Hour Period', () => {
    it('keyName = ["daniel","daniel","daniel","luis","luis","luis","luis"], keyTime = ["10:00","10:40","11:00","09:00","11:00","13:00","15:00"] should be ["daniel"]', () => {
        assert.deepEqual(
            func(
                ['daniel', 'daniel', 'daniel', 'luis', 'luis', 'luis', 'luis'],
                ['10:00', '10:40', '11:00', '09:00', '11:00', '13:00', '15:00']
            ),
            ['daniel']
        );
    });

    it('keyName = ["alice","alice","alice","bob","bob","bob","bob"], keyTime = ["12:01","12:00","18:00","21:00","21:20","21:30","23:00"] should be ["bob"]', () => {
        assert.deepEqual(
            func(
                ['alice', 'alice', 'alice', 'bob', 'bob', 'bob', 'bob'],
                ['12:01', '12:00', '18:00', '21:00', '21:20', '21:30', '23:00']
            ),
            ['bob']
        );
    });

    it('keyName = ["john","john","john"], keyTime = ["23:58","23:59","00:01"] should be []', () => {
        assert.deepEqual(
            func(['john', 'john', 'john'], ['23:58', '23:59', '00:01']),
            []
        );
    });

    it('keyName = ["leslie","leslie","leslie","clare","clare","clare","clare"], keyTime = ["13:00","13:20","14:00","18:00","18:51","19:30","19:49"] should be ["clare","leslie"]', () => {
        assert.deepEqual(
            func(
                [
                    'leslie',
                    'leslie',
                    'leslie',
                    'clare',
                    'clare',
                    'clare',
                    'clare',
                ],
                ['13:00', '13:20', '14:00', '18:00', '18:51', '19:30', '19:49']
            ),
            ['clare', 'leslie']
        );
    });

    it('keyName = ["a","a","a","a","a","a","b","b","b","b","b"], keyTime = ["23:27","03:14","12:57","13:35","13:18","21:58","22:39","10:49","19:37","14:14","10:41"] should be ["a"]', () => {
        assert.deepEqual(
            func(
                ['a', 'a', 'a', 'a', 'a', 'a', 'b', 'b', 'b', 'b', 'b'],
                [
                    '23:27',
                    '03:14',
                    '12:57',
                    '13:35',
                    '13:18',
                    '21:58',
                    '22:39',
                    '10:49',
                    '19:37',
                    '14:14',
                    '10:41',
                ]
            ),
            ['a']
        );
    });
});
