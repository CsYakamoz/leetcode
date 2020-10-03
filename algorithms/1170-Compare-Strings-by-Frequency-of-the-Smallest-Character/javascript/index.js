/**
 * @param {string} text
 * @returns {number}
 */
const calc = (text) => {
    let freq = 1;
    let smallestChar = text.charCodeAt(0);

    for (let i = 1, length = text.length; i < length; i++) {
        const char = text.charCodeAt(i);
        if (char < smallestChar) {
            smallestChar = char;
            freq = 1;
        } else if (char === smallestChar) {
            freq += 1;
        }
    }

    return freq;
};

/**
 * @param {string} query
 * @param {number[]} sortedWords
 */
const findSuitableElem = (query, sortedWords) => {
    const f = calc(query);

    let [left, right] = [0, sortedWords.length];
    while (left < right) {
        const mid = Math.floor((right - left) / 2) + left;
        if (sortedWords[mid] > f) {
            right = mid;
        } else {
            left = mid + 1;
        }
    }

    return sortedWords.length - right;
};

/**
 * @param {string[]} queries
 * @param {string[]} words
 * @return {number[]}
 */
const numSmallerByFrequency = function(queries, words) {
    const sortedWords = words.map(calc).sort((a, b) => a - b);

    return queries.map((query) => findSuitableElem(query, sortedWords));
};

module.exports = numSmallerByFrequency;
