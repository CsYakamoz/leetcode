/**
 * @param {number[][]} matrix
 * @return {void} Do not return anything, modify matrix in-place instead.
 */
const setZeroes = function (matrix) {
    const m = matrix.length;
    const n = matrix[0].length;

    let setRow = false;
    let setCol = false;
    for (let i = 0; i < m; i++) {
        if (matrix[i][0] === 0) {
            setCol = true;
            break;
        }
    }
    for (let j = 0; j < n; j++) {
        if (matrix[0][j] === 0) {
            setRow = true;
            break;
        }
    }

    for (let i = 1; i < m; i++) {
        for (let j = 1; j < n; j++) {
            if (matrix[i][j] === 0) {
                matrix[i][0] = 0;
                matrix[0][j] = 0;
            }
        }
    }

    for (let i = 1; i < m; i++) {
        if (matrix[i][0] === 0) {
            for (let j = 1; j < n; j++) {
                matrix[i][j] = 0;
            }
        }
    }
    for (let j = 1; j < n; j++) {
        if (matrix[0][j] === 0) {
            for (let i = 1; i < m; i++) {
                matrix[i][j] = 0;
            }
        }
    }

    if (setRow) {
        for (let j = 0; j < n; j++) {
            matrix[0][j] = 0;
        }
    }

    if (setCol) {
        for (let i = 0; i < m; i++) {
            matrix[i][0] = 0;
        }
    }
};

module.exports = setZeroes;
