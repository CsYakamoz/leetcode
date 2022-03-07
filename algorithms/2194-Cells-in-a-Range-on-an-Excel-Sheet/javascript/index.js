/**
 * @param {string} s
 * @return {string[]}
 */
const cellsInRange = function (s) {
    // col-begin, row-begin, col-end, row-end
    const [cb, rb, , ce, re] = s;
    const cBase = cb.charCodeAt(0);
    const rBase = parseInt(rb);
    const cLen = ce.charCodeAt(0) - cBase + 1;
    const rLen = parseInt(re) - rBase + 1;

    const charList = [];
    for (let i = 0; i < cLen; i++) {
        charList.push(String.fromCharCode(i + cBase));
    }

    return [...Array(cLen * rLen)].map((_, idx) => {
        const number = (idx % rLen) + rBase;
        const char = charList[parseInt(idx / rLen)];

        return char + number;
    });
};

module.exports = cellsInRange;

