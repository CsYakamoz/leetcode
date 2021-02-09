const isQuestionMark = (c) => c === '?';

/**
 * @param {string} time
 * @return {string}
 */
const maximumTime = function(time) {
    const list = time.split('');

    if (isQuestionMark(list[0]) && isQuestionMark(list[1])) {
        list[0] = '2';
        list[1] = '3';
    } else if (isQuestionMark(list[0])) {
        list[0] = parseInt(list[1]) > 3 ? '1' : '2';
    } else if (isQuestionMark(list[1])) {
        list[1] = parseInt(list[0]) === 2 ? '3' : '9';
    }

    if (isQuestionMark(list[3])) {
        list[3] = '5';
    }

    if (isQuestionMark(list[4])) {
        list[4] = '9';
    }

    return list.join('');
};

module.exports = maximumTime;
