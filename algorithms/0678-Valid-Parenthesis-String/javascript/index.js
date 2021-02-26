/**
 * @param {any[]} arr
 * @return {any}
 */
const lastElemOrNeg1 = (arr) => (arr.length !== 0 ? arr[arr.length - 1] : -1);

/**
 * @param {string} s
 * @return {boolean}
 */
const checkValidString = function (s) {
    const starStack = [];
    const leftBrackets = [];

    for (let i = 0, len = s.length; i < len; i++) {
        const char = s[i];

        switch (char) {
            case ')':
                {
                    if (leftBrackets.length !== 0) {
                        leftBrackets.pop();
                    } else if (starStack.length !== 0) {
                        starStack.pop();
                    } else {
                        return false;
                    }
                }
                break;
            case '*':
                starStack.push(i);
                break;
            case '(':
                leftBrackets.push(i);
                break;
            default:
                throw new Error(`unknown char: ${char}`);
        }
    }

    while (
        leftBrackets.length !== 0 &&
        lastElemOrNeg1(starStack) > lastElemOrNeg1(leftBrackets)
    ) {
        starStack.pop();
        leftBrackets.pop();
    }

    return leftBrackets.length === 0;
};

module.exports = checkValidString;
