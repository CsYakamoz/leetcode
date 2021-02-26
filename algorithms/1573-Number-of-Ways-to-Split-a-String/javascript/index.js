/**
 * @param {string} s
 * @return {number}
 */
const numWays = function (s) {
    const oneIdxRecord = [];
    const modulo = 10 ** 9 + 7;
    for (let i = 0, len = s.length; i < len; i++) {
        if (s[i] === '0') {
            continue;
        }

        oneIdxRecord.push(i);
    }

    if (oneIdxRecord.length % 3 !== 0) {
        return 0;
    }

    if (oneIdxRecord.length === 0) {
        return (((s.length - 1) * (s.length - 2)) / 2) % modulo;
    }

    const partCount = oneIdxRecord.length / 3;
    const firstPartEndingIdx = oneIdxRecord[partCount - 1];
    const secondPartStartingIdx = oneIdxRecord[partCount];
    const secondPartEndingIdx = oneIdxRecord[partCount * 2 - 1];
    const thirdPartStartingIdx = oneIdxRecord[partCount * 2];

    const zeroCountBtwFirstAndSecond =
        secondPartStartingIdx - firstPartEndingIdx - 1;
    const zeroCountBtwSecondeAndThird =
        thirdPartStartingIdx - secondPartEndingIdx - 1;

    return (
        ((zeroCountBtwFirstAndSecond + 1) * (zeroCountBtwSecondeAndThird + 1)) %
        modulo
    );
};

module.exports = numWays;
