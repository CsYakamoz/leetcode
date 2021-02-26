/**
 * @param {string} str
 * @returns {number[]}
 */
const getHourAndMinute = (str) =>
    [str.slice(0, 2), str.slice(3, 5)].map((item) => parseInt(item));

/**
 * @param {string[]} timeList
 * @returns {boolean}
 */
const check = (timeList) => {
    if (timeList.length < 3) {
        return false;
    }

    let [begin, end] = [0, 1];
    while (end < timeList.length) {
        const [beginHH, beginMM] = getHourAndMinute(timeList[begin]);
        const [endHH, endMM] = getHourAndMinute(timeList[end]);

        if (beginHH === endHH || (endHH - beginHH === 1 && beginMM >= endMM)) {
            end++;
        } else {
            begin++;
        }

        if (end - begin >= 3) {
            return true;
        }
    }

    return end - begin >= 3;
};

/**
 * @param {string[]} keyName
 * @param {string[]} keyTime
 * @return {string[]}
 */
const alertNames = function (keyName, keyTime) {
    /** @type {Map<string, string[]} */
    const dict = new Map();
    for (let i = 0, len = keyName.length; i < len; i++) {
        if (!dict.has(keyName[i])) {
            dict.set(keyName[i], []);
        }

        dict.get(keyName[i]).push(keyTime[i]);
    }

    /** @type {string[]} */
    const result = [];
    for (const [key, value] of dict) {
        if (check(value.sort())) {
            result.push(key);
        }
    }

    return result.sort();
};

module.exports = alertNames;
