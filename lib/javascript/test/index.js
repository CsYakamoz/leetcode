const { deepStrictEqual } = require('power-assert');

const {
    arrayToBinaryTree,
    arrayToLinkedList,
    preOrderTraversal,
    inOrderTraversal,
    postOrderTraversal,
} = require('../utils');
const { TreeNode, ListNode } = require('../structure');

const treeTestList = [
    {
        // empty tree
        arr: [],
        treeRoot: null,
        preOrder: [],
        inOrder: [],
        postOrder: [],
    },
    {
        /*
         *    1
         *  /   \
         * 2     3
         *  \     \
         *   5     4
         */
        arr: [1, 2, 3, null, 5, null, 4],
        treeRoot: new TreeNode(
            1,
            new TreeNode(2, null, new TreeNode(5)),
            new TreeNode(3, null, new TreeNode(4))
        ),
        preOrder: [1, 2, 5, 3, 4],
        inOrder: [2, 5, 1, 3, 4],
        postOrder: [5, 2, 4, 3, 1],
    },
    {
        /*
         *    6
         *  /   \
         * 3     5
         *  \    /
         *   2  0
         *    \
         *     1
         */
        arr: [6, 3, 5, null, 2, 0, null, null, 1],
        treeRoot: new TreeNode(
            6,
            new TreeNode(3, null, new TreeNode(2, null, new TreeNode(1))),
            new TreeNode(5, new TreeNode(0))
        ),
        preOrder: [6, 3, 2, 1, 5, 0],
        inOrder: [3, 2, 1, 6, 0, 5],
        postOrder: [1, 2, 3, 0, 5, 6],
    },
];

const linkedListTestList = [
    {
        arr: [],
        head: null,
    },
    {
        // 1->2->3->4->5->NULL
        arr: [1, 2, 3, 4, 5],
        head: new ListNode(
            1,
            new ListNode(2, new ListNode(3, new ListNode(4, new ListNode(5))))
        ),
    },
];

describe('lib/javascript', () => {
    describe('Tree and Related Utils', () => {
        for (const [idx, test] of Object.entries(treeTestList)) {
            it(`TestCase[${idx}]: arrayToBinaryTree`, () => {
                deepStrictEqual(arrayToBinaryTree(test.arr), test.treeRoot);
            });

            it(`TestCase[${idx}]: preOrderTraversal`, () => {
                deepStrictEqual(
                    preOrderTraversal(test.treeRoot),
                    test.preOrder
                );
            });

            it(`TestCase[${idx}]: inOrderTraversal`, () => {
                deepStrictEqual(inOrderTraversal(test.treeRoot), test.inOrder);
            });

            it(`TestCase[${idx}]: postOrderTraversal`, () => {
                deepStrictEqual(
                    postOrderTraversal(test.treeRoot),
                    test.postOrder
                );
            });
        }
    });

    describe('LinkedList and Related Utils', () => {
        for (const [idx, test] of Object.entries(linkedListTestList)) {
            it(`TestCase[${idx}]: arrayToLinkedList`, () => {
                deepStrictEqual(arrayToLinkedList(test.arr), test.head);
            });
        }
    });
});
