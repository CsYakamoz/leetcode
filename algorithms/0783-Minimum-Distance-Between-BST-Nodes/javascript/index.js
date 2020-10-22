const {
    structure: { TreeNode },
} = require('@lib/javascript');

/**
 * @param {TreeNode} root
 * @return {number}
 */
const minDiffInBST = function(root) {
    const stack = [];

    let prevVal = null;
    let minDiff = Number.MAX_SAFE_INTEGER;

    let curr = root;
    while (curr !== null || stack.length !== 0) {
        if (curr !== null) {
            stack.push(curr);
            curr = curr.left;
        } else {
            curr = stack.pop();
            if (prevVal !== null) {
                minDiff = Math.min(minDiff, curr.val - prevVal);
            }

            prevVal = curr.val;
            curr = curr.right;
        }
    }

    return minDiff;
};

module.exports = minDiffInBST;
