/**
 * @param {number[]} A
 * @return {boolean}
 */
const canThreePartsEqualSum = function(A) {
    let totalSum = 0;
    const sumList = [];

    for (const val of A) {
        totalSum += val;
        sumList.push(totalSum);
    }

    if (totalSum % 3 !== 0) {
        return false;
    }

    const partSum = totalSum / 3;

    const firstIdx = sumList.indexOf(partSum);
    if (firstIdx === -1 || firstIdx >= A.length - 2) {
        return false;
    }

    const secondIdx = sumList.indexOf(2 * partSum, firstIdx + 1);
    if (secondIdx === -1 || secondIdx >= A.length - 1) {
        return false;
    }

    return totalSum - sumList[secondIdx] === partSum;
};

module.exports = canThreePartsEqualSum;
