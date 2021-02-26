/**
 * @param {string} command
 * @return {string}
 */
const interpret = function (command) {
    const stack = [];
    const resultArr = [];

    for (let idx = 0, len = command.length; idx < len; idx++) {
        const c = command[idx];
        switch (c) {
            case 'G':
                resultArr.push(c);
                break;
            case '(':
                stack.push(idx);
                break;
            case ')':
                {
                    const lasted = stack.pop();
                    const str = command.slice(lasted + 1, idx);
                    resultArr.push(str.length === 0 ? 'o' : 'al');
                }
                break;
        }
    }

    return resultArr.join('');
};

module.exports = interpret;
