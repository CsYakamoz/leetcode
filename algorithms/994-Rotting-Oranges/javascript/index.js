/**
 * @param {number[][]} grid
 * @return {number}
 */
const orangesRotting = function(grid) {
    /** @type {Set<number>} */
    const rottingIdx = new Set();
    /** @type {Set<number>} */
    const freshIdx = new Set();

    const rowLength = grid.length;
    const colLength = grid[0].length;
    for (let i = 0; i < rowLength; i++) {
        for (let j = 0; j < colLength; j++) {
            const val = grid[i][j];
            if (val === 1) {
                freshIdx.add(i * colLength + j);
            } else if (val === 2) {
                rottingIdx.add(i * colLength + j);
            }
        }
    }

    // there are already no fresh oranges at minute 0
    if (freshIdx.size === 0) {
        return 0;
    }

    let minutes = 0;
    while (freshIdx.size !== 0) {
        let canRotten = false;

        const idxList = Array.from(rottingIdx.values());
        for (const idx of idxList) {
            const [row, col] = [Math.floor(idx / colLength), idx % colLength];

            // up
            const upIdx = idx - colLength;
            if (row > 0 && freshIdx.has(upIdx)) {
                canRotten = true;
                rottingIdx.add(upIdx);
                freshIdx.delete(upIdx);
            }

            // down
            const downIdx = idx + colLength;
            if (row < rowLength - 1 && freshIdx.has(downIdx)) {
                canRotten = true;
                rottingIdx.add(downIdx);
                freshIdx.delete(downIdx);
            }

            // left
            const leftIdx = idx - 1;
            if (col > 0 && freshIdx.has(leftIdx)) {
                canRotten = true;
                rottingIdx.add(leftIdx);
                freshIdx.delete(leftIdx);
            }

            // right
            const rightIdx = idx + 1;
            if (col < grid[0].length - 1 && freshIdx.has(rightIdx)) {
                canRotten = true;
                rottingIdx.add(rightIdx);
                freshIdx.delete(rightIdx);
            }

            rottingIdx.delete(idx);
        }

        if (canRotten) {
            minutes++;
        } else {
            break;
        }
    }

    // Return -1 if there are some fresh oranges that never be rotten
    return freshIdx.size === 0 ? minutes : -1;
};

module.exports = orangesRotting;
