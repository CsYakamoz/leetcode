const { deepStrictEqual } = require('power-assert');

const {
    arrayToBinaryTree,
    arrayToLinkedList,
    preOrderTraversal,
    inOrderTraversal,
    postOrderTraversal,
} = require('../utils');
const { TreeNode, ListNode } = require('../structure');

/*
 *    1
 *  /   \
 * 2     3
 *  \     \
 *   5     4
 */
const TreeRoot1 = new TreeNode(
    1,
    new TreeNode(2, null, new TreeNode(5)),
    new TreeNode(3, null, new TreeNode(4))
);

/*
 *    6
 *  /   \
 * 3     5
 *  \    /
 *   2  0
 *    \
 *     1
 */
const TreeRoot2 = new TreeNode(
    6,
    new TreeNode(3, null, new TreeNode(2, null, new TreeNode(1))),
    new TreeNode(5, new TreeNode(0))
);

// 1->2->3->4->5->NULL
const LinkedListHead = new ListNode(
    1,
    new ListNode(2, new ListNode(3, new ListNode(4, new ListNode(5))))
);

describe('lib/javascript', () => {
    describe('Tree and Related Utils', () => {
        it('arrayToBinaryTree', () => {
            deepStrictEqual(
                arrayToBinaryTree([1, 2, 3, null, 5, null, 4]),
                TreeRoot1
            );

            deepStrictEqual(
                arrayToBinaryTree([6, 3, 5, null, 2, 0, null, null, 1]),
                TreeRoot2
            );
        });

        it('preOrderTraversal', () => {
            deepStrictEqual(preOrderTraversal(TreeRoot1), [1, 2, 5, 3, 4]);

            deepStrictEqual(preOrderTraversal(TreeRoot2), [6, 3, 2, 1, 5, 0]);
        });

        it('inOrderTraversal', () => {
            deepStrictEqual(inOrderTraversal(TreeRoot1), [2, 5, 1, 3, 4]);

            deepStrictEqual(inOrderTraversal(TreeRoot2), [3, 2, 1, 6, 0, 5]);
        });

        it('postOrderTraversal', () => {
            deepStrictEqual(postOrderTraversal(TreeRoot1), [5, 2, 4, 3, 1]);

            deepStrictEqual(postOrderTraversal(TreeRoot2), [1, 2, 3, 0, 5, 6]);
        });

    /*
     * it('prettyPrintBinaryTree', () => {
     *     const output1 = [
     *         '│       ┌── 4',
     *         '│   ┌── 3',
     *         '└── 1',
     *         '    │   ┌── 5',
     *         '    └── 2',
     *     ].join('\n');
     *     deepStrictEqual(prettyPrintBinaryTree(TreeRoot1), output1);
     * });
     */
    });

    describe('LinkedList and Related Utils', () => {
        it('arrayToLinkedList', () => {
            deepStrictEqual(arrayToLinkedList([1, 2, 3, 4, 5]), LinkedListHead);
        });
    });
});
