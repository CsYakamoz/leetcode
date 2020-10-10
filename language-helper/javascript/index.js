class ListNode {
    constructor(val) {
        this.val = val;
        /** @type {ListNode} */
        this.next = null;
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

/**
 * @param {any[]} arr
 * @return {ListNode}
 */
const arrayToLinkedList = (arr) => {
    if (arr === undefined || arr.length === 0) {
        return null;
    }

    const head = new ListNode(null);
    arr.reduce((prev, val) => {
        const curr = new ListNode(val);
        prev.next = curr;
        return curr;
    }, head);

    return head.next;
};

/**
 * @param {any[]} arr
 * @returns {TreeNode}
 */
const arrayToBinaryTree = (arr) => {
    if (arr === undefined || arr.length === 0) {
        return null;
    }

    const root = new TreeNode(arr[0]);
    const queue = [root];

    let isLeftChild = true;

    for (let i = 1, len = arr.length; i < len; i++) {
        const val = arr[i];
        const curr = val !== null ? new TreeNode(val) : null;

        if (isLeftChild) {
            queue[0].left = curr;
        } else {
            queue[0].right = curr;
            queue.shift();
        }

        if (curr !== null) {
            queue.push(curr);
        }

        isLeftChild = !isLeftChild;
    }

    return root;
};

/**
 * @param {TreeNode} root
 * @returns {any[]}
 */
const preOrderTraversal = (root) => {
    if (root === null) {
        return [];
    }

    const res = [];
    const stack = [root];

    while (stack.length !== 0) {
        const curr = stack.pop();
        res.push(curr.val);

        if (curr.right !== null) {
            stack.push(curr.right);
        }

        if (curr.left !== null) {
            stack.push(curr.left);
        }
    }

    return res;
};

/**
 * @param {TreeNode} root
 * @returns {any[]}
 */
const inOrderTraversal = (root) => {
    if (root === null) {
        return [];
    }

    const res = [];
    /** @type {TreeNode[]} */
    const stack = [];

    while (root !== null || stack.length !== 0) {
        if (root !== null) {
            stack.push(root);
            root = root.left;
        } else {
            const curr = stack.pop();
            res.push(curr.val);
            root = curr.right;
        }
    }

    return res;
};

/**
 * TODO: loop method instead of recursion method
 * @param {TreeNode} root
 * @returns {any[]}
 */
const postOrderTraversal = (root) => {
    /**
   * @param {TreeNode} root
   * @param {any[]} res
   * @returns {any[]}
   */
    const helper = (root, res) => {
        if (root === null) {
            return res;
        }

        helper(root.left, res);
        helper(root.right);
        res.push(root.val);

        return res;
    };

    return helper(root, []);
};

/**
 * @param {TreeNode} node
 */
const prettyPrintBinaryTree = (node, prefix = '', isLeft = true) => {
    if (node === null) {
        process.stdout.write('Empty tree\n');
        return;
    }

    if (node.right) {
        prettyPrintBinaryTree(
            node.right,
            prefix + (isLeft ? '│   ' : '    '),
            false
        );
    }

    process.stdout.write(prefix + (isLeft ? '└── ' : '┌── ') + node.val + '\n');

    if (node.left) {
        prettyPrintBinaryTree(node.left, prefix + (isLeft ? '    ' : '│   '), true);
    }
};

/**
 * @param {ListNode} node
 */
const prettyPrintLinkedList = (node) => {
    while (node && node.next) {
        process.stdout.write(node.val + ' -> ');
        node = node.next;
    }

    if (node) {
        process.stdout.write(node.val + '\n');
    } else {
        process.stdout.write('Empty LinkedList\n');
    }
};

module.exports = {
    ListNode,
    TreeNode,

    arrayToLinkedList,
    arrayToBinaryTree,

    preOrderTraversal,
    inOrderTraversal,
    postOrderTraversal,

    prettyPrintBinaryTree,
    prettyPrintLinkedList,
};
