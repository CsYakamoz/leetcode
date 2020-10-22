const {
    structure: { TreeNode },
} = require('@lib/javascript');

/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */

/**
 * @param {TreeNode} root
 * @returns {Iterator<number, number>}
 */
const getIterator = (root) => {
    /** @type {TreeNode[]} */
    const stack = [];

    return {
        next: () => {
            while (root !== null || stack.length > 0) {
                if (root !== null) {
                    stack.push(root);
                    root = root.left;
                } else {
                    const curr = stack.pop();
                    root = curr.right;
                    return { value: curr.val, done: false };
                }
            }

            return { done: true };
        },
    };
};

/**
 * @param {TreeNode} root1
 * @param {TreeNode} root2
 * @return {number[]}
 */
const getAllElements = function(root1, root2) {
    const result = [];

    const iter1 = getIterator(root1);
    const iter2 = getIterator(root2);

    let next1 = iter1.next();
    let next2 = iter2.next();

    while (!next1.done || !next2.done) {
        if (!next1.done && !next2.done) {
            if (next1.value < next2.value) {
                result.push(next1.value);
                next1 = iter1.next();
            } else {
                result.push(next2.value);
                next2 = iter2.next();
            }
        } else if (next1.done) {
            result.push(next2.value);
            next2 = iter2.next();
        } else {
            result.push(next1.value);
            next1 = iter1.next();
        }
    }

    return result;
};

module.exports = getAllElements;
