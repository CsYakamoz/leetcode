/**
 * @param {string} text
 * @returns {object}
 */
const createCountDict = function(text) {
    return text.split('').reduce((prev, curr) => {
        if (prev[curr] === undefined) {
            prev[curr] = 1;
        } else {
            prev[curr] += 1;
        }

        return prev;
    }, {});
};

/**
 * @param {string} text
 * @return {number}
 */
const maxNumberOfBalloons = function(text) {
    const keyword = 'balloon';
    const keywordDict = createCountDict(keyword);
    const textDict = createCountDict(text);

    let maxNumber = Number.MAX_SAFE_INTEGER;
    for (const c in keywordDict) {
        const division = keywordDict[c];
        const count = Math.floor((textDict[c] || 0) / division);
        maxNumber = Math.min(maxNumber, count);
    }

    return maxNumber;
};

module.exports = maxNumberOfBalloons;
