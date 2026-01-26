// Definition for singly-linked list.
function ListNode(val, next) {
    this.val = (val===undefined ? 0 : val);
    this.next = (next===undefined ? null : next);
}

/**
 * @param {ListNode} list1
 * @param {ListNode} list2
 * @return {ListNode}
 */
var mergeTwoLists = function(list1, list2) {
    if (list1 === null && list2 === null) {
        return null;
    } else if (list1 === null) {
        return list2;
    } else if (list2 === null) {
        return list1;
    }

    let dummy = new ListNode();

    let tail = dummy;

    while (list1 !== null && list2 !== null) {
        if (list1.val < list2.val) {
            tail.next = list1;
            list1 = list1.next;
        } else {
            tail.next = list2;
            list2 = list2.next;
        }
        tail = tail.next;
    }

    // append remaining list head to the result list tail
    if (list1 !== null) {
        tail.next = list1;
    }

    if (list2 !== null) {
        tail.next = list2;
    }

    // return the next element appended to the dummy list
    return dummy.next;
};

module.exports = {
    mergeTwoLists,
}