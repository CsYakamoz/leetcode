/**
 * @param {string} S
 * @return {string}
 */
var removeDuplicates = function(S) {
    const stack = [];

    for (const c of S) {
        stack.push(c);

        while (stack.length > 1) {
            const last1 = stack.pop();
            const last2 = stack.pop();
            if (last1 !== last2) {
                stack.push(last2);
                stack.push(last1);
                break;
            }
        }
    }

    return stack.join('');
};

module.exports = removeDuplicates;
