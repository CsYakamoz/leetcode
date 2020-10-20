/**
 * @param {number[]} A
 * @param {number} K
 * @return {number}
 */
const subarraysDivByK = function (A, K) {
    const prefixRemainderDict = new Map();

    let result = 0;
    let remainder = 0;
    for (let i = 0, len = A.length; i < len; i++) {
        remainder = (remainder + A[i]) % K;

        if (remainder === 0) {
            result += 1;
        } else {
            const anotherRemainder =
                remainder - K * (remainder / Math.abs(remainder));
            result += prefixRemainderDict.get(anotherRemainder) || 0;
        }

        const count = prefixRemainderDict.get(remainder) || 0;
        result += count;
        prefixRemainderDict.set(remainder, count + 1);
    }

    return result;
};

module.exports = subarraysDivByK;

