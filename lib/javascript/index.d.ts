declare class TreeNode {
    val: any;
    left: TreeNode;
    right: TreeNode;
    constructor(val: any, left: TreeNode, right: TreeNode);
}

declare class ListNode {
    val: any;
    next: TreeNode;
    constructor(val: any, next: ListNode);
}

export const structure: {
    TreeNode: TreeNode;
    ListNode: ListNode;
};

export const utils: {
    arrayToLinkedList(arr: any[]): ListNode;
    arrayToBinaryTree(arr: any[]): TreeNode;
    preOrderTraversal(root: TreeNode): any[];
    inOrderTraversal(root: TreeNode): any[];
    postOrderTraversal(root: TreeNode): any[];
    prettyPrintBinaryTree(root: TreeNode): void;
    prettyPrintLinkedList(root: TreeNode): void;
};
