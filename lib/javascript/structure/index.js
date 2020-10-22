class ListNode {
    constructor(val, next) {
        this.val = val;
        /** @type {ListNode} */
        this.next = next === undefined ? null : next;
    }
}

class TreeNode {
    constructor(val, left, right) {
        this.val = val;
        /** @type {TreeNode} */
        this.left = left === undefined ? null : left;
        /** @type {TreeNode} */
        this.right = right === undefined ? null : right;
    }
}

module.exports = { ListNode, TreeNode };
