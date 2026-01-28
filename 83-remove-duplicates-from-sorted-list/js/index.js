/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var deleteDuplicates = function(head) {
    // iterate over the list, if the next element value equals to the current val, reset current next
    
    let tail = head;

    while (tail && tail.next) {
        if (tail.val === tail.next.val) {
            tail.next = tail.next.next;
        } else {
            tail = tail.next;
        }
    }

    // 1, 1, 2, 3, 3
    // tail = 1
    // tail.next = 1
    return head;
};
