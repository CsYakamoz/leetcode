/**
 * @param {number[]} A
 * @param {number} K
 * @return {number}
 */
const largestSumAfterKNegations = function(A, K) {
    let positiveSum = 0;
    let negativeSum = 0;
    let maxNegative = Number.MIN_SAFE_INTEGER;
    let minPositive = Number.MAX_SAFE_INTEGER;

    const negativeAndZeroList = [];

    for (const val of A) {
        if (val > 0) {
            positiveSum += val;
            minPositive = Math.min(minPositive, val);
        } else {
            negativeAndZeroList.push(val);
            maxNegative = Math.max(maxNegative, val);
            negativeSum += val;
        }
    }

    const nzLen = negativeAndZeroList.length;

    if (nzLen === 0) {
        return positiveSum - (K % 2 === 1 ? minPositive * 2 : 0);
    }

    if (nzLen <= K) {
        positiveSum -= negativeSum;
        negativeSum = 0;

        const remain = K - nzLen;
        if (remain % 2 !== 0) {
            const min = Math.min(minPositive, -maxNegative);
            negativeSum -= min;
            positiveSum -= min;
        }
    } else {
        negativeAndZeroList.sort((a, b) => a - b);
        for (let i = 0; i < K; i++) {
            const val = negativeAndZeroList[i];
            negativeSum -= val;
            positiveSum -= val;
        }
    }

    return positiveSum + negativeSum;
};

module.exports = largestSumAfterKNegations;
