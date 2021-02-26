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

/**
 * @param {number[]} preorder
 * @return {TreeNode}
 */
var bstFromPreorder = function (preorder) {
    if (preorder.length === 0) {
        return null;
    }

    const rootVal = preorder[0];
    const root = new TreeNode(rootVal);

    /*
     * find the first element that is greater than rootVal
     * after the loop
     * end will point it if exist, else point the preorder.length
     */
    let begin = 1;
    let end = preorder.length;
    while (begin < end) {
        const mid = begin + Math.floor((end - begin) / 2);
        const val = preorder[mid];

        if (rootVal < val) {
            end = mid;
        } else {
            begin = mid + 1;
        }
    }

    root.left = bstFromPreorder(preorder.slice(1, end));
    root.right = bstFromPreorder(preorder.slice(end, preorder.length));

    return root;
};

module.exports = bstFromPreorder;
