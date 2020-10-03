/**
 * @param {string[]} arr
 * @return {string}
 */
const lastElem = (arr) => arr[arr.length - 1];

/**
 * @param {string} c1
 * @param {string} c2
 * @return {boolean}
 */
const needToRemove = (c1, c2) =>
    Math.abs(c1.charCodeAt(0) - c2.charCodeAt(0)) === 32;

/**
 * @param {string} s
 * @return {string}
 */
const makeGood = function(s) {
    if (s.length < 2) {
        return s;
    }

    const stack = [s[0]];

    for (let i = 1, len = s.length; i < len; i++) {
        const c = s[i];

        if (stack.length === 0) {
            stack.push(c);
            continue;
        }

        if (needToRemove(c, lastElem(stack))) {
            stack.pop();
            continue;
        }

        stack.push(c);
    }

    return stack.join('');
};

module.exports = makeGood;
