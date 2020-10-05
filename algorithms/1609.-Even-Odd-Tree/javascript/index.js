const { TreeNode } = require('@javascript/language-helper');

/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */

/**
 * @param {TreeNode[]} arr
 * @param {boolean} isEvenLevel
 * @returns {boolean}
 */
const check = (arr, isEvenLevel) => {
    if (arr.some((item) => item.val % 2 === (isEvenLevel ? 0 : 1))) {
        return false;
    }

    for (let i = 1, len = arr.length; i < len; i++) {
        if (
            (isEvenLevel && arr[i - 1].val >= arr[i].val) ||
            (!isEvenLevel && arr[i - 1].val <= arr[i].val)
        ) {
            return false;
        }
    }

    return true;
};

/**
 * @param {TreeNode} root
 * @return {boolean}
 */
const isEvenOddTree = function(root) {
    let level = 0;
    const queue = [root];

    while (queue.length > 0) {
        const length = queue.length;

        if (!check(queue, level % 2 === 0)) {
            return false;
        }

        for (let i = 0; i < length; i++) {
            const curr = queue.shift();
            if (curr.left !== null) {
                queue.push(curr.left);
            }

            if (curr.right !== null) {
                queue.push(curr.right);
            }
        }

        level++;
    }

    return true;
};

module.exports = isEvenOddTree;
