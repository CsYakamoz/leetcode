const {
    structure: { TreeNode },
} = require('@lib/javascript');

/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */

class IdxToLastSum {
    /**
     * @param {number} target
     */
    constructor(target) {
        this.count = 0;
        this.target = target;
        /** @type {number[]} */
        this.idxToLastSum = [];
    }

    push(val) {
        this.count = this.target === val ? 1 : 0;
        this.idxToLastSum = this.idxToLastSum.map((sum) => {
            sum += val;
            if (sum === this.target) {
                this.count++;
            }

            return sum;
        });

        this.idxToLastSum.push(val);
    }

    pop() {
        const val = this.idxToLastSum.pop();
        this.idxToLastSum = this.idxToLastSum.map((sum) => (sum -= val));
    }
}

const helper = (root, _) => {
    if (root === null) {
        return;
    }

    _.valList.push(root.val);
    _.res += _.valList.count;
    helper(root.left, _);
    helper(root.right, _);
    _.valList.pop();
};

/**
 * @param {TreeNode} root
 * @param {number} sum
 * @return {number}
 */
const pathSum = function (root, sum) {
    const _ = {
        valList: new IdxToLastSum(sum),
        res: 0,
    };

    helper(root, _);

    return _.res;
};

module.exports = pathSum;
