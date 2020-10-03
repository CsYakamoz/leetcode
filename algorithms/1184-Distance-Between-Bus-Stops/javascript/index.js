/**
 * @param {number} num
 * @param {number} n
 */
const dec = (num, n) => (num !== 0 ? num - 1 : n - 1);

/**
 * @param {number[]} distance
 * @param {number} start
 * @param {number} destination
 * @return {number}
 */
const distanceBetweenBusStops = function(distance, start, destination) {
    let clockwise = 0;
    let clockwiseStart = start;
    while (clockwiseStart !== destination) {
        clockwise += distance[clockwiseStart];

        clockwiseStart = (clockwiseStart + 1) % distance.length;
    }

    let counterclockwise = 0;
    let counterclockwiseStart = start;
    while (counterclockwiseStart !== destination) {
        counterclockwise +=
            distance[dec(counterclockwiseStart, distance.length)];

        counterclockwiseStart = dec(counterclockwiseStart, distance.length);
    }

    return Math.min(clockwise, counterclockwise);
};

module.exports = distanceBetweenBusStops;
