/**
 * @param {string} text
 * @return {string}
 */
const arrangeWords = function(text) {
    const wordList = text.toLowerCase().split(' ');

    wordList.sort((a, b) => {
        if (a.length !== b.length) {
            return a.length - b.length;
        }

        return 0;
    });

    const sortText = wordList.join(' ');

    return sortText.substring(0, 1).toUpperCase() + sortText.substring(1);
};

module.exports = arrangeWords;
