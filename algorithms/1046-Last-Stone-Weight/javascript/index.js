class PriorityQueue {
    constructor(rawData, compare) {
        this._data = [null].concat(Array.from(rawData));
        this._compare = compare;

        for (let i = Math.floor(this.length / 2); i > 0; i--) {
            this._sink(i);
        }
    }

    get length() {
        return this._data.length - 1;
    }

    insert(elem) {
        this._data.push(elem);
        this._swim(this.length);
    }

    pop() {
        if (this.length === 0) {
            throw new Error('empty queue');
        }

        const val = this._data[1];

        this._exch(1, this.length);
        this._data.pop();
        this._sink(1);

        return val;
    }

    head() {
        if (this.length === 0) {
            throw new Error('empty queue');
        }

        return this._data[1];
    }

    empty() {
        return this.length === 0;
    }

    _swim(i) {
        while (
            i > 1 &&
            this._compare(this._data[i], this._data[Math.floor(i / 2)]) < 0
        ) {
            this._exch(i, Math.floor(i / 2));
            i = Math.floor(i / 2);
        }
    }

    _sink(i) {
        while (2 * i <= this.length) {
            let j = 2 * i;
            if (
                j < this.length &&
                this._compare(this._data[j], this._data[j + 1]) > 0
            ) {
                j++;
            }

            if (this._compare(this._data[i], this._data[j]) > 0) {
                this._exch(i, j);
                i = j;
            } else {
                break;
            }
        }
    }

    _exch(i, j) {
        const tmp = this._data[j];
        this._data[j] = this._data[i];
        this._data[i] = tmp;
    }
}

/**
 * @param {number[]} stones
 * @return {number}
 */
var lastStoneWeight = function(stones) {
    if (stones.length < 2) {
        return stones.pop() || 0;
    }

    const pq = new PriorityQueue(stones, (a, b) => b - a);

    while (pq.length > 1) {
        const i = pq.pop();
        const j = pq.pop();

        if (i === j) {
            continue;
        }

        pq.insert(Math.abs(i - j));
    }

    return pq.empty() ? 0 : pq.pop();
};

module.exports = lastStoneWeight;
